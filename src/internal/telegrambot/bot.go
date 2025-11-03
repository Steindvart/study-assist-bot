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
	api         *bot.Bot
	i18nService *localization.Service
}

type Config struct {
	Token       string
	I18nService *localization.Service
}

func NewBot(cfg Config) (*Bot, error) {
	i18nMiddleware := middlewares.NewLocalizationMiddleware(cfg.I18nService)

	opts := []bot.Option{
		bot.WithMiddlewares(i18nMiddleware.Handler, middlewares.LogMessageWithText),

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
		api:         b,
		i18nService: cfg.I18nService,
	}, nil
}

// Start запускает бота в polling режиме
func (b *Bot) Start(ctx context.Context) error {
	log.Println("Starting Telegram bot...")
	log.Printf("Supported languages: %v", b.i18nService.SupportedLanguages())

	b.api.Start(ctx)

	return nil
}
