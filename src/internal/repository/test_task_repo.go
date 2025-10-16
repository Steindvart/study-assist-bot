package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"study-assist-bot-go/internal/models"
)

// TestTaskRepo defines the repository for test tasks.
type TestTaskRepo struct {
	db *sql.DB
}

// NewTestTaskRepo creates a new instance of TestTaskRepo.
func NewTestTaskRepo(db *sql.DB) *TestTaskRepo {
	return &TestTaskRepo{db: db}
}

// GetTestTask retrieves a test task by its ID.
func (r *TestTaskRepo) GetTestTask(id int) (*models.TestTask, error) {
	var task models.TestTask
	query := "SELECT id, question, options, correct_answer FROM test_tasks WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Question, &task.Options, &task.CorrectAnswer)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("test task not found")
		}
		return nil, fmt.Errorf("failed to get test task: %w", err)
	}
	return &task, nil
}

// CreateTestTask inserts a new test task into the database.
func (r *TestTaskRepo) CreateTestTask(task *models.TestTask) error {
	query := "INSERT INTO test_tasks (question, options, correct_answer) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, task.Question, task.Options, task.CorrectAnswer)
	if err != nil {
		return fmt.Errorf("failed to create test task: %w", err)
	}
	return nil
}

// UpdateTestTask updates an existing test task in the database.
func (r *TestTaskRepo) UpdateTestTask(task *models.TestTask) error {
	query := "UPDATE test_tasks SET question = ?, options = ?, correct_answer = ? WHERE id = ?"
	_, err := r.db.Exec(query, task.Question, task.Options, task.CorrectAnswer, task.ID)
	if err != nil {
		return fmt.Errorf("failed to update test task: %w", err)
	}
	return nil
}

// DeleteTestTask removes a test task from the database by its ID.
func (r *TestTaskRepo) DeleteTestTask(id int) error {
	query := "DELETE FROM test_tasks WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete test task: %w", err)
	}
	return nil
}

// GetAllTestTasks retrieves all test tasks from the database.
func (r *TestTaskRepo) GetAllTestTasks() ([]models.TestTask, error) {
	rows, err := r.db.Query("SELECT id, question, options, correct_answer FROM test_tasks")
	if err != nil {
		return nil, fmt.Errorf("failed to get test tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.TestTask
	for rows.Next() {
		var task models.TestTask
		if err := rows.Scan(&task.ID, &task.Question, &task.Options, &task.CorrectAnswer); err != nil {
			log.Println("Error scanning test task:", err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}