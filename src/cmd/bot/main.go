package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"study-assist-tgbot/internal/config"
	"study-assist-tgbot/internal/localization"
	"study-assist-tgbot/internal/telegrambot"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	localizationService, err := localization.NewService(localization.Config{
		DefaultLanguage:  "ru",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"ru", "en"},
	})

	if err != nil {
		log.Fatalf("Failed to initialize localization service: %v", err)
	}

	bot, err := telegrambot.NewBot(telegrambot.Config{
		Token:               cfg.TelegramToken,
		LocalizationService: localizationService,
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := bot.Start(ctx); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	log.Println("Bot stopped gracefully")
}
