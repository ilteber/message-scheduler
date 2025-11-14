package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ilteber/message-scheduler/internal/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// Initialize establishes database connection
func Initialize(cfg *config.DatabaseConfig) error {
	var err error

	dsn := cfg.DSN()
	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Test connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	log.Println("Database initialized successfully")
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return DB
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
