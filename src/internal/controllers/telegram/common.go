package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot/v2"
)

// SendMessage sends a message to a specific chat ID with the given text.
func SendMessage(ctx context.Context, b *bot.Bot, chatID int64, text string) error {
	_, err := b.Send(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	if err != nil {
		log.Printf("Failed to send message to chat %d: %v", chatID, err)
		return err
	}
	return nil
}

// HandleError logs and sends an error message to the user.
func HandleError(ctx context.Context, b *bot.Bot, chatID int64, err error) {
	log.Printf("Error occurred: %v", err)
	errMsg := fmt.Sprintf("An error occurred: %v", err)
	if sendErr := SendMessage(ctx, b, chatID, errMsg); sendErr != nil {
		log.Printf("Failed to send error message to chat %d: %v", chatID, sendErr)
	}
}

// GetUserLanguage retrieves the user's preferred language from the context or defaults to English.
func GetUserLanguage(ctx context.Context) string {
	lang, ok := ctx.Value("language").(string)
	if !ok {
		return "en" // Default language
	}
	return lang
}