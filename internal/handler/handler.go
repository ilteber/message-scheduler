package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ilteber/message-scheduler/internal/scheduler"
	"github.com/ilteber/message-scheduler/internal/service"
)

// Handler contains all HTTP handlers
type Handler struct {
	messageService *service.MessageService
	scheduler      *scheduler.Scheduler
}

// NewHandler creates a new handler
func NewHandler(messageService *service.MessageService, scheduler *scheduler.Scheduler) *Handler {
	return &Handler{
		messageService: messageService,
		scheduler:      scheduler,
	}
}

// Response represents a generic API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SchedulerStatusResponse represents scheduler status
type SchedulerStatusResponse struct {
	IsRunning bool   `json:"is_running"`
	Message   string `json:"message"`
}

// SchedulerCommandRequest represents the scheduler command request
type SchedulerCommandRequest struct {
	Command string `json:"command"`
}

// SentMessagesRequest represents the request to get sent messages
type SentMessagesRequest struct {
	PhoneNumber string `json:"phone_number"` // "all" for all messages, or specific phone number
	Limit       int    `json:"limit"`        // pagination limit
	Offset      int    `json:"offset"`       // pagination offset
	SentAfter   string `json:"sent_after"`   // optional: get messages sent after this date (RFC3339 format)
	SentBefore  string `json:"sent_before"`  // optional: get messages sent before this date (RFC3339 format)
}

// writeJSON writes JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// SchedulerControl handles POST /api/scheduler with command in payload
func (h *Handler) SchedulerControl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var req SchedulerCommandRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request payload: " + err.Error(),
		})
		return
	}

	switch req.Command {
	case "start":
		if err := h.scheduler.Start(); err != nil {
			writeJSON(w, http.StatusInternalServerError, Response{
				Success: false,
				Error:   err.Error(),
			})
			return
		}

		writeJSON(w, http.StatusOK, Response{
			Success: true,
			Message: "Scheduler started successfully",
			Data: SchedulerStatusResponse{
				IsRunning: true,
				Message:   "Automatic message sending is now active",
			},
		})

	case "stop":
		h.scheduler.Stop()

		writeJSON(w, http.StatusOK, Response{
			Success: true,
			Message: "Scheduler stopped successfully",
			Data: SchedulerStatusResponse{
				IsRunning: false,
				Message:   "Automatic message sending is now inactive",
			},
		})

	default:
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid command. Must be 'start' or 'stop'",
		})
	}
}

// GetSentMessages handles POST /api/messages/sent with phone number filter
func (h *Handler) GetSentMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var req SentMessagesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Error:   "Invalid request payload: " + err.Error(),
		})
		return
	}

	// Set default pagination values
	if req.Limit <= 0 {
		req.Limit = 50 // Default: 50
	}
	// Cap maximum limit at 50
	if req.Limit > 50 {
		req.Limit = 50 // Max: 50
	}
	// Ensure offset is non-negative
	if req.Offset < 0 {
		req.Offset = 0 // Min: 0
	}
	// Optional: Cap maximum offset to prevent abuse (e.g., 10000)
	// Uncomment if you want to limit deep pagination
	// if req.Offset > 10000 {
	// 	req.Offset = 10000
	// }

	// Validate phone_number field
	if req.PhoneNumber == "" {
		writeJSON(w, http.StatusBadRequest, Response{
			Success: false,
			Error:   "phone_number field is required",
		})
		return
	}

	var messages interface{}
	var err error

	if req.PhoneNumber == "all" {
		// Get all sent messages
		messages, err = h.messageService.GetSentMessagesWithFilters(req.Limit, req.Offset, req.SentAfter, req.SentBefore)
	} else {
		// Get sent messages for specific phone number
		messages, err = h.messageService.GetSentMessagesByPhoneWithFilters(req.PhoneNumber, req.Limit, req.Offset, req.SentAfter, req.SentBefore)
	}

	// Handle errors
	if err != nil {
		if err.Error() == "phone number not found" {
			writeJSON(w, http.StatusNotFound, Response{
				Success: false,
				Error:   "Phone number not found",
			})
			return
		}
		writeJSON(w, http.StatusInternalServerError, Response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, Response{
		Success: true,
		Data:    messages,
	})
}

// HealthCheck handles GET /health
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, Response{
		Success: true,
		Message: "API is running",
		Data: map[string]interface{}{
			"status":           "healthy",
			"scheduler_status": h.scheduler.IsRunning(),
		},
	})
}
