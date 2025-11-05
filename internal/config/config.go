package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"study-assist-tgbot/internal/database"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	Database      database.Config
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	dbConfig, err := loadDatabaseConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %w", err)
	}

	return &Config{
		TelegramToken: token,
		Database:      dbConfig,
	}, nil
}

func loadDatabaseConfig() (database.Config, error) {
	port, err := strconv.Atoi(getEnvOrDefault("DB_PORT", "5432"))
	if err != nil {
		return database.Config{}, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	maxOpenConns, err := strconv.Atoi(getEnvOrDefault("DB_MAX_OPEN_CONNS", "10"))
	if err != nil {
		return database.Config{}, fmt.Errorf("invalid DB_MAX_OPEN_CONNS: %w", err)
	}

	maxIdleConns, err := strconv.Atoi(getEnvOrDefault("DB_MAX_IDLE_CONNS", "5"))
	if err != nil {
		return database.Config{}, fmt.Errorf("invalid DB_MAX_IDLE_CONNS: %w", err)
	}

	maxLifetime, err := time.ParseDuration(getEnvOrDefault("DB_MAX_LIFETIME", "1h"))
	if err != nil {
		return database.Config{}, fmt.Errorf("invalid DB_MAX_LIFETIME: %w", err)
	}

	return database.Config{
		Host:         getEnvOrDefault("DB_HOST", "localhost"),
		Port:         port,
		User:         getEnvOrDefault("DB_USER", "postgres"),
		Password:     getEnvOrDefault("DB_PASSWORD", "postgres"),
		Database:     getEnvOrDefault("DB_NAME", "study_assist"),
		SSLMode:      getEnvOrDefault("DB_SSL_MODE", "disable"),
		MaxOpenConns: maxOpenConns,
		MaxIdleConns: maxIdleConns,
		MaxLifetime:  maxLifetime,
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
