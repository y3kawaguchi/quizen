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
			&item.Explanation,
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
		&item.Explanation,
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
func (q *QuizRepository) GetChoicesByQuizID(quizId int64) ([]domains.Choice, error) {
	db := q.connection.GetDB()

	query := `SELECT * FROM choices WHERE quiz_id = $1`
	rows, err := db.Queryx(query, quizId)
	if err != nil {
		// TODO: nilを返すのが適切か考える
		return []domains.Choice{}, errors.WithStack(err)
	}
	defer rows.Close()

	choices := make([]domains.Choice, 0)
	for rows.Next() {
		var choice domains.Choice
		err := rows.StructScan(&choice)
		if err != nil {
			return []domains.Choice{}, errors.WithStack(err)
		}
		choices = append(choices, choice)
	}

	return choices, nil
}

// Save ...
func (q *QuizRepository) Save(quiz *domains.Quiz) (int64, error) {
	now := time.Now()

	row, err := q.connection.GetDB().Query(`INSERT INTO quizzes
		(
			title,
			question,
			explanation,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5
		) RETURNING id`,
		quiz.Title,
		quiz.Question,
		quiz.Explanation,
		now,
		now,
	)
	if err != nil {
		return -1, err
	}
	defer row.Close()

	row.Next()
	if err := row.Scan(&quiz.ID); err != nil {
		return -1, err
	}

	return quiz.ID, nil
}

// SaveChoice ...
func (q *QuizRepository) SaveChoice(choice *domains.Choice) (int64, error) {
	now := time.Now()

	_, err := q.connection.GetDB().Exec(`INSERT INTO choices
		(
			quiz_id,
			choice_id,
			content,
			is_correct,
			created_at,
			updated_at
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		)`,
		choice.QuizID,
		choice.ChoiceID,
		choice.Content,
		choice.IsCorrect,
		now,
		now,
	)
	if err != nil {
		return -1, err
	}

	return 0, nil
}

// Update ...
func (q *QuizRepository) Update(quiz *domains.Quiz) (int64, error) {
	now := time.Now()

	_, err := q.connection.GetDB().Exec(
		`UPDATE quizzes
			SET
				title = $1,
				question = $2,
				explanation = $3,
				updated_at = $4
			WHERE id = $5`,
		quiz.Title,
		quiz.Question,
		quiz.Explanation,
		now,
		quiz.ID,
	)
	if err != nil {
		return -1, err
	}

	return 0, err
}
