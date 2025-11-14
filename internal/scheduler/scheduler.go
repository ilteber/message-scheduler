package scheduler

import (
	"log"
	"sync"
	"time"

	"github.com/ilteber/message-scheduler/internal/config"
	"github.com/ilteber/message-scheduler/internal/service"
)

// Scheduler handles automatic message sending
type Scheduler struct {
	messageService *service.MessageService
	config         *config.SchedulerConfig
	ticker         *time.Ticker
	stopChan       chan bool
	isRunning      bool
	mu             sync.RWMutex
}

// NewScheduler creates a new scheduler instance
func NewScheduler(messageService *service.MessageService, cfg *config.SchedulerConfig) *Scheduler {
	return &Scheduler{
		messageService: messageService,
		config:         cfg,
		stopChan:       make(chan bool),
		isRunning:      false,
	}
}

// Start begins the automatic message sending process
func (s *Scheduler) Start() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRunning {
		log.Println("Scheduler is already running")
		return nil
	}

	s.isRunning = true
	interval := time.Duration(s.config.IntervalSeconds) * time.Second
	s.ticker = time.NewTicker(interval)

	log.Printf("Scheduler started - sending %d messages every %d seconds",
		s.config.MessagesPerBatch, s.config.IntervalSeconds)

	// Start the scheduler in a goroutine
	go s.run()

	// Send messages immediately on start
	go func() {
		if err := s.sendBatch(); err != nil {
			log.Printf("Error in initial batch send: %v", err)
		}
	}()

	return nil
}

// Stop halts the automatic message sending process
func (s *Scheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		log.Println("Scheduler is not running")
		return
	}

	s.isRunning = false
	if s.ticker != nil {
		s.ticker.Stop()
	}
	s.stopChan <- true

	log.Println("Scheduler stopped")
}

// IsRunning returns the current status of the scheduler
func (s *Scheduler) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isRunning
}

// run is the main scheduler loop
func (s *Scheduler) run() {
	for {
		select {
		case <-s.ticker.C:
			if err := s.sendBatch(); err != nil {
				log.Printf("Error sending batch: %v", err)
			}
		case <-s.stopChan:
			log.Println("Scheduler loop terminated")
			return
		}
	}
}

// sendBatch sends a batch of pending messages
func (s *Scheduler) sendBatch() error {
	s.mu.RLock()
	if !s.isRunning {
		s.mu.RUnlock()
		return nil
	}
	s.mu.RUnlock()

	log.Printf("Sending batch of %d messages...", s.config.MessagesPerBatch)

	sent, failed, err := s.messageService.SendPendingMessages(s.config.MessagesPerBatch)
	if err != nil {
		log.Printf("Batch send completed with errors - Sent: %d, Failed: %d, Error: %v",
			sent, failed, err)
		return err
	}

	if sent > 0 || failed > 0 {
		log.Printf("Batch completed - Sent: %d, Failed: %d", sent, failed)
	} else {
		log.Println("No pending messages to send")
	}

	return nil
}
