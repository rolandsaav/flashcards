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

	testcard := database.Flashcard{
		Term:       "Test please",
		Definition: "Test definition",
		OwnerID:    20,
	}

	cardid, err := database.CreateFlashcard(db, testcard)

	if err != nil {
		fmt.Println("create flashcard error")
	}

	fmt.Println(cardid)

	flashcards, err := database.GetFlashcardsByOwner(db, 19)

	if err != nil {
		fmt.Println("couldn't get flashcards")
	}

	for _, card := range flashcards {
		fmt.Println(card.String())
	}

	router := gin.Default()

	router.GET("/flashcards", routes.HandleGetFlashcards(app))

	router.Run("localhost:8080")
}
