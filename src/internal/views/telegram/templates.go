package telegram

import (
	"fmt"
	"strings"
)

// MessageTemplate holds the structure for dynamic message templates.
type MessageTemplate struct {
	Title   string
	Content string
}

// GenerateSectionMessage creates a message for a specific section.
func GenerateSectionMessage(sectionTitle string, topics []string) MessageTemplate {
	content := fmt.Sprintf("Welcome to the section: %s\nTopics covered:\n- %s", sectionTitle, strings.Join(topics, "\n- "))
	return MessageTemplate{
		Title:   sectionTitle,
		Content: content,
	}
}

// GenerateTestResultMessage creates a message summarizing the test results.
func GenerateTestResultMessage(score int, total int) MessageTemplate {
	content := fmt.Sprintf("You scored %d out of %d.\n", score, total)
	if score == total {
		content += "Excellent work! You answered all questions correctly."
	} else {
		content += "Keep practicing to improve your score!"
	}
	return MessageTemplate{
		Title:   "Test Results",
		Content: content,
	}
}

// GenerateExplanationMessage creates a message providing an explanation for a question.
func GenerateExplanationMessage(question string, explanation string) MessageTemplate {
	content := fmt.Sprintf("Question: %s\nExplanation: %s", question, explanation)
	return MessageTemplate{
		Title:   "Explanation",
		Content: content,
	}
}

// GenerateLocalizationMessage creates a localized message based on the user's language preference.
func GenerateLocalizationMessage(key string, lang string) MessageTemplate {
	// This function would typically fetch the localized message from a map or database.
	// For simplicity, we will return a placeholder message.
	content := fmt.Sprintf("Localized message for key: %s in language: %s", key, lang)
	return MessageTemplate{
		Title:   "Localization",
		Content: content,
	}
}