package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host         string        `env:"DB_HOST" envDefault:"localhost"`
	Port         int           `env:"DB_PORT" envDefault:"5432"`
	User         string        `env:"DB_USER" envDefault:"postgres"`
	Password     string        `env:"DB_PASSWORD" envDefault:"postgres"`
	Database     string        `env:"DB_NAME" envDefault:"study_assist"`
	SSLMode      string        `env:"DB_SSL_MODE" envDefault:"disable"`
	MaxOpenConns int           `env:"DB_MAX_OPEN_CONNS" envDefault:"10"`
	MaxIdleConns int           `env:"DB_MAX_IDLE_CONNS" envDefault:"5"`
	MaxLifetime  time.Duration `env:"DB_MAX_LIFETIME" envDefault:"1h"`
}

type Database struct {
	db *gorm.DB
}

func New(cfg Config) (*Database, error) {
	connectionStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.Open(connectionStr), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.MaxLifetime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	if d.db != nil {
		sqlDB, err := d.db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

func (d *Database) DB() *gorm.DB {
	return d.db
}

func (d *Database) Health(ctx context.Context) error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return sqlDB.PingContext(ctx)
}
