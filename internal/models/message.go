package models

import (
	"time"
)

// MessageStatus represents the status of a message
type MessageStatus string

const (
	StatusPending MessageStatus = "pending"
	StatusSent    MessageStatus = "sent"
	StatusFailed  MessageStatus = "failed"
)

// Message represents a message to be sent
type Message struct {
	ID          uint          `gorm:"primaryKey" json:"id"`
	PhoneNumber string        `gorm:"type:varchar(20);not null" json:"phone_number"`
	Content     string        `gorm:"type:varchar(500);not null" json:"content"`
	Status      MessageStatus `gorm:"type:varchar(20);default:'pending';index" json:"status"`
	MessageID   string        `gorm:"type:varchar(100)" json:"message_id,omitempty"`
	SentAt      *time.Time    `gorm:"index" json:"sent_at,omitempty"`
	CreatedAt   time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name for Message model
func (Message) TableName() string {
	return "messages"
}

// SendMessageRequest represents the request payload for sending a message
type SendMessageRequest struct {
	To      string `json:"to" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// SendMessageResponse represents the response from webhook
type SendMessageResponse struct {
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}

// MessageListResponse represents the response for listing messages
type MessageListResponse struct {
	Total    int64     `json:"total"`
	Messages []Message `json:"messages"`
}
