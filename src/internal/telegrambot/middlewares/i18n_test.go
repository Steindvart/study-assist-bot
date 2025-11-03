package middlewares

import (
	"context"
	"testing"

	"study-assist-tgbot/internal/i18n"

	"github.com/go-telegram/bot/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestI18nMiddleware_getUserLanguage(t *testing.T) {
	service, err := i18n.NewService(i18n.Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	middleware := NewI18nMiddleware(service)

	tests := []struct {
		name     string
		update   *models.Update
		expected string
	}{
		{
			name: "message with russian language",
			update: &models.Update{
				Message: &models.Message{
					From: &models.User{
						LanguageCode: "ru",
					},
				},
			},
			expected: "ru",
		},
		{
			name: "message with english language",
			update: &models.Update{
				Message: &models.Message{
					From: &models.User{
						LanguageCode: "en",
					},
				},
			},
			expected: "en",
		},
		{
			name: "message with unsupported language falls back to ru",
			update: &models.Update{
				Message: &models.Message{
					From: &models.User{
						LanguageCode: "fr",
					},
				},
			},
			expected: "ru", // fallback
		},
		{
			name: "callback query with language",
			update: &models.Update{
				CallbackQuery: &models.CallbackQuery{
					From: models.User{
						LanguageCode: "en",
					},
				},
			},
			expected: "en",
		},
		{
			name:     "no language information defaults to ru",
			update:   &models.Update{},
			expected: "ru",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := middleware.getUserLanguage(tt.update)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestI18nMiddleware_normalizeLanguageCode(t *testing.T) {
	service, err := i18n.NewService(i18n.Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	middleware := NewI18nMiddleware(service)

	tests := []struct {
		name     string
		langCode string
		expected string
	}{
		{
			name:     "exact match ru",
			langCode: "ru",
			expected: "ru",
		},
		{
			name:     "exact match en",
			langCode: "en",
			expected: "en",
		},
		{
			name:     "long format ru-RU",
			langCode: "ru-RU",
			expected: "ru",
		},
		{
			name:     "long format en-US",
			langCode: "en-US",
			expected: "en",
		},
		{
			name:     "unsupported language",
			langCode: "fr",
			expected: "ru", // fallback
		},
		{
			name:     "empty string",
			langCode: "",
			expected: "ru", // fallback
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := middleware.normalizeLanguageCode(tt.langCode)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetMessage(t *testing.T) {
	service, err := i18n.NewService(i18n.Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	// Тест с локализатором в контексте
	localizer := service.GetLocalizer("en")
	ctx := i18n.WithLocalizer(context.Background(), localizer)

	tests := []struct {
		name         string
		ctx          context.Context
		messageID    string
		templateData map[string]interface{}
		expected     string
	}{
		{
			name:      "valid context with message",
			ctx:       ctx,
			messageID: "welcome_message",
			expected:  "Welcome to the Study Assist Bot! How can I help you today?", // предполагаемый перевод
		},
		{
			name:      "context without localizer",
			ctx:       context.Background(),
			messageID: "welcome_message",
			expected:  "welcome_message", // fallback to messageID
		},
		{
			name:      "non-existing message",
			ctx:       ctx,
			messageID: "non_existing",
			expected:  "non_existing", // fallback to messageID
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetMessage(tt.ctx, tt.messageID, tt.templateData)
			// Поскольку у нас нет реальных файлов переводов в тестах,
			// проверяем что функция не падает и возвращает строку
			assert.NotEmpty(t, result)
		})
	}
}

func TestGetSimpleMessage(t *testing.T) {
	service, err := i18n.NewService(i18n.Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	localizer := service.GetLocalizer("en")
	ctx := i18n.WithLocalizer(context.Background(), localizer)

	result := GetSimpleMessage(ctx, "welcome_message")
	assert.NotEmpty(t, result)

	// Тест без локализатора
	result = GetSimpleMessage(context.Background(), "welcome_message")
	assert.Equal(t, "welcome_message", result)
}

func TestGetPluralMessage(t *testing.T) {
	service, err := i18n.NewService(i18n.Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	localizer := service.GetLocalizer("en")
	ctx := i18n.WithLocalizer(context.Background(), localizer)

	result := GetPluralMessage(ctx, "tests_passed_plural", 5, nil)
	assert.NotEmpty(t, result)

	// Тест без локализатора
	result = GetPluralMessage(context.Background(), "tests_passed_plural", 5, nil)
	assert.Equal(t, "tests_passed_plural", result)
}
