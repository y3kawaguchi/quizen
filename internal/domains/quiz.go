package domains

import "time"

// Quiz ...
type Quiz struct {
	ID        int64
	Title     string
	Question  string
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
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
		ID:        q.ID,
		Title:     item.Title,
		Question:  item.Question,
		Answer:    item.Answer,
		CreatedAt: q.CreatedAt,
		UpdatedAt: q.UpdatedAt,
	}
}
