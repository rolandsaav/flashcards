package routes

import (
	"backend/app"
	"backend/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetFlashcards(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		flashcards, err := database.GetFlashcards(app.DB)

		if err != nil {
			fmt.Println("Get flashcards handler error")
		}

		c.IndentedJSON(http.StatusOK, flashcards)
	}
}
