package localization

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Embed локальные файлы в бинарник для production
//
//go:embed locales/*.json
var localesFS embed.FS

type Service struct {
	bundle       *i18n.Bundle
	localizers   map[string]*i18n.Localizer
	defaultLang  language.Tag
	fallbackLang language.Tag
}

type Config struct {
	DefaultLanguage  string
	FallbackLanguage string
	SupportedLangs   []string
}

func NewService(cfg Config) (*Service, error) {
	defaultLang, err := language.Parse(cfg.DefaultLanguage)
	if err != nil {
		return nil, fmt.Errorf("invalid default language %q: %w", cfg.DefaultLanguage, err)
	}

	fallbackLang, err := language.Parse(cfg.FallbackLanguage)
	if err != nil {
		return nil, fmt.Errorf("invalid fallback language %q: %w", cfg.FallbackLanguage, err)
	}

	bundle := i18n.NewBundle(defaultLang)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	service := &Service{
		bundle:       bundle,
		localizers:   make(map[string]*i18n.Localizer),
		defaultLang:  defaultLang,
		fallbackLang: fallbackLang,
	}

	for _, lang := range cfg.SupportedLangs {
		if err := service.loadLanguage(lang); err != nil {
			return nil, fmt.Errorf("failed to load language %s: %w", lang, err)
		}
	}

	return service, nil
}

func (s *Service) loadLanguage(langCode string) error {
	filePath := fmt.Sprintf("locales/%s.json", langCode)
	data, err := localesFS.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read embedded locale file %s: %w", filePath, err)
	}

	if _, err := s.bundle.ParseMessageFileBytes(data, filePath); err != nil {
		return fmt.Errorf("failed to parse locale file %s: %w", filePath, err)
	}

	langTag, err := language.Parse(langCode)
	if err != nil {
		return fmt.Errorf("invalid language code %s: %w", langCode, err)
	}

	s.localizers[langCode] = i18n.NewLocalizer(s.bundle, langTag.String(), s.fallbackLang.String())
	return nil
}

func (s *Service) GetLocalizer(langCode string) *i18n.Localizer {
	if localizer, exists := s.localizers[langCode]; exists {
		return localizer
	}

	return s.localizers[s.fallbackLang.String()]
}

func (s *Service) Localize(langCode, messageID string, templateData map[string]any) string {
	localizer := s.GetLocalizer(langCode)

	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})

	if err != nil {
		log.Print("Localization error: ", err)
		return messageID
	}

	return message
}

func (s *Service) LocalizeSimple(langCode, messageID string) string {
	return s.Localize(langCode, messageID, nil)
}

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
