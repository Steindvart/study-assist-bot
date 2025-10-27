package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config содержит конфигурацию приложения
type Config struct {
	TelegramToken string
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() (*Config, error) {
	// Пытаемся загрузить .env файл (необязательно)
	_ = godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	return &Config{
		TelegramToken: token,
	}, nil
}
