package main

import (
	"backend/app"
	"backend/database"
	"backend/middleware"
	flashcardRoutes "backend/routes/flashcards"
	userRoutes "backend/routes/users"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "quizlet",
		ParseTime: true,
	}

	db, err := database.ConnectToDB(cfg)

	if err != nil {
		fmt.Println("Connect to database not working")
	} else {
		fmt.Println("Connected to database")
	}

	app := &app.App{DB: database.Database{DB: db}}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Frontend on localhost:5173
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},   // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,
	}))

	authRequired := router.Group("")

	authRequired.Use(middleware.ValidateAndUpdateSession(app))
	{
		authRequired.GET("/flashcards", flashcardRoutes.HandleGetFlashcards(app))
		authRequired.POST("/flashcards", flashcardRoutes.HandleCreateFlashcard(app))
		authRequired.PATCH("/flashcards", flashcardRoutes.HandleUpdateFlashcard(app))
		authRequired.DELETE("/flashcards/:flashcardId", flashcardRoutes.HandleDeleteFlashcard(app))
	}

	router.POST("/register", userRoutes.HandleRegister(app))
	router.POST("/login", userRoutes.HandleLogin(app))
	router.POST("/logout", userRoutes.HandleLogout(app))

	router.Run("localhost:8080")
}
