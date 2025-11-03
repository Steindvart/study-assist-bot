package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"study-assist-tgbot/internal/config"
	"study-assist-tgbot/internal/i18n"
	"study-assist-tgbot/internal/telegrambot"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	i18nService, err := i18n.NewService(i18n.Config{
		DefaultLanguage:  "ru",                 // Основной язык
		FallbackLanguage: "en",                 // Fallback язык
		SupportedLangs:   []string{"ru", "en"}, // Поддерживаемые языки
		LocalesPath:      "",                   // Пустая строка = используем embedded файлы
	})
	if err != nil {
		log.Fatalf("Failed to initialize i18n service: %v", err)
	}

	// Создаём бота с i18n
	bot, err := telegrambot.NewBot(telegrambot.Config{
		Token:       cfg.TelegramToken,
		I18nService: i18nService,
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Println("Bot is starting with i18n support...")

	if err := bot.Start(ctx); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	log.Println("Bot stopped gracefully")
}
