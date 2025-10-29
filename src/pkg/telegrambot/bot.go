package telegrambot

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	api *bot.Bot
}

func NewBot(token string) (*Bot, error) {
	opts := []bot.Option{
		bot.WithMiddlewares(logMiddleware),
		bot.WithDefaultHandler(echoHandler),
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

	b.api.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeExact, echoHandler)
	b.api.Start(ctx)

	return nil
}

func logMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			log.Printf("Received message from user %d in chat %d: %s",
				update.Message.From.ID,
				update.Message.Chat.ID,
				update.Message.Text,
			)
		}
		next(ctx, b, update)
	}
}

func echoHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Echo: " + update.Message.Text,
	})
}
