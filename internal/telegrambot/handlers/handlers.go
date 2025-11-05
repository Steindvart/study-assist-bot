package handlers

import (
	"context"
	"fmt"
	"study-assist-tgbot/internal/localization"
	"study-assist-tgbot/internal/repositories"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Echo(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	echoText := fmt.Sprintf("Echo: %s", update.Message.Text)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   echoText,
	})
}

func CommandStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	welcomeText := localization.GetSimpleText(ctx, "welcome_message")

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   welcomeText,
	})
}

func CommandHelp(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	helpText := localization.GetSimpleText(ctx, "help_message")

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   helpText,
	})
}

func CommandLanguage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

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

	selectText := localization.GetSimpleText(ctx, "language_selection")

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        selectText,
		ReplyMarkup: keyboard,
	})
}

func HandleLanguageCallback(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.CallbackQuery == nil {
		return
	}

	var langCode, langName string

	switch update.CallbackQuery.Data {
	case "set_lang_ru":
		langCode = "ru"
		langName = "русский"
	case "set_lang_en":
		langCode = "en"
		langName = "English"
	default:
		return
	}

	userRepo, ok := repositories.GetUserRepository(ctx)
	if ok {
		telegramID := update.CallbackQuery.From.ID
		_, err := userRepo.UpsertLanguage(ctx, telegramID, langCode)
		if err != nil {
			fmt.Printf("Failed to save language preference for user %d: %v\n", telegramID, err)
		}
	}

	confirmText := localization.GetText(ctx, "language_changed", map[string]interface{}{
		"Language": langName,
	})

	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		Text:            confirmText,
		ShowAlert:       false,
	})

	b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
		MessageID: update.CallbackQuery.Message.Message.ID,
		Text:      confirmText,
	})
}

func CommandTest(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	// Пример использования плюрализации
	testCount := 5
	testText := localization.GetPluralText(ctx, "tests_passed_plural", testCount, nil)

	// Пример использования с параметрами
	resultText := localization.GetText(ctx, "test_completed", map[string]interface{}{
		"Score": 8,
		"Total": 10,
	})

	responseText := fmt.Sprintf("%s\n%s", testText, resultText)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   responseText,
	})
}
