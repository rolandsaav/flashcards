package routes

import (
	"backend/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetFlashcards(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		flashcards, err := app.FlashcardDB.GetFlashcards()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
		}

		if len(flashcards) == 0 {
			c.IndentedJSON(http.StatusOK, "No flashcards found")
		}

		c.IndentedJSON(http.StatusOK, flashcards)
	}
}
