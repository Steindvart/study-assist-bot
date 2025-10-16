package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/bot/updates"
	"study-assist-bot-go/internal/service"
)

// Handler struct to hold dependencies
type Handler struct {
	TestService      *service.TestService
	StatsService     *service.StatsService
	ExplanationService *service.ExplanationService
	LocalizationService *service.LocalizationService
}

// NewHandler creates a new Handler instance
func NewHandler(testService *service.TestService, statsService *service.StatsService, explanationService *service.ExplanationService, localizationService *service.LocalizationService) *Handler {
	return &Handler{
		TestService:      testService,
		StatsService:     statsService,
		ExplanationService: explanationService,
		LocalizationService: localizationService,
	}
}

// HandleMessage processes incoming messages
func (h *Handler) HandleMessage(ctx context.Context, update updates.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "start":
		h.handleStart(ctx, update.Message)
	case "test":
		h.handleTest(ctx, update.Message)
	case "stats":
		h.handleStats(ctx, update.Message)
	default:
		h.handleUnknownCommand(ctx, update.Message)
	}
}

// handleStart handles the /start command
func (h *Handler) handleStart(ctx context.Context, message *models.Message) {
	response := h.LocalizationService.GetMessage("welcome")
	if _, err := bot.SendMessage(message.Chat.ID, response); err != nil {
		log.Printf("Error sending welcome message: %v", err)
	}
}

// handleTest handles the /test command
func (h *Handler) handleTest(ctx context.Context, message *models.Message) {
	// Logic to start a test session
	testSession, err := h.TestService.StartTestSession(message.Chat.ID)
	if err != nil {
		log.Printf("Error starting test session: %v", err)
		return
	}
	
	// Send the first question
	question := testSession.GetCurrentQuestion()
	if _, err := bot.SendMessage(message.Chat.ID, question.Text); err != nil {
		log.Printf("Error sending question: %v", err)
	}
}

// handleStats handles the /stats command
func (h *Handler) handleStats(ctx context.Context, message *models.Message) {
	stats, err := h.StatsService.GetUserStats(message.Chat.ID)
	if err != nil {
		log.Printf("Error retrieving user stats: %v", err)
		return
	}

	response := fmt.Sprintf("Your stats: %v", stats)
	if _, err := bot.SendMessage(message.Chat.ID, response); err != nil {
		log.Printf("Error sending stats message: %v", err)
	}
}

// handleUnknownCommand handles unrecognized commands
func (h *Handler) handleUnknownCommand(ctx context.Context, message *models.Message) {
	response := h.LocalizationService.GetMessage("unknown_command")
	if _, err := bot.SendMessage(message.Chat.ID, response); err != nil {
		log.Printf("Error sending unknown command message: %v", err)
	}
}