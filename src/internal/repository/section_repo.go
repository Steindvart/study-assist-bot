package repository

import (
	"context"
	"database/sql"
	"errors"

	"study-assist-bot-go/internal/models"
)

// SectionRepository defines the interface for section data access.
type SectionRepository interface {
	GetSectionByID(ctx context.Context, id int) (*models.Section, error)
	GetAllSections(ctx context.Context) ([]models.Section, error)
	CreateSection(ctx context.Context, section *models.Section) error
	UpdateSection(ctx context.Context, section *models.Section) error
	DeleteSection(ctx context.Context, id int) error
}

// sectionRepo implements the SectionRepository interface.
type sectionRepo struct {
	db *sql.DB
}

// NewSectionRepository creates a new instance of SectionRepository.
func NewSectionRepository(db *sql.DB) SectionRepository {
	return &sectionRepo{db: db}
}

// GetSectionByID retrieves a section by its ID.
func (r *sectionRepo) GetSectionByID(ctx context.Context, id int) (*models.Section, error) {
	section := &models.Section{}
	query := "SELECT id, name, description FROM sections WHERE id = ?"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&section.ID, &section.Name, &section.Description)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Section not found
		}
		return nil, err
	}
	return section, nil
}

// GetAllSections retrieves all sections from the database.
func (r *sectionRepo) GetAllSections(ctx context.Context) ([]models.Section, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, description FROM sections")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sections []models.Section
	for rows.Next() {
		var section models.Section
		if err := rows.Scan(&section.ID, &section.Name, &section.Description); err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}
	return sections, rows.Err()
}

// CreateSection inserts a new section into the database.
func (r *sectionRepo) CreateSection(ctx context.Context, section *models.Section) error {
	query := "INSERT INTO sections (name, description) VALUES (?, ?)"
	_, err := r.db.ExecContext(ctx, query, section.Name, section.Description)
	return err
}

// UpdateSection updates an existing section in the database.
func (r *sectionRepo) UpdateSection(ctx context.Context, section *models.Section) error {
	query := "UPDATE sections SET name = ?, description = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, section.Name, section.Description, section.ID)
	return err
}

// DeleteSection removes a section from the database by its ID.
func (r *sectionRepo) DeleteSection(ctx context.Context, id int) error {
	query := "DELETE FROM sections WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}