package routes

import (
	"backend/app"
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUpdateFlashcard(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateFlashcard database.Flashcard

		if err := c.BindJSON(&updateFlashcard); err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Issue with request")
			return
		}

		if updateFlashcard.Term == "" && updateFlashcard.Definition == "" {
			c.IndentedJSON(http.StatusBadRequest, "Term or Definition must be updated.")
			return
		}

		resultFlashcard, err := app.FlashcardDB.UpdateFlashcard(updateFlashcard)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, resultFlashcard)
	}
}
