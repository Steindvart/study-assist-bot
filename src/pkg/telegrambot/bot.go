package telegrambot

import (
	"context"
	"log"
	"log/slog"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Bot представляет обертку над go-telegram/bot
type Bot struct {
	api     *bot.Bot
	handler func(ctx context.Context, b *bot.Bot, update *models.Update)
}

// NewBot создает новый экземпляр бота с переданным токеном
func NewBot(token string) (*Bot, error) {
	// Опции для создания бота
	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	// Создание экземпляра бота
	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:     b,
		handler: defaultHandler,
	}, nil
}

// SetHandler устанавливает пользовательский обработчик сообщений
func (b *Bot) SetHandler(handler func(ctx context.Context, bot *bot.Bot, update *models.Update)) {
	b.handler = handler
}

// Start запускает бота в polling режиме
func (b *Bot) Start(ctx context.Context) error {
	log.Println("Starting Telegram bot...")

	// Регистрируем наш обработчик
	b.api.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeExact, b.handler)

	// Запускаем polling
	b.api.Start(ctx)

	return nil
}

// SendMessage отправляет текстовое сообщение
func (b *Bot) SendMessage(ctx context.Context, chatID int64, text string) error {
	_, err := b.api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	})
	return err
}

// SendMessageWithKeyboard отправляет сообщение с клавиатурой
func (b *Bot) SendMessageWithKeyboard(ctx context.Context, chatID int64, text string, keyboard *models.InlineKeyboardMarkup) error {
	_, err := b.api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: keyboard,
	})
	return err
}

// GetBot возвращает внутренний экземпляр бота для расширенного использования
func (b *Bot) GetBot() *bot.Bot {
	return b.api
}

// defaultHandler - обработчик по умолчанию
func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	slog.Info("Received message",
		"user_id", update.Message.From.ID,
		"text", update.Message.Text,
		"chat_id", update.Message.Chat.ID,
	)
}
