package form

import (
	"github.com/y3kawaguchi/quizen/internal/domains"
)

// Quiz ...
type Quiz struct {
	Title            string   `json:"title" binding:"required"`
	Question         string   `json:"question" binding:"required"`
	Explanation      string   `json:"explanation" binding:"required"`
	ChoiceContents   []string `json:"choice_contents" binding:"required"`
	ChoiceIsCorrects []bool   `json:"choice_is_corrects" binding:"required"`
}

// BuildDomain ...
func (f *Quiz) BuildDomain() *domains.Quiz {
	contents := f.ChoiceContents
	isCorrects := f.ChoiceIsCorrects

	choices := make([]domains.Choice, 0)
	for i := 0; i < len(contents); i++ {
		choices = append(choices, domains.Choice{
			ChoiceID:  int16(i + 1),
			Content:   contents[i],
			IsCorrect: isCorrects[i],
		})
	}

	return &domains.Quiz{
		Title:       f.Title,
		Question:    f.Question,
		Choices:     choices,
		Explanation: f.Explanation,
	}
}
