package router

import (
	"log"
	"net/http"

	"github.com/ilteber/message-scheduler/internal/handler"
)

// SetupRouter configures all routes using standard net/http
func SetupRouter(h *handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	// Might not be exposed to user in production, can be removed if needed as the requirement is to provide 2 endpoints only
	mux.HandleFunc("/health", h.HealthCheck)

	// 2 Required API Endpoints:
	// 1. Scheduler control (start/stop via command in payload)
	mux.HandleFunc("/api/scheduler", h.SchedulerControl)

	// 2. Retrieve sent messages (with phone number filter and pagination)
	mux.HandleFunc("/api/messages/sent", h.GetSentMessages)

	return mux
}

// LoggingMiddleware logs HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
