package domains

import (
	"time"
)

type Choice struct {
	QuizID    int64     `db:"quiz_id" json:"quiz_id"`
	ChoiceID  int16     `db:"choice_id" json:"choice_id"`
	Content   string    `db:"content" json:"content"`
	IsCorrect bool      `db:"is_correct" json:"is_correct"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
