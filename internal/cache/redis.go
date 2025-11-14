package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ilteber/message-scheduler/internal/config"
)

var Client *redis.Client
var ctx = context.Background()

// MessageCacheData represents cached message data
type MessageCacheData struct {
	MessageID   string    `json:"message_id"`
	SentAt      time.Time `json:"sent_at"`
	PhoneNumber string    `json:"phone_number"`
	Content     string    `json:"content"`
}

// Initialize establishes Redis connection
func Initialize(cfg *config.RedisConfig) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Test connection
	if err := Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Println("Redis initialized successfully")
	return nil
}

// CacheMessageSent stores message send information in Redis
func CacheMessageSent(messageID string, phoneNumber string, content string, sentAt time.Time) error {
	if Client == nil {
		return fmt.Errorf("redis client not initialized")
	}

	data := MessageCacheData{
		MessageID:   messageID,
		SentAt:      sentAt,
		PhoneNumber: phoneNumber,
		Content:     content,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal cache data: %w", err)
	}

	// Store with key format: message:sent:{messageID}
	key := fmt.Sprintf("message:sent:%s", messageID)

	// Set with no expiration (you can add TTL if needed)
	if err := Client.Set(ctx, key, jsonData, 0).Err(); err != nil {
		return fmt.Errorf("failed to cache message: %w", err)
	}

	return nil
}

// GetCachedMessage retrieves cached message data
func GetCachedMessage(messageID string) (*MessageCacheData, error) {
	if Client == nil {
		return nil, fmt.Errorf("redis client not initialized")
	}

	key := fmt.Sprintf("message:sent:%s", messageID)

	val, err := Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("message not found in cache")
	} else if err != nil {
		return nil, fmt.Errorf("failed to get cached message: %w", err)
	}

	var data MessageCacheData
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cache data: %w", err)
	}

	return &data, nil
}

// Close closes the Redis connection
func Close() error {
	if Client != nil {
		return Client.Close()
	}
	return nil
}
