package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	TelegramID   int64  `gorm:"primaryKey;column:telegram_id" json:"telegram_id"`
	LanguageCode string `gorm:"type:varchar(10);default:'en';not null" json:"language_code"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.LanguageCode == "" {
		u.LanguageCode = "en"
	}
	return nil
}
