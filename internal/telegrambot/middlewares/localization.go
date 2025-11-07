package middlewares

import (
	"context"
	"log"
	"study-assist-tgbot/internal/localization"
	"study-assist-tgbot/internal/repositories"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const defaultLanguageCode = "en"

type Localization struct {
	service        *localization.Service
	userRepository repositories.UserRepository
}

func NewLocalization(service *localization.Service, userRepository repositories.UserRepository) *Localization {
	return &Localization{
		service:        service,
		userRepository: userRepository,
	}
}

func (m *Localization) Handler(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		userLang := m.getUserLanguage(ctx, update)
		localizer := m.service.GetLocalizer(userLang)
		ctx = localization.WithLocalizer(ctx, localizer)

		next(ctx, b, update)
	}
}

func (m *Localization) getUserLanguage(ctx context.Context, update *models.Update) string {
	telegramID := m.getTelegramID(update)
	if telegramID == 0 {
		return defaultLanguageCode
	}

	user, err := m.userRepository.GetByTelegramID(ctx, telegramID)
	if err != nil {
		log.Printf("Failed to get user from database: %v", err)
	} else if user != nil {
		return user.LanguageCode
	}

	telegramLang := m.getTelegramLanguage(update)
	if telegramLang != "" {
		normalizedLang := m.normalizeLanguageCode(telegramLang)

		_, err := m.userRepository.UpsertLanguage(ctx, telegramID, normalizedLang)
		if err != nil {
			log.Printf("Failed to upsert user language: %v", err)
		}

		return normalizedLang
	}

	return defaultLanguageCode
}

func (m *Localization) getTelegramID(update *models.Update) int64 {
	if update.Message != nil && update.Message.From != nil {
		return update.Message.From.ID
	}
	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}
	if update.InlineQuery != nil {
		return update.InlineQuery.From.ID
	}
	return 0
}

func (m *Localization) getTelegramLanguage(update *models.Update) string {
	if update.Message != nil && update.Message.From != nil {
		return update.Message.From.LanguageCode
	}

	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.LanguageCode
	}

	if update.InlineQuery != nil {
		return update.InlineQuery.From.LanguageCode
	}

	return ""
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
