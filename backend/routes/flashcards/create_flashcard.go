package routes

import (
	"backend/app"
	"backend/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateFlashcard(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newFlashcard database.Flashcard

		if err := c.BindJSON(&newFlashcard); err != nil {
			fmt.Println("Error with binding json")
			c.IndentedJSON(http.StatusBadRequest, "Issue with request")
			return
		}

		flashcard, err := app.FlashcardDB.CreateFlashcard(newFlashcard)

		if err != nil {
			fmt.Println("error with creating flashcard")
			c.IndentedJSON(http.StatusInternalServerError, "error with creating flashcard")
			return
		}

		c.IndentedJSON(http.StatusOK, flashcard)
	}
}
