package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"study-assist-bot-go/internal/config"
	"study-assist-bot-go/pkg/telegrambot"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Создаем экземпляр бота
	bot, err := telegrambot.NewBot(cfg.TelegramToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Создаем контекст с возможностью отмены
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log.Println("Bot is starting...")

	// Запускаем бота
	if err := bot.Start(ctx); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	log.Println("Bot stopped gracefully")
}
