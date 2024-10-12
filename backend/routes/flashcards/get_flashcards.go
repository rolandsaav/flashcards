package routes

import (
	"backend/app"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetFlashcards(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		flashcards, err := app.FlashcardDB.GetFlashcards()

		if err != nil {
			fmt.Println("Get flashcards handler error")
		}

		if len(flashcards) == 0 {
			c.IndentedJSON(http.StatusOK, "No flashcards found")
			return
		}

		c.IndentedJSON(http.StatusOK, flashcards)
	}
}
