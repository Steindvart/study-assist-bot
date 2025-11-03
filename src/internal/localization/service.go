package localization

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Embed локальные файлы в бинарник для production
//
//go:embed locales/*.json
var localesFS embed.FS

// Service управляет локализацией в приложении
type Service struct {
	bundle       *i18n.Bundle
	localizers   map[string]*i18n.Localizer
	defaultLang  language.Tag
	fallbackLang language.Tag
}

// Config для настройки i18n сервиса
type Config struct {
	DefaultLanguage  string   // "ru"
	FallbackLanguage string   // "en"
	SupportedLangs   []string // ["ru", "en"]
	LocalesPath      string   // путь к файлам локализации
}

// NewService создаёт новый i18n сервис с конфигурацией
func NewService(cfg Config) (*Service, error) {
	// Парсим языковые теги
	defaultLang, err := language.Parse(cfg.DefaultLanguage)
	if err != nil {
		return nil, fmt.Errorf("invalid default language %q: %w", cfg.DefaultLanguage, err)
	}

	fallbackLang, err := language.Parse(cfg.FallbackLanguage)
	if err != nil {
		return nil, fmt.Errorf("invalid fallback language %q: %w", cfg.FallbackLanguage, err)
	}

	// Создаём bundle с базовым языком
	bundle := i18n.NewBundle(defaultLang)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal) // Поддержка JSON

	service := &Service{
		bundle:       bundle,
		localizers:   make(map[string]*i18n.Localizer),
		defaultLang:  defaultLang,
		fallbackLang: fallbackLang,
	}

	// Загружаем переводы для всех поддерживаемых языков
	for _, lang := range cfg.SupportedLangs {
		if err := service.loadLanguage(lang, cfg.LocalesPath); err != nil {
			return nil, fmt.Errorf("failed to load language %s: %w", lang, err)
		}
	}

	return service, nil
}

// loadLanguage загружает переводы для конкретного языка
func (s *Service) loadLanguage(langCode, localesPath string) error {
	var filePath string
	if localesPath == "" {
		// Используем embedded файлы
		filePath = fmt.Sprintf("locales/%s.json", langCode)
		data, err := localesFS.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read embedded locale file %s: %w", filePath, err)
		}

		// Парсим и загружаем в bundle
		if _, err := s.bundle.ParseMessageFileBytes(data, filePath); err != nil {
			return fmt.Errorf("failed to parse locale file %s: %w", filePath, err)
		}
	} else {
		// Загружаем из файловой системы
		filePath = fmt.Sprintf("%s/%s.json", localesPath, langCode)
		if _, err := s.bundle.LoadMessageFile(filePath); err != nil {
			return fmt.Errorf("failed to load locale file %s: %w", filePath, err)
		}
	}

	// Создаём локализатор для языка
	langTag, err := language.Parse(langCode)
	if err != nil {
		return fmt.Errorf("invalid language code %s: %w", langCode, err)
	}

	s.localizers[langCode] = i18n.NewLocalizer(s.bundle, langTag.String(), s.fallbackLang.String())
	return nil
}

// GetLocalizer возвращает локализатор для языка пользователя
func (s *Service) GetLocalizer(langCode string) *i18n.Localizer {
	if localizer, exists := s.localizers[langCode]; exists {
		return localizer
	}

	// Возвращаем fallback локализатор
	return s.localizers[s.fallbackLang.String()]
}

// Localize переводит сообщение с параметрами
func (s *Service) Localize(langCode, messageID string, templateData map[string]interface{}) string {
	localizer := s.GetLocalizer(langCode)

	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})

	if err != nil {
		// В случае ошибки возвращаем messageID как fallback
		return messageID
	}

	return message
}

// LocalizeSimple переводит простое сообщение без параметров
func (s *Service) LocalizeSimple(langCode, messageID string) string {
	return s.Localize(langCode, messageID, nil)
}

// LocalizePlural переводит сообщения с плюрализацией
func (s *Service) LocalizePlural(langCode, messageID string, count int, templateData map[string]interface{}) string {
	localizer := s.GetLocalizer(langCode)

	if templateData == nil {
		templateData = make(map[string]interface{})
	}
	templateData["Count"] = count

	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		PluralCount:  count,
		TemplateData: templateData,
	})

	if err != nil {
		return messageID
	}

	return message
}

// SupportedLanguages возвращает список поддерживаемых языков
func (s *Service) SupportedLanguages() []string {
	langs := make([]string, 0, len(s.localizers))
	for lang := range s.localizers {
		langs = append(langs, lang)
	}
	return langs
}

// Ключ для хранения локализатора в контексте
type localizerKey struct{}

// WithLocalizer добавляет локализатор в контекст
func WithLocalizer(ctx context.Context, localizer *i18n.Localizer) context.Context {
	return context.WithValue(ctx, localizerKey{}, localizer)
}

// FromContext извлекает локализатор из контекста
func FromContext(ctx context.Context) (*i18n.Localizer, bool) {
	localizer, ok := ctx.Value(localizerKey{}).(*i18n.Localizer)
	return localizer, ok
}

// MustFromContext извлекает локализатор из контекста или panic
func MustFromContext(ctx context.Context) *i18n.Localizer {
	localizer, ok := FromContext(ctx)
	if !ok {
		panic("localizer not found in context")
	}
	return localizer
}
