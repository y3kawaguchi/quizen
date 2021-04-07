package repositories

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/joho/godotenv"
	"github.com/y3kawaguchi/quizen/internal/db"
	"github.com/y3kawaguchi/quizen/internal/domains"
)

var sqldb *sql.DB
var repo *QuizRepository

func setup() {
	teardown()

	loc, _ := time.LoadLocation("Asia/Tokyo")
	t1, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:26+09:00", loc)
	t2, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:27+09:00", loc)
	t3, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:28+09:00", loc)
	t4, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:29+09:00", loc)
	t5, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:30+09:00", loc)
	t6, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:31+09:00", loc)

	query := `INSERT INTO quizzes(
		title,
		question,
		answer,
		created_at,
		updated_at
	) VALUES
		('test_title_1', 'test_question_1', 'test_answer_1', $1, $2),
		('test_title_2', 'test_question_2', 'test_answer_2', $3, $4),
		('test_title_3', 'test_question_3', 'test_answer_3', $5, $6);
	`

	if _, err := sqldb.Exec(query, t1, t2, t3, t4, t5, t6); err != nil {
		log.Fatal(err)
	}
}

func teardown() {
	query := `
		TRUNCATE TABLE quizzes RESTART IDENTITY;
		select setval ('quizzes_id_seq', 1, false);
	`
	_, err := sqldb.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	envPath, err := filepath.Abs("../../.env.ci")
	if err != nil {
		log.Fatal(err)
	}
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file : ", err)
	}

	// PostgreSQL setup
	dbConfig, err := db.GetPostgreSQLConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	dbConnection, err := db.ConnectPostgreSQL(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()
	sqldb = dbConnection.GetDB()
	repo = NewQuizRepository(dbConnection)

	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func TestQuizRepository_FindByID(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	createdAt, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:26+09:00", loc)
	updatedAt, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:27+09:00", loc)

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    *domains.Quiz
		wantErr bool
	}{
		{
			name: "Returns the quiz specified by ID",
			args: args{
				id: 1,
			},
			want: &domains.Quiz{
				ID:        1,
				Title:     "test_title_1",
				Question:  "test_question_1",
				Answer:    "test_answer_1",
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuizRepository.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("QuizRepository.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuizRepository_Update(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	createdAt, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:26+09:00", loc)
	updatedAt, _ := time.ParseInLocation(time.RFC3339, "2021-01-04T19:52:27+09:00", loc)

	type args struct {
		quiz *domains.Quiz
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "update quiz and return zero",
			args: args{
				quiz: &domains.Quiz{
					ID:        1,
					Title:     "test_title_1_update",
					Question:  "test_question_1_update",
					Answer:    "test_answer_1_update",
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
				},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.Update(tt.args.quiz)
			if (err != nil) != tt.wantErr {
				t.Errorf("QuizRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QuizRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
