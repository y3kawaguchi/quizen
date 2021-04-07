package main

import (
	"time"

	"github.com/y3kawaguchi/quizen/internal/controller/handlers"
	"github.com/y3kawaguchi/quizen/internal/db"
	"github.com/y3kawaguchi/quizen/internal/repositories"
	"github.com/y3kawaguchi/quizen/internal/usecases"

	"github.com/gin-gonic/gin"
)

const (
	location = "Asia/Tokyo"
	offset   = 9 * 60 * 60
)

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, offset)
	}
	time.Local = loc
}

func main() {
	// PostgreSQL setup
	dbConfig, err := db.GetPostgreSQLConfigFromEnv()
	if err != nil {
		// TODO: plan to implement logging
	}
	dbConnection, err := db.ConnectPostgreSQL(dbConfig)
	if err != nil {
		// TODO: plan to implement logging
	}
	defer dbConnection.Close()

	// inject dbConnection to repository
	quizRepository := repositories.NewQuizRepository(dbConnection)

	// inject quizRepository to usecase
	quizUsecase := usecases.NewQuizUsecase(quizRepository)

	// inject quizUsecase to handler
	quizAPI := handlers.NewQuizAPI(quizUsecase)

	r := gin.Default()
	r.GET("/quizzes", quizAPI.QuizzesGet())
	r.POST("/quizzes", quizAPI.QuizPost())
	r.PUT("/quizzes/:quiz_id", quizAPI.QuizPut())

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
