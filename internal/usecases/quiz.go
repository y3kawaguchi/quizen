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
	return q.repository.Save(quiz)
}

// Get ...
func (q *QuizUsecase) Get() (*domains.Quizzes, error) {
	quizzes, err := q.repository.FindAll()
	fmt.Printf("QuizUsecase.Get(): %#v\n", quizzes)
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
