package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"study-assist-tgbot/internal/config"
	"study-assist-tgbot/internal/database"
	"study-assist-tgbot/internal/localization"
	"study-assist-tgbot/internal/models"
	"study-assist-tgbot/internal/repositories"
	"study-assist-tgbot/internal/telegrambot"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	if err := db.DB().AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to run auto-migration: %v", err)
	}

	log.Println("Database auto-migration completed successfully")

	userRepository := repositories.NewUserRepository(db.DB())

	localizationService, err := localization.NewService(localization.ServiceConfig{
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
		UserRepository:      userRepository,
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
