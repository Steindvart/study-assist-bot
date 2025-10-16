package service

import (
	"context"
	"fmt"
	"sync"

	"study-assist-bot-go/internal/models"
	"study-assist-bot-go/internal/repository"
)

// StatsService provides methods for tracking and reporting user statistics.
type StatsService struct {
	repo repository.UserStatsRepository
	mu   sync.Mutex
}

// NewStatsService creates a new instance of StatsService.
func NewStatsService(repo repository.UserStatsRepository) *StatsService {
	return &StatsService{repo: repo}
}

// TrackUserResponse tracks a user's response to a test question.
func (s *StatsService) TrackUserResponse(ctx context.Context, userID int, sectionID int, topicID int, correct bool) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	stats, err := s.repo.GetUserStats(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user stats: %w", err)
	}

	if correct {
		stats.CorrectAnswers++
	} else {
		stats.IncorrectAnswers++
	}

	if err := s.repo.UpdateUserStats(ctx, stats); err != nil {
		return fmt.Errorf("failed to update user stats: %w", err)
	}

	return nil
}

// GetUserStats retrieves the statistics for a specific user.
func (s *StatsService) GetUserStats(ctx context.Context, userID int) (*models.UserStats, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	stats, err := s.repo.GetUserStats(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user stats: %w", err)
	}

	return stats, nil
}

// ReportUserStats generates a report of user statistics.
func (s *StatsService) ReportUserStats(ctx context.Context, userID int) (string, error) {
	stats, err := s.GetUserStats(ctx, userID)
	if err != nil {
		return "", err
	}

	report := fmt.Sprintf("User ID: %d\nCorrect Answers: %d\nIncorrect Answers: %d\n", userID, stats.CorrectAnswers, stats.IncorrectAnswers)
	return report, nil
}
