package routes

import (
	"backend/app"
	"backend/database"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponseBody struct {
	Token string `json:"token"`
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

		newSession := database.Session{
			UserId:     user.Id,
			Token:      uuid.New().String(),
			Expiration: time.Now().Add(time.Minute * 30),
			Created:    time.Now(),
		}

		session, err := app.DB.CreateSession(newSession)

		c.IndentedJSON(http.StatusOK, LoginResponseBody{Token: session.Token})
		return
	}
}
