package models

// UserStats represents the statistics of a user interacting with the bot.
type UserStats struct {
    UserID         int64   `json:"user_id"`          // Unique identifier for the user
    TotalTests     int     `json:"total_tests"`      // Total number of tests taken by the user
    CorrectAnswers  int     `json:"correct_answers"`   // Total number of correct answers given by the user
    IncorrectAnswers int     `json:"incorrect_answers"` // Total number of incorrect answers given by the user
    LastTestScore   float64 `json:"last_test_score"`   // Score of the last test taken by the user
    Language        string  `json:"language"`         // Preferred language of the user for localization
}

// NewUserStats creates a new UserStats instance for a user.
func NewUserStats(userID int64, language string) *UserStats {
    return &UserStats{
        UserID:         userID,
        TotalTests:     0,
        CorrectAnswers:  0,
        IncorrectAnswers: 0,
        LastTestScore:   0.0,
        Language:        language,
    }
}

// UpdateStats updates the user's statistics based on the test results.
func (us *UserStats) UpdateStats(correct bool, score float64) {
    us.TotalTests++
    if correct {
        us.CorrectAnswers++
    } else {
        us.IncorrectAnswers++
    }
    us.LastTestScore = score
}