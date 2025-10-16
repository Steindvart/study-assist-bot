package models

// Topic represents a topic in the study assist bot.
type Topic struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SectionID   int64  `json:"section_id"` // Foreign key to the Section model
}

// NewTopic creates a new Topic instance.
func NewTopic(id int64, title, description string, sectionID int64) *Topic {
	return &Topic{
		ID:          id,
		Title:       title,
		Description: description,
		SectionID:   sectionID,
	}
}

// GetID returns the ID of the topic.
func (t *Topic) GetID() int64 {
	return t.ID
}

// GetTitle returns the title of the topic.
func (t *Topic) GetTitle() string {
	return t.Title
}

// GetDescription returns the description of the topic.
func (t *Topic) GetDescription() string {
	return t.Description
}

// GetSectionID returns the associated section ID of the topic.
func (t *Topic) GetSectionID() int64 {
	return t.SectionID
}