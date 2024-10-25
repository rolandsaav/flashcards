package routes

import (
	"backend/app"
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetFlashcardsResponse struct {
	Data  []database.Flashcard `json:"data"`
	Error string               `json:"error"`
}

func HandleGetFlashcards(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		flashcards, err := app.DB.GetFlashcards()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, &GetFlashcardsResponse{
				Error: err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusOK, &GetFlashcardsResponse{
			Data: flashcards,
		})
	}
}
