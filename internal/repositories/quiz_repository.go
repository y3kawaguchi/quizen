package repositories

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/y3kawaguchi/quizen/internal/db"
	"github.com/y3kawaguchi/quizen/internal/domains"
	"github.com/y3kawaguchi/quizen/pkg/location"
)

// QuizRepository ...
type QuizRepository struct {
	connection db.Connection
}

// NewQuizRepository ...
func NewQuizRepository(connection db.Connection) *QuizRepository {
	return &QuizRepository{
		connection: connection,
	}
}

// FindAll ...
func (a *QuizRepository) FindAll() (*domains.Quizzes, error) {
	db := a.connection.GetDB()

	query := `SELECT * FROM quizzes`
	rows, err := db.Query(query)
	if err != nil {
		// TODO: nilを返すのが適切か考える
		return nil, err
	}
	defer rows.Close()

	quizzes := domains.QuizzesNew()
	for rows.Next() {
		fmt.Printf("rows: %#v\n", rows)

		item := domains.Quiz{}
		err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Question,
			&item.Answer,
			&item.CreatedAt,
			&item.UpdatedAt,
		)
		if err != nil {
			// TODO: nilを返すのが適切か考える
			return nil, err
		}
		quizzes.Add(item)
	}

	fmt.Printf("QuizRepository.FindAll(): %#v\n", quizzes)

	return quizzes, nil
}

// FindByID ...
func (q *QuizRepository) FindByID(id int64) (*domains.Quiz, error) {
	db := q.connection.GetDB()

	// fmt.Printf("db: %#v\n", db)
	// fmt.Printf("id: %#v\n", id)

	query := `SELECT * FROM quizzes where id = $1`

	item := domains.Quiz{}
	err := db.QueryRow(query, id).Scan(
		&item.ID,
		&item.Title,
		&item.Question,
		&item.Answer,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		// TODO: nilを返すのが適切か考える
		return nil, errors.WithStack(err)
	}

	// fmt.Printf("QuizRepository.FindByID(): %#v\n", item)

	item.CreatedAt = item.CreatedAt.In(location.JP())
	item.UpdatedAt = item.UpdatedAt.In(location.JP())

	return &item, nil
}

// GetChoicesByQuizID ...
func (q *QuizRepository) GetChoicesByQuizID(quizId int64) ([]string, error) {
	db := q.connection.GetDB()

	query := `SELECT content FROM choices WHERE quiz_id = $1`
	rows, err := db.Query(query, quizId)
	if err != nil {
		// TODO: nilを返すのが適切か考える
		return []string{}, errors.WithStack(err)
	}
	defer rows.Close()

	choices := make([]string, 0)
	for rows.Next() {
		var choice string
		err := rows.Scan(&choice)
		if err != nil {
			return []string{}, errors.WithStack(err)
		}
		choices = append(choices, choice)
	}

	return choices, nil
}

// Save ...
func (q *QuizRepository) Save(quiz *domains.Quiz) (int64, error) {
	now := time.Now()

	_, err := q.connection.GetDB().Exec(`INSERT INTO quizzes
		(
			title,
			question,
			answer,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		)`,
		quiz.Title,
		quiz.Question,
		quiz.Answer,
		now,
		now,
	)
	if err != nil {
		return -1, err
	}

	return 0, err
}

// Update ...
func (q *QuizRepository) Update(quiz *domains.Quiz) (int64, error) {
	now := time.Now()

	_, err := q.connection.GetDB().Exec(
		`UPDATE quizzes
			SET
				title = $1,
				question = $2,
				answer = $3,
				updated_at = $4
			WHERE id = $5`,
		quiz.Title,
		quiz.Question,
		quiz.Answer,
		now,
		quiz.ID,
	)
	if err != nil {
		return -1, err
	}

	return 0, err
}
