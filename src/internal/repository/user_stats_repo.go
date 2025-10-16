package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"study-assist-bot-go/internal/models"
)

// UserStatsRepo defines the repository for user statistics.
type UserStatsRepo struct {
	db *sql.DB
}

// NewUserStatsRepo creates a new instance of UserStatsRepo.
func NewUserStatsRepo(db *sql.DB) *UserStatsRepo {
	return &UserStatsRepo{db: db}
}

// CreateUserStats inserts a new user statistics record into the database.
func (repo *UserStatsRepo) CreateUserStats(ctx context.Context, stats *models.UserStats) error {
	query := `INSERT INTO user_stats (user_id, correct_answers, total_answers) VALUES (?, ?, ?)`
	_, err := repo.db.ExecContext(ctx, query, stats.UserID, stats.CorrectAnswers, stats.TotalAnswers)
	if err != nil {
		log.Printf("Error inserting user stats: %v", err)
		return fmt.Errorf("could not insert user stats: %w", err)
	}
	return nil
}

// GetUserStats retrieves user statistics by user ID.
func (repo *UserStatsRepo) GetUserStats(ctx context.Context, userID int64) (*models.UserStats, error) {
	query := `SELECT user_id, correct_answers, total_answers FROM user_stats WHERE user_id = ?`
	row := repo.db.QueryRowContext(ctx, query, userID)

	var stats models.UserStats
	if err := row.Scan(&stats.UserID, &stats.CorrectAnswers, &stats.TotalAnswers); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No stats found for this user
		}
		log.Printf("Error retrieving user stats: %v", err)
		return nil, fmt.Errorf("could not retrieve user stats: %w", err)
	}
	return &stats, nil
}

// UpdateUserStats updates the user statistics for a given user ID.
func (repo *UserStatsRepo) UpdateUserStats(ctx context.Context, stats *models.UserStats) error {
	query := `UPDATE user_stats SET correct_answers = ?, total_answers = ? WHERE user_id = ?`
	_, err := repo.db.ExecContext(ctx, query, stats.CorrectAnswers, stats.TotalAnswers, stats.UserID)
	if err != nil {
		log.Printf("Error updating user stats: %v", err)
		return fmt.Errorf("could not update user stats: %w", err)
	}
	return nil
}