package domains

import (
	"reflect"
	"testing"
	"time"
)

func TestQuiz_Change(t *testing.T) {
	createdAt, _ := time.Parse(time.RFC3339, "2000-01-01T12:34:56+00:00")
	updatedAt, _ := time.Parse(time.RFC3339, "2000-01-01T12:34:56+00:00")

	type fields struct {
		ID        int64
		Title     string
		Question  string
		Answer    string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		item Quiz
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Quiz
	}{
		{
			name: "change title and question and answer",
			fields: fields{
				ID:        1,
				Title:     "test_title_1",
				Question:  "test_question_1",
				Answer:    "test_answer_1",
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			args: args{
				item: Quiz{
					ID:        1,
					Title:     "test_title_1_update",
					Question:  "test_question_1_update",
					Answer:    "test_answer_1_update",
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
				},
			},
			want: &Quiz{
				ID:        1,
				Title:     "test_title_1_update",
				Content:   "test_content_1_update",
				Answer:    "test_answer_1_update",
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Quiz{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Question:  tt.fields.Question,
				Answer:    tt.fields.Answer,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if got := q.Change(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quiz.Change() = %v, want %v", got, tt.want)
			}
		})
	}
}
