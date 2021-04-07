package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/y3kawaguchi/quizen/internal/controller/form"
	"github.com/y3kawaguchi/quizen/internal/domains"

	"github.com/gin-gonic/gin"
)

// Quiz ...
type Quiz interface {
	Create(quiz *domains.Quiz) (int64, error)
	Get() (*domains.Quizzes, error)
	Update(quiz *domains.Quiz) (int64, error)
}

// QuizAPI ...
type QuizAPI struct {
	quiz Quiz
}

// NewQuizAPI ...
func NewQuizAPI(quiz Quiz) *QuizAPI {
	return &QuizAPI{
		quiz: quiz,
	}
}

// QuizzesGet ...
func (q QuizAPI) QuizzesGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := q.quiz.Get()
		if err != nil {
			c.Error(err).SetMeta(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, result)
	}
}

type quizPostRequest struct {
	Title    string `json:"title" binding:"required"`
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

// QuizPost ...
func (q QuizAPI) QuizPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var f form.Quiz
		if err := c.ShouldBindJSON(&f); err != nil {
			c.Error(err).SetMeta(http.StatusUnprocessableEntity)
			return
		}
		quiz := f.BuildDomain()
		_, err := q.quiz.Create(quiz)
		if err != nil {
			fmt.Printf("error: %#v\n", err)
			c.Error(err).SetMeta(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}

// QuizPut ...
func (q QuizAPI) QuizPut() gin.HandlerFunc {
	return func(c *gin.Context) {
		var f form.Quiz
		if err := c.ShouldBindJSON(&f); err != nil {
			c.Error(err).SetMeta(http.StatusUnprocessableEntity)
			return
		}

		id, err := strconv.Atoi(c.Param("quiz_id"))
		if err != nil {
			c.Error(err).SetMeta(http.StatusNotFound)
			return
		}

		quiz := f.BuildDomain()
		quiz.ID = int64(id)

		if _, err := q.quiz.Update(quiz); err != nil {
			fmt.Printf("error: %#v\n", err)
			c.Error(err).SetMeta(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusNoContent)
		}
	}
}
