package usecases

import (
	"fmt"

	"github.com/y3kawaguchi/quizen/internal/domains"
	"github.com/y3kawaguchi/quizen/internal/repositories"
)

// QuizUsecase ...
type QuizUsecase struct {
	repository repositories.QuizRepository
}

// NewQuizUsecase ...
func NewQuizUsecase(repository *repositories.QuizRepository) *QuizUsecase {
	usecase := &QuizUsecase{
		repository: *repository,
	}
	return usecase
}

// Create ...
func (q *QuizUsecase) Create(quiz *domains.Quiz) (int64, error) {
	quizId, err := q.repository.Save(quiz)
	if err != nil {
		return -1, err
	}

	for _, choice := range quiz.Choices {
		choice.QuizID = quizId
		q.repository.SaveChoice(&choice)
		if err != nil {
			return -1, err
		}
	}

	return quizId, nil
}

// Get ...
func (q *QuizUsecase) Get(quizId int64) (*domains.Quiz, error) {
	quiz, err := q.repository.FindByID(quizId)
	if err != nil {
		return quiz, err
	}

	choices, err := q.repository.GetChoicesByQuizID(quiz.ID)
	if err != nil {
		return quiz, err
	}
	quiz.Choices = choices

	return quiz, nil
}

// Search ...
func (q *QuizUsecase) Search() (*domains.Quizzes, error) {
	quizzes, err := q.repository.FindAll()
	fmt.Printf("QuizUsecase.Search(): %#v\n", quizzes)
	return quizzes, err
}

// Update ...
func (q *QuizUsecase) Update(quiz *domains.Quiz) (int64, error) {
	target, err := q.repository.FindByID(quiz.ID)
	if err != nil {
		return -1, err
	}
	changed := target.Change(*quiz)
	return q.repository.Update(changed)
}
