package middlewares

import (
	"context"
	"study-assist-tgbot/internal/localization"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const defaultLanguageCode = "en"

type Localization struct {
	service *localization.Service
}

func NewLocalization(service *localization.Service) *Localization {
	return &Localization{
		service: service,
	}
}

func (m *Localization) Handler(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		// Определяем язык пользователя
		userLang := m.getUserLanguage(update)

		// Получаем соответствующий локализатор
		localizer := m.service.GetLocalizer(userLang)

		// Обогащаем контекст локализатором
		ctx = localization.WithLocalizer(ctx, localizer)

		next(ctx, b, update)
	}
}

func (m *Localization) getUserLanguage(update *models.Update) string {
	// Проверяем разные источники языка в порядке приоритета
	// 1. Из сообщения пользователя
	if update.Message != nil && update.Message.From != nil {
		if langCode := update.Message.From.LanguageCode; langCode != "" {
			return m.normalizeLanguageCode(langCode)
		}
	}

	// 2. Из callback query
	if update.CallbackQuery != nil && update.CallbackQuery.From.LanguageCode != "" {
		return m.normalizeLanguageCode(update.CallbackQuery.From.LanguageCode)
	}

	// 3. Из inline query
	if update.InlineQuery != nil && update.InlineQuery.From.LanguageCode != "" {
		return m.normalizeLanguageCode(update.InlineQuery.From.LanguageCode)
	}

	return defaultLanguageCode
}

func (m *Localization) normalizeLanguageCode(langCode string) string {
	supportedLangs := m.service.SupportedLanguages()

	for _, supported := range supportedLangs {
		if langCode == supported {
			return langCode
		}
	}

	if len(langCode) >= 2 {
		shortCode := langCode[:2]
		for _, supported := range supportedLangs {
			if shortCode == supported {
				return shortCode
			}
		}
	}

	return defaultLanguageCode
}
