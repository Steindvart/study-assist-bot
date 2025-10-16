package telegram

import (
	"github.com/go-telegram/bot"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Keyboard represents a custom keyboard layout for the Telegram bot.
type Keyboard struct {
	Buttons [][]string
}

// NewMainMenuKeyboard creates a keyboard for the main menu.
func NewMainMenuKeyboard() *Keyboard {
	return &Keyboard{
		Buttons: [][]string{
			{"Start Test", "View Statistics"},
			{"Help", "Settings"},
		},
	}
}

// NewSectionKeyboard creates a keyboard for selecting sections.
func NewSectionKeyboard(sections []string) *Keyboard {
	buttons := make([][]string, len(sections))
	for i, section := range sections {
		buttons[i] = []string{section}
	}
	return &Keyboard{Buttons: buttons}
}

// NewLanguageKeyboard creates a keyboard for language selection.
func NewLanguageKeyboard() *Keyboard {
	return &Keyboard{
		Buttons: [][]string{
			{"English", "Русский"},
		},
	}
}

// SendKeyboard sends a custom keyboard to the user.
func SendKeyboard(chatID int64, keyboard *Keyboard, bot *bot.Bot) error {
	replyMarkup := bot.NewReplyKeyboardMarkup(keyboard.Buttons)
	_, err := bot.SendMessage(chatID, "Please choose an option:", replyMarkup)
	return err
}

// LocalizedKeyboard creates a localized keyboard based on the user's language preference.
func LocalizedKeyboard(lang language.Tag) *Keyboard {
	p := message.NewPrinter(lang)
	if lang == language.Russian {
		return &Keyboard{
			Buttons: [][]string{
				{p.Sprintf("Начать тест"), p.Sprintf("Посмотреть статистику")},
				{p.Sprintf("Помощь"), p.Sprintf("Настройки")},
			},
		}
	}
	return NewMainMenuKeyboard()
}