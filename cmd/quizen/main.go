package main

import (
	"time"

	"github.com/y3kawaguchi/quizen/internal/controller/handlers"
	"github.com/y3kawaguchi/quizen/internal/db"
	"github.com/y3kawaguchi/quizen/internal/repositories"
	"github.com/y3kawaguchi/quizen/internal/usecases"

	"github.com/gin-contrib/cors"
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

	// cors setting
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.GET("/quizzes", quizAPI.QuizzesGet())
	r.POST("/quizzes", quizAPI.QuizPost())
	r.PUT("/quizzes/:quiz_id", quizAPI.QuizPut())

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
