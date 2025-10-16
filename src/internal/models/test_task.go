package models

// TestTask represents an individual test question with its associated data.
type TestTask struct {
	ID          int64  `json:"id"`
	Question    string `json:"question"`
	Options     []string `json:"options"`
	CorrectAnswer string `json:"correct_answer"`
	Explanation string `json:"explanation"`
	SectionID   int64  `json:"section_id"`
	TopicID     int64  `json:"topic_id"`
}

// NewTestTask creates a new TestTask instance.
func NewTestTask(id int64, question string, options []string, correctAnswer string, explanation string, sectionID int64, topicID int64) *TestTask {
	return &TestTask{
		ID:            id,
		Question:      question,
		Options:       options,
		CorrectAnswer: correctAnswer,
		Explanation:   explanation,
		SectionID:     sectionID,
		TopicID:       topicID,
	}
}

// IsCorrect checks if the provided answer is correct.
func (t *TestTask) IsCorrect(answer string) bool {
	return t.CorrectAnswer == answer
}