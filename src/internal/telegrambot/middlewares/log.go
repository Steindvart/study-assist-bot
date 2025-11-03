package middlewares

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func LogMessage(next bot.HandlerFunc) bot.HandlerFunc {
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
