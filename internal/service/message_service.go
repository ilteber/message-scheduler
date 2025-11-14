package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/ilteber/message-scheduler/internal/cache"
	"github.com/ilteber/message-scheduler/internal/config"
	"github.com/ilteber/message-scheduler/internal/database"
	"github.com/ilteber/message-scheduler/internal/models"
)

// MessageService handles message-related operations
type MessageService struct {
	db      *sql.DB
	webhook *config.WebhookConfig
}

// NewMessageService creates a new message service
func NewMessageService(webhook *config.WebhookConfig) *MessageService {
	return &MessageService{
		db:      database.GetDB(),
		webhook: webhook,
	}
}

// SendPendingMessages retrieves and sends pending messages
func (s *MessageService) SendPendingMessages(limit int) (int, int, error) {
	// Fetch pending messages
	query := `
		SELECT id, phone_number, content, status, created_at, updated_at
		FROM messages
		WHERE status = $1
		ORDER BY created_at ASC
		LIMIT $2
	`

	rows, err := s.db.Query(query, models.StatusPending, limit)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to fetch pending messages: %w", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		var sentAt sql.NullTime
		var messageID sql.NullString

		err := rows.Scan(
			&msg.ID,
			&msg.PhoneNumber,
			&msg.Content,
			&msg.Status,
			&msg.CreatedAt,
			&msg.UpdatedAt,
		)
		if err != nil {
			return 0, 0, fmt.Errorf("failed to scan message: %w", err)
		}

		if sentAt.Valid {
			msg.SentAt = &sentAt.Time
		}
		if messageID.Valid {
			msg.MessageID = messageID.String
		}

		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return 0, 0, fmt.Errorf("error iterating messages: %w", err)
	}

	if len(messages) == 0 {
		return 0, 0, nil
	}

	sent := 0
	failed := 0

	for i := range messages {
		if err := s.sendMessage(&messages[i]); err != nil {
			log.Printf("Failed to send message ID %d: %v", messages[i].ID, err)
			failed++
		} else {
			sent++
		}
	}

	return sent, failed, nil
}

// sendMessage sends a single message via webhook
func (s *MessageService) sendMessage(message *models.Message) error {
	// Prepare request payload
	payload := models.SendMessageRequest{
		To:      message.PhoneNumber,
		Content: message.Content,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		s.updateMessageStatus(message.ID, models.StatusFailed, "", nil)
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", s.webhook.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		s.updateMessageStatus(message.ID, models.StatusFailed, "", nil)
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ins-auth-key", s.webhook.AuthKey)

	// Send request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		s.updateMessageStatus(message.ID, models.StatusFailed, "", nil)
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.updateMessageStatus(message.ID, models.StatusFailed, "", nil)
		return fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		s.updateMessageStatus(message.ID, models.StatusFailed, "", nil)
		return fmt.Errorf("webhook returned non-success status: %d - %s", resp.StatusCode, string(body))
	}

	// Parse response
	var webhookResp models.SendMessageResponse
	if err := json.Unmarshal(body, &webhookResp); err != nil {
		log.Printf("Warning: failed to parse response for message ID %d: %v", message.ID, err)
		// Still mark as sent if we got a success status code
		webhookResp.MessageID = fmt.Sprintf("unknown_%d", time.Now().Unix())
	}

	// Update message status
	sentAt := time.Now()
	if err := s.updateMessageStatus(message.ID, models.StatusSent, webhookResp.MessageID, &sentAt); err != nil {
		return fmt.Errorf("failed to update message status: %w", err)
	}

	// Cache to Redis (bonus feature)
	if err := cache.CacheMessageSent(webhookResp.MessageID, message.PhoneNumber, message.Content, sentAt); err != nil {
		log.Printf("Warning: failed to cache message ID %d to Redis: %v", message.ID, err)
		// Don't fail the entire operation if caching fails, Write a log to the console (or db in prod case)
	}

	log.Printf("Message ID %d sent successfully - MessageID: %s", message.ID, webhookResp.MessageID)
	return nil
}

// updateMessageStatus updates the status of a message
func (s *MessageService) updateMessageStatus(messageID uint, status models.MessageStatus, webhookMessageID string, sentAt *time.Time) error {
	query := `
		UPDATE messages
		SET status = $1, message_id = $2, sent_at = $3, updated_at = $4
		WHERE id = $5
	`

	var sentAtValue interface{}
	if sentAt != nil {
		sentAtValue = sentAt
	}

	_, err := s.db.Exec(query, status, webhookMessageID, sentAtValue, time.Now(), messageID)
	if err != nil {
		return fmt.Errorf("failed to update message: %w", err)
	}

	return nil
}

