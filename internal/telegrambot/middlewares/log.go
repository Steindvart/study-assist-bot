package middlewares

import (
	"context"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func LogMessageWithText(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			if update.Message.Text == "" {
				return
			}

			log.Printf("Received message from user %d (%s) in chat %d (%s): %s",
				update.Message.From.ID,
				update.Message.From.Username,
				update.Message.Chat.ID,
				update.Message.Chat.Username,
				update.Message.Text,
			)
		}
		next(ctx, b, update)
	}
}
