package routes

import (
	"backend/database"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetFlashcards(c *gin.Context, db *sql.DB) {
	flashcards, err := database.GetFlashcards(db)

	if err != nil {
		fmt.Println("Get flashcards handler error")
	}

	c.IndentedJSON(http.StatusOK, flashcards)
}
