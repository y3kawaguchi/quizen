package form

import (
	"github.com/y3kawaguchi/quizen/internal/domains"
)

// Quiz ...
type Quiz struct {
	Title    string `json:"title" binding:"required"`
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

// BuildDomain ...
func (f *Quiz) BuildDomain() *domains.Quiz {
	return &domains.Quiz{
		Title:    f.Title,
		Question: f.Question,
		Answer:   f.Answer,
	}
}
