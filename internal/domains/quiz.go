package domains

import "time"

// Quiz ...
type Quiz struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Question    string    `json:"question"`
	Choices     []Choice  `json:"choices"`
	Explanation string    `json:"explanation"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Quizzes ...
type Quizzes struct {
	Items []Quiz
}

// QuizzesNew ...
func QuizzesNew() *Quizzes {
	return &Quizzes{}
}

// Add ...
func (q *Quizzes) Add(item Quiz) {
	q.Items = append(q.Items, item)
}

// GetAll ...
func (q *Quizzes) GetAll() []Quiz {
	return q.Items
}

// Change ...
func (q *Quiz) Change(item Quiz) *Quiz {
	return &Quiz{
		ID:          q.ID,
		Title:       item.Title,
		Question:    item.Question,
		Choices:     item.Choices,
		Explanation: item.Explanation,
		CreatedAt:   q.CreatedAt,
		UpdatedAt:   q.UpdatedAt,
	}
}
