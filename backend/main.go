package main

import (
	"backend/app"
	"backend/database"
	routes "backend/routes/flashcards"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "quizlet",
	}

	db, err := database.ConnectToDB(cfg)

	if err != nil {
		fmt.Println("Connect to database not working")
	} else {
		fmt.Println("Connected to database")
	}

	app := &app.App{FlashcardDB: database.FlashcardDB{DB: db}}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Frontend on localhost:5173
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,
	}))

	router.GET("/flashcards", routes.HandleGetFlashcards(app))
	router.POST("/flashcards", routes.HandleCreateFlashcard(app))
	router.PATCH("/flashcards", routes.HandleUpdateFlashcard(app))
	router.DELETE("/flashcards/:flashcardId", routes.HandleDeleteFlashcard(app))

	router.Run("localhost:8080")
}
