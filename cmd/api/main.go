package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilteber/message-scheduler/internal/cache"
	"github.com/ilteber/message-scheduler/internal/config"
	"github.com/ilteber/message-scheduler/internal/database"
	"github.com/ilteber/message-scheduler/internal/handler"
	"github.com/ilteber/message-scheduler/internal/router"
	"github.com/ilteber/message-scheduler/internal/scheduler"
	"github.com/ilteber/message-scheduler/internal/service"
)

func main() {
	log.Println("Starting Message Scheduler Service...")

	// Load configuration
	cfg := config.Load()

	// Initialize database
	if err := database.Initialize(&cfg.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database connected successfully")

	// Initialize Redis
	if err := cache.Initialize(&cfg.Redis); err != nil {
		log.Printf("Warning: Failed to initialize Redis: %v", err)
		log.Println("Continuing without Redis caching...")
	} else {
		log.Println("Redis connected successfully")
	}

	// Initialize services
	messageService := service.NewMessageService(&cfg.Webhook)

	// Initialize scheduler
	sched := scheduler.NewScheduler(messageService, &cfg.Scheduler)

	// Start scheduler automatically on deployment
	log.Println("Auto-starting scheduler...")
	if err := sched.Start(); err != nil {
		log.Printf("Warning: Failed to auto-start scheduler: %v", err)
	}

	// Initialize handlers
	h := handler.NewHandler(messageService, sched)

	// Setup router
	mux := router.SetupRouter(h)

	// Wrap with logging middleware
	httpHandler := router.LoggingMiddleware(mux)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("\nShutting down gracefully...")

		// Stop scheduler
		sched.Stop()

		// Close database
		if err := database.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}

		// Close Redis
		if err := cache.Close(); err != nil {
			log.Printf("Error closing Redis: %v", err)
		}

		log.Println("Shutdown complete")
		os.Exit(0)
	}()

	// Start server
	serverAddr := ":" + cfg.Server.Port
	log.Printf("Server starting on %s", serverAddr)
	log.Printf("API endpoints:")
	log.Printf("  GET  /health")
	log.Printf("  POST /api/scheduler (payload: {\"command\": \"start\"} or {\"command\": \"stop\"})")
	log.Printf("  POST /api/messages/sent (payload: {\"phone_number\": \"all\", \"limit\": 50, \"offset\": 0})")

	if err := http.ListenAndServe(serverAddr, httpHandler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
