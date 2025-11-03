package localization

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: Config{
				DefaultLanguage:  "ru",
				FallbackLanguage: "en",
				SupportedLangs:   []string{"ru", "en"},
				LocalesPath:      "testdata",
			},
			wantErr: false,
		},
		{
			name: "invalid default language",
			config: Config{
				DefaultLanguage:  "invalid-lang",
				FallbackLanguage: "en",
				SupportedLangs:   []string{"en"},
			},
			wantErr: true,
		},
		{
			name: "invalid fallback language",
			config: Config{
				DefaultLanguage:  "en",
				FallbackLanguage: "invalid-lang",
				SupportedLangs:   []string{"en"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, err := NewService(tt.config)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, service)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, service)
			}
		})
	}
}

func TestService_GetLocalizer(t *testing.T) {
	// Создаём тестовые переводы
	setupTestLocales(t)

	service, err := NewService(Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	tests := []struct {
		name     string
		langCode string
		wantNil  bool
	}{
		{
			name:     "existing language",
			langCode: "en",
			wantNil:  false,
		},
		{
			name:     "existing language ru",
			langCode: "ru",
			wantNil:  false,
		},
		{
			name:     "non-existing language returns fallback",
			langCode: "fr",
			wantNil:  false, // должен вернуть fallback localizer
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localizer := service.GetLocalizer(tt.langCode)

			if tt.wantNil {
				assert.Nil(t, localizer)
			} else {
				assert.NotNil(t, localizer)
			}
		})
	}
}

func TestService_LocalizeSimple(t *testing.T) {
	setupTestLocales(t)

	service, err := NewService(Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	tests := []struct {
		name      string
		langCode  string
		messageID string
		expected  string
	}{
		{
			name:      "english translation",
			langCode:  "en",
			messageID: "welcome_message",
			expected:  "Welcome to the Study Assist Bot! How can I help you today?",
		},
		{
			name:      "russian translation",
			langCode:  "ru",
			messageID: "welcome_message",
			expected:  "Добро пожаловать в Study Assist Bot! Как я могу помочь вам сегодня?",
		},
		{
			name:      "non-existing message returns messageID",
			langCode:  "en",
			messageID: "non_existing_message",
			expected:  "non_existing_message",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.LocalizeSimple(tt.langCode, tt.messageID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestService_Localize(t *testing.T) {
	setupTestLocales(t)

	service, err := NewService(Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	tests := []struct {
		name         string
		langCode     string
		messageID    string
		templateData map[string]interface{}
		expected     string
	}{
		{
			name:      "english with template data",
			langCode:  "en",
			messageID: "test_completed",
			templateData: map[string]interface{}{
				"Score": 8,
				"Total": 10,
			},
			expected: "Test completed! Your score: 8 out of 10.",
		},
		{
			name:      "russian with template data",
			langCode:  "ru",
			messageID: "test_completed",
			templateData: map[string]interface{}{
				"Score": 8,
				"Total": 10,
			},
			expected: "Тест завершён! Ваш результат: 8 из 10.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.Localize(tt.langCode, tt.messageID, tt.templateData)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestService_LocalizePlural(t *testing.T) {
	setupTestLocales(t)

	service, err := NewService(Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	tests := []struct {
		name      string
		langCode  string
		messageID string
		count     int
		expected  string
	}{
		{
			name:      "english singular",
			langCode:  "en",
			messageID: "tests_passed_plural",
			count:     1,
			expected:  "1 test passed",
		},
		{
			name:      "english plural",
			langCode:  "en",
			messageID: "tests_passed_plural",
			count:     5,
			expected:  "5 tests passed",
		},
		{
			name:      "russian singular",
			langCode:  "ru",
			messageID: "tests_passed_plural",
			count:     1,
			expected:  "1 тест пройден",
		},
		{
			name:      "russian few",
			langCode:  "ru",
			messageID: "tests_passed_plural",
			count:     2,
			expected:  "2 теста пройдено",
		},
		{
			name:      "russian many",
			langCode:  "ru",
			messageID: "tests_passed_plural",
			count:     5,
			expected:  "5 тестов пройдено",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.LocalizePlural(tt.langCode, tt.messageID, tt.count, nil)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestContext(t *testing.T) {
	setupTestLocales(t)

	service, err := NewService(Config{
		DefaultLanguage:  "en",
		FallbackLanguage: "en",
		SupportedLangs:   []string{"en", "ru"},
		LocalesPath:      "testdata",
	})
	require.NoError(t, err)

	// Тест добавления в контекст
	localizer := service.GetLocalizer("en")
	ctx := WithLocalizer(context.Background(), localizer)

	// Тест извлечения из контекста
	retrievedLocalizer, ok := FromContext(ctx)
	assert.True(t, ok)
	assert.Equal(t, localizer, retrievedLocalizer)

	// Тест с пустым контекстом
	emptyLocalizer, ok := FromContext(context.Background())
	assert.False(t, ok)
	assert.Nil(t, emptyLocalizer)

	// Тест MustFromContext с валидным контекстом
	mustLocalizer := MustFromContext(ctx)
	assert.Equal(t, localizer, mustLocalizer)

	// Тест MustFromContext с пустым контекстом (должен panic)
	assert.Panics(t, func() {
		MustFromContext(context.Background())
	})
}

func setupTestLocales(t *testing.T) {
	// Здесь нужно создать тестовые файлы локализации
	// В реальном проекте можно использовать embedded файлы
	// или создавать временные файлы для тестов
}
