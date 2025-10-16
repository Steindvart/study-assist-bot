package models

// Section represents a section in the study assist bot.
type Section struct {
	ID          int64  `json:"id"`          // Unique identifier for the section
	Title       string `json:"title"`       // Title of the section
	Description string `json:"description"` // Description of the section
	Topics      []int64 `json:"topics"`     // List of topic IDs associated with the section
}

// NewSection creates a new Section instance.
func NewSection(id int64, title, description string, topics []int64) *Section {
	return &Section{
		ID:          id,
		Title:       title,
		Description: description,
		Topics:      topics,
	}
}

// AddTopic adds a topic ID to the section.
func (s *Section) AddTopic(topicID int64) {
	s.Topics = append(s.Topics, topicID)
}

// RemoveTopic removes a topic ID from the section.
func (s *Section) RemoveTopic(topicID int64) {
	for i, id := range s.Topics {
		if id == topicID {
			s.Topics = append(s.Topics[:i], s.Topics[i+1:]...)
			break
		}
	}
}