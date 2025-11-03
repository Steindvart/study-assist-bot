package middlewares

import (
	"context"
	"log"
	"study-assist-tgbot/internal/i18n"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
)

// I18nMiddleware обогащает контекст локализатором на основе языка пользователя
type I18nMiddleware struct {
	i18nService *i18n.Service
}

// NewI18nMiddleware создаёт новый middleware для i18n
func NewI18nMiddleware(i18nService *i18n.Service) *I18nMiddleware {
	return &I18nMiddleware{
		i18nService: i18nService,
	}
}

// Handler - основной обработчик middleware
func (m *I18nMiddleware) Handler(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		// Определяем язык пользователя
		userLang := m.getUserLanguage(update)

		// Получаем соответствующий локализатор
		localizer := m.i18nService.GetLocalizer(userLang)

		// Обогащаем контекст локализатором
		ctx = i18n.WithLocalizer(ctx, localizer)

		// Передаём управление следующему обработчику
		next(ctx, b, update)
	}
}

// getUserLanguage определяет язык пользователя из обновления Telegram
func (m *I18nMiddleware) getUserLanguage(update *models.Update) string {
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

	// 4. Fallback на русский (ваш дефолтный язык)
	return "ru"
}

// normalizeLanguageCode нормализует код языка к поддерживаемому формату
func (m *I18nMiddleware) normalizeLanguageCode(langCode string) string {
	// Проверяем поддерживаемые языки
	supportedLangs := m.i18nService.SupportedLanguages()

	// Точное совпадение
	for _, supported := range supportedLangs {
		if langCode == supported {
			return langCode
		}
	}

	// Проверяем первые два символа для языков типа "en-US" -> "en"
	if len(langCode) >= 2 {
		shortCode := langCode[:2]
		for _, supported := range supportedLangs {
			if shortCode == supported {
				return shortCode
			}
		}
	}

	// Специальная обработка для русского языка
	if langCode == "ru" || langCode == "ru-RU" {
		return "ru"
	}

	// Специальная обработка для английского языка
	if langCode == "en" || langCode[:2] == "en" {
		return "en"
	}

	// Fallback на русский язык по умолчанию
	return "ru"
}

// Helper функции для использования в handlers

// GetMessage получает переведённое сообщение из контекста
func GetMessage(ctx context.Context, messageID string, templateData map[string]interface{}) string {
	localizer, ok := i18n.FromContext(ctx)
	if !ok {
		// Fallback если локализатор не найден
		return messageID
	}

	config := &goi18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	}

	message, err := localizer.Localize(config)
	if err != nil {
		log.Fatal("Localization error: ", err)
		return messageID
	}

	return message
}

// GetSimpleMessage получает простое переведённое сообщение без параметров
func GetSimpleMessage(ctx context.Context, messageID string) string {
	return GetMessage(ctx, messageID, nil)
}

// GetPluralMessage получает переведённое сообщение с плюрализацией
func GetPluralMessage(ctx context.Context, messageID string, count int, templateData map[string]interface{}) string {
	localizer, ok := i18n.FromContext(ctx)
	if !ok {
		return messageID
	}

	if templateData == nil {
		templateData = make(map[string]interface{})
	}
	templateData["Count"] = count

	config := &goi18n.LocalizeConfig{
		MessageID:    messageID,
		PluralCount:  count,
		TemplateData: templateData,
	}

	message, err := localizer.Localize(config)
	if err != nil {
		return messageID
	}

	return message
}
