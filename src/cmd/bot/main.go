package main

import (
	"log"
	"os"

	"study-assist-bot-go/internal/config"
	"study-assist-bot-go/pkg/telegrambot"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the Telegram bot
	bot, err := telegrambot.NewBot(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Error initializing bot: %v", err)
	}

	// Start listening for updates
	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	for update := range updates {
		if update.Message == nil { // ignore non-message updates
			continue
		}

		// Process incoming messages
		if err := bot.ProcessMessage(update.Message); err != nil {
			log.Printf("Error processing message: %v", err)
		}
	}
}
