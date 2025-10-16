package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

// Localization holds the localized messages for different languages.
type Localization struct {
	mu         sync.RWMutex
	messages   map[string]map[string]string
	defaultLang string
}

// NewLocalization initializes a new Localization instance.
func NewLocalization(defaultLang string) *Localization {
	return &Localization{
		messages:   make(map[string]map[string]string),
		defaultLang: defaultLang,
	}
}

// LoadMessages loads localized messages from JSON files.
func (l *Localization) LoadMessages(lang string) error {
	filePath := "locales/" + lang + ".json"
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	var messages map[string]string
	if err := json.Unmarshal(data, &messages); err != nil {
		return err
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	l.messages[lang] = messages
	return nil
}

// GetMessage retrieves a localized message by key and language.
func (l *Localization) GetMessage(lang, key string) string {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if messages, ok := l.messages[lang]; ok {
		if msg, exists := messages[key]; exists {
			return msg
		}
	}
	// Fallback to default language if the message is not found
	if messages, ok := l.messages[l.defaultLang]; ok {
		return messages[key]
	}
	return ""
}

// LoadAllMessages loads messages for all supported languages.
func (l *Localization) LoadAllMessages(languages []string) error {
	for _, lang := range languages {
		if err := l.LoadMessages(lang); err != nil {
			return err
		}
	}
	return nil
}