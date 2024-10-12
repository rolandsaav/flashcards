package main

import (
	"backend/app"
	"backend/database"
	routes "backend/routes/flashcards"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"

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

	fmt.Println(cfg.User, cfg.Passwd)

	db, err := database.ConnectToDB(cfg)

	if err != nil {
		fmt.Println("Connect to database not working")
	} else {
		fmt.Println("Connected to database")
	}

	app := &app.App{DB: db}

	router := gin.Default()

	router.GET("/flashcards", routes.HandleGetFlashcards(app))

	router.Run("localhost:8080")
}