// GetSentMessagesWithFilters retrieves all sent messages with optional date filters
func (s *MessageService) GetSentMessagesWithFilters(limit, offset int, sentAfter, sentBefore string) (*models.MessageListResponse, error) {
	var args []interface{}
	countQuery := `SELECT COUNT(*) FROM messages WHERE status = $1`
	query := `SELECT id, phone_number, content, status, message_id, sent_at, created_at, updated_at FROM messages WHERE status = $1`
	args = append(args, models.StatusSent)

	// Add date filters if provided
	if sentAfter != "" {
		afterTime, err := time.Parse(time.RFC3339, sentAfter)
		if err != nil {
			return nil, fmt.Errorf("invalid sent_after format, use RFC3339 (e.g., 2025-11-14T10:00:00Z): %w", err)
		}
		args = append(args, afterTime)
		countQuery += fmt.Sprintf(" AND sent_at > $%d", len(args))
		query += fmt.Sprintf(" AND sent_at > $%d", len(args))
	}

	if sentBefore != "" {
		beforeTime, err := time.Parse(time.RFC3339, sentBefore)
		if err != nil {
			return nil, fmt.Errorf("invalid sent_before format, use RFC3339 (e.g., 2025-11-14T10:00:00Z): %w", err)
		}
		args = append(args, beforeTime)
		countQuery += fmt.Sprintf(" AND sent_at < $%d", len(args))
		query += fmt.Sprintf(" AND sent_at < $%d", len(args))
	}

	// Count total sent messages with filters
	var total int64
	err := s.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to count messages: %w", err)
	}

	// Add ORDER BY, LIMIT, OFFSET
	args = append(args, limit, offset)
	query += fmt.Sprintf(" ORDER BY sent_at DESC LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch messages: %w", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		var sentAt sql.NullTime
		var messageID sql.NullString

		err := rows.Scan(
			&msg.ID,
			&msg.PhoneNumber,
			&msg.Content,
			&msg.Status,
			&messageID,
			&sentAt,
			&msg.CreatedAt,
			&msg.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}

		if sentAt.Valid {
			msg.SentAt = &sentAt.Time
		}
		if messageID.Valid {
			msg.MessageID = messageID.String
		}

		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating messages: %w", err)
	}

	return &models.MessageListResponse{
		Total:    total,
		Messages: messages,
	}, nil
}

// GetSentMessagesByPhoneWithFilters retrieves sent messages for a specific phone number with optional date filters
func (s *MessageService) GetSentMessagesByPhoneWithFilters(phoneNumber string, limit, offset int, sentAfter, sentBefore string) (*models.MessageListResponse, error) {
	var args []interface{}
	args = append(args, phoneNumber, models.StatusSent)

	// Build base queries
	checkQuery := `SELECT COUNT(*) FROM messages WHERE phone_number = $1 AND status = $2`
	countQuery := `SELECT COUNT(*) FROM messages WHERE phone_number = $1 AND status = $2`
	query := `SELECT id, phone_number, content, status, message_id, sent_at, created_at, updated_at FROM messages WHERE phone_number = $1 AND status = $2`

	// Add date filters if provided
	if sentAfter != "" {
		afterTime, err := time.Parse(time.RFC3339, sentAfter)
		if err != nil {
			return nil, fmt.Errorf("invalid sent_after format, use RFC3339 (e.g., 2025-11-14T10:00:00Z): %w", err)
		}
		args = append(args, afterTime)
		checkQuery += fmt.Sprintf(" AND sent_at > $%d", len(args))
		countQuery += fmt.Sprintf(" AND sent_at > $%d", len(args))
		query += fmt.Sprintf(" AND sent_at > $%d", len(args))
	}

	if sentBefore != "" {
		beforeTime, err := time.Parse(time.RFC3339, sentBefore)
		if err != nil {
			return nil, fmt.Errorf("invalid sent_before format, use RFC3339 (e.g., 2025-11-14T10:00:00Z): %w", err)
		}
		args = append(args, beforeTime)
		checkQuery += fmt.Sprintf(" AND sent_at < $%d", len(args))
		countQuery += fmt.Sprintf(" AND sent_at < $%d", len(args))
		query += fmt.Sprintf(" AND sent_at < $%d", len(args))
	}

	// First check if phone number exists with filters
	var exists int
	err := s.db.QueryRow(checkQuery, args...).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("failed to check phone number: %w", err)
	}

	if exists == 0 {
		return nil, fmt.Errorf("phone number not found")
	}

	// Count total sent messages for this phone number with filters
	var total int64
	err = s.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("failed to count messages: %w", err)
	}

	// Add ORDER BY, LIMIT, OFFSET
	args = append(args, limit, offset)
	query += fmt.Sprintf(" ORDER BY sent_at DESC LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch messages: %w", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		var sentAt sql.NullTime
		var messageID sql.NullString

		err := rows.Scan(
			&msg.ID,
			&msg.PhoneNumber,
			&msg.Content,
			&msg.Status,
			&messageID,
			&sentAt,
			&msg.CreatedAt,
			&msg.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}

		if sentAt.Valid {
			msg.SentAt = &sentAt.Time
		}
		if messageID.Valid {
			msg.MessageID = messageID.String
		}

		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating messages: %w", err)
	}

	return &models.MessageListResponse{
		Total:    total,
		Messages: messages,
	}, nil
}
