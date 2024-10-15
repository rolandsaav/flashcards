package routes

import (
	"backend/app"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleDeleteFlashcard(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		flashcardId := c.Param("flashcardId")

		id, err := strconv.ParseInt(flashcardId, 10, 64)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		deleted, err := app.DB.DeleteFlashcard(id)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, deleted)
	}
}
