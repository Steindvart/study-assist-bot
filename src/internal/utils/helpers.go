package utils

import (
	"regexp"
	"strings"
)

// SanitizeInput removes unwanted characters from user input.
func SanitizeInput(input string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9а-яА-ЯёЁ\s]`)
	return re.ReplaceAllString(input, "")
}

// ToTitleCase converts a string to title case.
func ToTitleCase(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, " ")
}

// IsValidEmail checks if the provided email address is valid.
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// FormatResponse formats the response message for the user.
func FormatResponse(message string) string {
	return strings.TrimSpace(message)
}