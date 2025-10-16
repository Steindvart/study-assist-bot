package telegram

import (
	"context"
	"fmt"
	"log"

	"study-assist-bot-go/internal/models"
	"study-assist-bot-go/internal/repository"
	"study-assist-bot-go/internal/service"

	"github.com/go-telegram/bot"
)

// SectionController handles section-related operations for the Telegram bot.
type SectionController struct {
	repo repository.SectionRepository
}

// NewSectionController creates a new SectionController.
func NewSectionController(repo repository.SectionRepository) *SectionController {
	return &SectionController{repo: repo}
}

// HandleSectionCommand processes the /section command from the user.
func (sc *SectionController) HandleSectionCommand(ctx context.Context, update bot.Update) {
	sections, err := sc.repo.GetAllSections(ctx)
	if err != nil {
		log.Printf("Error fetching sections: %v", err)
		return
	}

	// Send sections to the user
	for _, section := range sections {
		message := fmt.Sprintf("Section: %s\nDescription: %s", section.Title, section.Description)
		if _, err := bot.SendMessage(update.Message.Chat.ID, message); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}
}

// GetSectionDetails fetches details for a specific section.
func (sc *SectionController) GetSectionDetails(ctx context.Context, sectionID int) (*models.Section, error) {
	section, err := sc.repo.GetSectionByID(ctx, sectionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get section details: %w", err)
	}
	return section, nil
}

// TrackUserResponse tracks user responses for a specific section.
func (sc *SectionController) TrackUserResponse(ctx context.Context, userID int, sectionID int, response string) error {
	// Logic to track user response
	return service.TrackResponse(ctx, userID, sectionID, response)
}

// ProvideExplanation provides an explanation for a question in a section.
func (sc *SectionController) ProvideExplanation(ctx context.Context, questionID int) (string, error) {
	explanation, err := service.GetExplanation(ctx, questionID)
	if err != nil {
		return "", fmt.Errorf("failed to get explanation: %w", err)
	}
	return explanation, nil
}
