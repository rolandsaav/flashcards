package routes

import (
	"backend/app"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body LoginRequestBody

		if err := c.BindJSON(&body); err != nil {
			fmt.Print("Error binding json")
			c.IndentedJSON(http.StatusBadRequest, "Issue with request")
			return
		}

		user, err := app.DB.GetUserFromUsername(body.Username)
		if err != nil {
			fmt.Print("Error getting user")
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		base64salt, err := base64.StdEncoding.DecodeString(user.Salt)
		if err != nil {
			fmt.Print("Error decoding salt")
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		concatBytes := append(base64salt, body.Password...)

		err = bcrypt.CompareHashAndPassword([]byte(user.Hashed), concatBytes)
		if err != nil {
			fmt.Print("Incorrect password")
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, "You are logged in")
		return
	}
}
