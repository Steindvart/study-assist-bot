package localization

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.json
var localesFS embed.FS

type Service struct {
	bundle       *i18n.Bundle
	localizers   map[string]Localizer
	defaultLang  language.Tag
	fallbackLang language.Tag
}

type ServiceConfig struct {
	DefaultLanguage  string
	FallbackLanguage string
	SupportedLangs   []string
}

func NewService(cfg ServiceConfig) (*Service, error) {
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
		localizers:   make(map[string]Localizer, len(cfg.SupportedLangs)),
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

func (s *Service) GetLocalizer(langCode string) Localizer {
	if localizer, exists := s.localizers[langCode]; exists {
		return localizer
	}

	return s.localizers[s.fallbackLang.String()]
}

func (s *Service) SupportedLanguages() []string {
	langs := make([]string, 0, len(s.localizers))
	for lang := range s.localizers {
		langs = append(langs, lang)
	}
	return langs
}
