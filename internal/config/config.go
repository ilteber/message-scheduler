package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Database  DatabaseConfig
	Redis     RedisConfig
	Webhook   WebhookConfig
	Scheduler SchedulerConfig
	Server    ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type WebhookConfig struct {
	URL     string
	AuthKey string
}

type SchedulerConfig struct {
	IntervalSeconds  int
	MessagesPerBatch int
}

type ServerConfig struct {
	Port string
}

func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "insider"),
			Password: getEnv("DB_PASSWORD", "insider123"),
			DBName:   getEnv("DB_NAME", "insider_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnv("REDIS_PORT", "6379"),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
		},
		Webhook: WebhookConfig{
			URL:     getEnv("WEBHOOK_URL", "https://webhook.site/9e83d1aa-0046-4f29-a2d8-deb68ce93ec0"),
			AuthKey: getEnv("WEBHOOK_AUTH_KEY", "INS.me1x9uMcyYGlhKKQVPoc.bO3j9aZwRTOcA2Ywo"),
		},
		Scheduler: SchedulerConfig{
			IntervalSeconds:  getEnvAsInt("SCHEDULER_INTERVAL_SECONDS", 120),
			MessagesPerBatch: getEnvAsInt("MESSAGES_PER_BATCH", 2),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

func (c *RedisConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
