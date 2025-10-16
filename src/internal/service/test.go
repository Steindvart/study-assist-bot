package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-telegram/bot/v2"
	"study-assist-bot-go/internal/models"
	"study-assist-bot-go/internal/repository"
)

// TestService handles the business logic for conducting tests.
type TestService struct {
	repo repository.TestTaskRepository
	mu   sync.Mutex
}

// NewTestService creates a new instance of TestService.
func NewTestService(repo repository.TestTaskRepository) *TestService {
	return &TestService{repo: repo}
}

// ConductTest manages the test session for a user.
func (s *TestService) ConductTest(userID int64, sectionID int) ([]models.TestTask, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks, err := s.repo.GetTestTasksBySection(sectionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get test tasks: %w", err)
	}

	if len(tasks) == 0 {
		return nil, errors.New("no test tasks available for this section")
	}

	// Logic to manage user test session can be added here.
	return tasks, nil
}

// ScoreTest evaluates the user's answers and provides feedback.
func (s *TestService) ScoreTest(userID int64, answers map[int]int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	score := 0
	for questionID, userAnswer := range answers {
		correctAnswer, err := s.repo.GetCorrectAnswer(questionID)
		if err != nil {
			return 0, fmt.Errorf("failed to get correct answer for question %d: %w", questionID, err)
		}
		if userAnswer == correctAnswer {
			score++
		}
	}

	return score, nil
}

// ProvideFeedback generates feedback based on the user's performance.
func (s *TestService) ProvideFeedback(score int, total int) string {
	percentage := (float64(score) / float64(total)) * 100
	if percentage >= 80 {
		return "Excellent work! You have a strong understanding of the material."
	} else if percentage >= 50 {
		return "Good job! You have a decent understanding, but there's room for improvement."
	}
	return "Keep trying! Review the material and don't hesitate to ask for help."
}