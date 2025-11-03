package telegrambot

import (
	"context"
	"log"
	"study-assist-tgbot/internal/telegrambot/handlers"
	"study-assist-tgbot/internal/telegrambot/middlewares"

	"github.com/go-telegram/bot"
)

type Bot struct {
	api *bot.Bot
}

func NewBot(token string) (*Bot, error) {
	opts := []bot.Option{
		bot.WithMiddlewares(middlewares.LogMessage),
		bot.WithDefaultHandler(handlers.Echo),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api: b,
	}, nil
}

// Start запускает бота в polling режиме
func (b *Bot) Start(ctx context.Context) error {
	log.Println("Starting Telegram bot...")

	b.api.Start(ctx)

	return nil
}
