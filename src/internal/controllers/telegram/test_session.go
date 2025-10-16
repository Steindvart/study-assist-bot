package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/bot/service"
)

// TestSession represents a user's test session.
type TestSession struct {
	UserID      int64
	CurrentTask int
	Answers     []string
}

// TestManager manages user test sessions.
type TestManager struct {
	sessions map[int64]*TestSession
	repo     repository.TestTaskRepository
}

// NewTestManager creates a new TestManager.
func NewTestManager(repo repository.TestTaskRepository) *TestManager {
	return &TestManager{
		sessions: make(map[int64]*TestSession),
		repo:     repo,
	}
}

// StartTest initializes a new test session for the user.
func (tm *TestManager) StartTest(userID int64) error {
	tasks, err := tm.repo.GetTestTasks()
	if err != nil {
		return fmt.Errorf("failed to get test tasks: %w", err)
	}

	tm.sessions[userID] = &TestSession{
		UserID:      userID,
		CurrentTask: 0,
		Answers:     make([]string, len(tasks)),
	}

	return nil
}

// AnswerQuestion records the user's answer to the current question.
func (tm *TestManager) AnswerQuestion(userID int64, answer string) error {
	session, exists := tm.sessions[userID]
	if !exists {
		return fmt.Errorf("no active session for user %d", userID)
	}

	session.Answers[session.CurrentTask] = answer
	session.CurrentTask++

	return nil
}

// FinishTest finalizes the test session and provides feedback.
func (tm *TestManager) FinishTest(userID int64) (string, error) {
	session, exists := tm.sessions[userID]
	if !exists {
		return "", fmt.Errorf("no active session for user %d", userID)
	}

	// Logic to evaluate answers and provide feedback
	score := evaluateAnswers(session.Answers)
	delete(tm.sessions, userID)

	return fmt.Sprintf("Test completed! Your score: %d", score), nil
}

// evaluateAnswers evaluates the user's answers and returns a score.
func evaluateAnswers(answers []string) int {
	// Placeholder for actual evaluation logic
	return 0
}

// HandleTestSession handles incoming messages related to test sessions.
func HandleTestSession(ctx context.Context, b *bot.Bot, update models.Update, tm *TestManager) {
	userID := update.Message.From.ID

	switch update.Message.Text {
	case "/start_test":
		if err := tm.StartTest(userID); err != nil {
			log.Println("Error starting test:", err)
			b.SendMessage(update.Message.Chat.ID, "Failed to start test.")
			return
		}
		b.SendMessage(update.Message.Chat.ID, "Test started! Please answer the questions.")
	default:
		if err := tm.AnswerQuestion(userID, update.Message.Text); err != nil {
			log.Println("Error answering question:", err)
			b.SendMessage(update.Message.Chat.ID, "Error recording your answer.")
			return
		}
		b.SendMessage(update.Message.Chat.ID, "Your answer has been recorded.")
	}
}
