package handlers

import (
	"context"
	"fmt"
	"study-assist-tgbot/internal/telegrambot/middlewares"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Echo обработчик для эхо-сообщений с локализацией
func Echo(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	// Используем локализованное сообщение для эхо
	echoText := fmt.Sprintf("Echo: %s", update.Message.Text)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   echoText,
	})
}

// CommandStart обработчик команды /start с локализацией
func CommandStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	// Получаем локализованное приветствие
	welcomeText := middlewares.GetSimpleMessage(ctx, "welcome_message")

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   welcomeText,
	})
}

// CommandHelp обработчик команды /help с локализацией
func CommandHelp(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	// Получаем локализованную справку
	helpText := middlewares.GetSimpleMessage(ctx, "help_message")

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   helpText,
	})
}

// CommandLanguage обработчик команды /lang для смены языка
func CommandLanguage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	// Создаём inline клавиатуру для выбора языка
	keyboard := models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{
					Text:         "🇷🇺 Русский",
					CallbackData: "set_lang_ru",
				},
				{
					Text:         "🇺🇸 English",
					CallbackData: "set_lang_en",
				},
			},
		},
	}

	// Получаем локализованное сообщение выбора языка
	selectText := middlewares.GetSimpleMessage(ctx, "language_selection")

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        selectText,
		ReplyMarkup: keyboard,
	})
}

// HandleLanguageCallback обработчик callback для смены языка
func HandleLanguageCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.CallbackQuery == nil {
		return
	}

	var langName string

	switch update.CallbackQuery.Data {
	case "set_lang_ru":
		langName = "русский"
	case "set_lang_en":
		langName = "English"
	default:
		return
	}

	// Получаем локализованное сообщение подтверждения
	confirmText := middlewares.GetMessage(ctx, "language_changed", map[string]interface{}{
		"Language": langName,
	})

	// Отвечаем на callback query
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            confirmText,
		ShowAlert:       false,
	})

	// Обновляем сообщение
	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
		Text:      confirmText,
	})

	// В реальном приложении здесь нужно сохранить выбор языка в базу данных
	// для персистентности между сессиями
}

// CommandTest демонстрационный обработчик для тестирования с плюрализацией
func CommandTest(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	// Пример использования плюрализации
	testCount := 5
	testText := middlewares.GetPluralMessage(ctx, "tests_passed_plural", testCount, nil)

	// Пример использования с параметрами
	resultText := middlewares.GetMessage(ctx, "test_completed", map[string]interface{}{
		"Score": 8,
		"Total": 10,
	})

	responseText := fmt.Sprintf("%s\n%s", testText, resultText)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   responseText,
	})
}
