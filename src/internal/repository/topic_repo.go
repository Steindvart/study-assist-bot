package repository

import (
	"database/sql"
	"errors"

	"study-assist-bot-go/internal/models"
)

// TopicRepository defines the interface for topic-related database operations.
type TopicRepository interface {
	GetTopicByID(id int) (*models.Topic, error)
	GetAllTopics() ([]models.Topic, error)
	CreateTopic(topic *models.Topic) error
	UpdateTopic(topic *models.Topic) error
	DeleteTopic(id int) error
}

// topicRepo implements the TopicRepository interface.
type topicRepo struct {
	db *sql.DB
}

// NewTopicRepo creates a new instance of TopicRepository.
func NewTopicRepo(db *sql.DB) TopicRepository {
	return &topicRepo{db: db}
}

// GetTopicByID retrieves a topic by its ID.
func (r *topicRepo) GetTopicByID(id int) (*models.Topic, error) {
	var topic models.Topic
	query := "SELECT id, name, section_id FROM topics WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&topic.ID, &topic.Name, &topic.SectionID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("topic not found")
		}
		return nil, err
	}
	return &topic, nil
}

// GetAllTopics retrieves all topics from the database.
func (r *topicRepo) GetAllTopics() ([]models.Topic, error) {
	rows, err := r.db.Query("SELECT id, name, section_id FROM topics")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topics []models.Topic
	for rows.Next() {
		var topic models.Topic
		if err := rows.Scan(&topic.ID, &topic.Name, &topic.SectionID); err != nil {
			return nil, err
		}
		topics = append(topics, topic)
	}
	return topics, rows.Err()
}

// CreateTopic inserts a new topic into the database.
func (r *topicRepo) CreateTopic(topic *models.Topic) error {
	query := "INSERT INTO topics (name, section_id) VALUES (?, ?)"
	_, err := r.db.Exec(query, topic.Name, topic.SectionID)
	return err
}

// UpdateTopic updates an existing topic in the database.
func (r *topicRepo) UpdateTopic(topic *models.Topic) error {
	query := "UPDATE topics SET name = ?, section_id = ? WHERE id = ?"
	_, err := r.db.Exec(query, topic.Name, topic.SectionID, topic.ID)
	return err
}

// DeleteTopic removes a topic from the database by its ID.
func (r *topicRepo) DeleteTopic(id int) error {
	query := "DELETE FROM topics WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}