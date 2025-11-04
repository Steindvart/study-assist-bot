package telegrambot

import (
	"context"
	"log"
	"study-assist-tgbot/internal/localization"
	"study-assist-tgbot/internal/telegrambot/handlers"
	"study-assist-tgbot/internal/telegrambot/middlewares"

	"github.com/go-telegram/bot"
)

type Bot struct {
	api          *bot.Bot
	localization *localization.Service
}

type Config struct {
	Token               string
	LocalizationService *localization.Service
}

func NewBot(cfg Config) (*Bot, error) {
	localizationMiddleware := middlewares.NewLocalization(cfg.LocalizationService)

	opts := []bot.Option{
		bot.WithMiddlewares(localizationMiddleware.Handler, middlewares.LogMessageWithText),

		bot.WithMessageTextHandler("start", bot.MatchTypeCommand, handlers.CommandStart),
		bot.WithMessageTextHandler("help", bot.MatchTypeCommand, handlers.CommandHelp),
		bot.WithMessageTextHandler("lang", bot.MatchTypeCommand, handlers.CommandLanguage),
		bot.WithMessageTextHandler("test", bot.MatchTypeCommand, handlers.CommandTest),

		bot.WithCallbackQueryDataHandler("set_lang_", bot.MatchTypePrefix, handlers.HandleLanguageCallback),

		bot.WithDefaultHandler(handlers.Echo),
	}

	b, err := bot.New(cfg.Token, opts...)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:          b,
		localization: cfg.LocalizationService,
	}, nil
}

// Start запускает бота в polling режиме
func (b *Bot) Start(ctx context.Context) error {
	log.Println("Starting Telegram bot...")
	log.Printf("Supported languages: %v", b.localization.SupportedLanguages())

	b.api.Start(ctx)

	return nil
}
