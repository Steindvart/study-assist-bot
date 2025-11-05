package repositories

import (
	"context"
	"fmt"
	"study-assist-tgbot/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByTelegramID(ctx context.Context, telegramID int64) (*models.User, error)
	Create(ctx context.Context, telegramID int64, languageCode string) (*models.User, error)
	UpdateLanguage(ctx context.Context, telegramID int64, languageCode string) error
	UpsertLanguage(ctx context.Context, telegramID int64, languageCode string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByTelegramID(ctx context.Context, telegramID int64) (*models.User, error) {
	var user models.User

	result := r.db.WithContext(ctx).Where("telegram_id = ?", telegramID).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by telegram_id %d: %w", telegramID, result.Error)
	}

	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, telegramID int64, languageCode string) (*models.User, error) {
	user := &models.User{
		TelegramID:   telegramID,
		LanguageCode: languageCode,
	}

	result := r.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to create user: %w", result.Error)
	}

	return user, nil
}

func (r *userRepository) UpdateLanguage(ctx context.Context, telegramID int64, languageCode string) error {
	result := r.db.WithContext(ctx).Model(&models.User{}).
		Where("telegram_id = ?", telegramID).
		Update("language_code", languageCode)

	if result.Error != nil {
		return fmt.Errorf("failed to update user language: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with telegram_id %d not found", telegramID)
	}

	return nil
}

func (r *userRepository) UpsertLanguage(ctx context.Context, telegramID int64, languageCode string) (*models.User, error) {
	existingUser, err := r.GetByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	if existingUser != nil {
		if err := r.UpdateLanguage(ctx, telegramID, languageCode); err != nil {
			return nil, err
		}
		existingUser.LanguageCode = languageCode
		return existingUser, nil
	}

	return r.Create(ctx, telegramID, languageCode)
}

type userRepositoryKey struct{}

func WithUserRepository(ctx context.Context, repo UserRepository) context.Context {
	return context.WithValue(ctx, userRepositoryKey{}, repo)
}

func GetUserRepository(ctx context.Context) (UserRepository, bool) {
	repo, ok := ctx.Value(userRepositoryKey{}).(UserRepository)
	return repo, ok
}
