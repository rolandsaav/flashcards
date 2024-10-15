package routes

import (
	"backend/app"
	"backend/database"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleRegister(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body RegisterRequestBody
		if err := c.BindJSON(&body); err != nil {
			fmt.Println("Error binding json")
			c.IndentedJSON(http.StatusBadRequest, "Issue with request")
			return
		}

		unique, err := app.DB.IsUniqueUsername(body.Username)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Database error")
			return
		}

		if !unique {
			c.IndentedJSON(http.StatusConflict, "Username already taken")
			return
		}

		salt := make([]byte, 32)
		_, err = rand.Read(salt)
		if err != nil {
			c.IndentedJSON(http.StatusConflict, "Error creating salt")
			return
		}

		stringSalt := base64.StdEncoding.EncodeToString(salt)

		concatBytes := append(salt, body.Password...)

		hashed, err := bcrypt.GenerateFromPassword(concatBytes, bcrypt.DefaultCost)

		if err != nil {
			c.IndentedJSON(http.StatusConflict, "Error hasing password")
			return
		}

		user := database.User{
			Username: body.Username,
			Hashed:   string(hashed),
			Salt:     stringSalt,
		}

		_, err = app.DB.CreateUser(user)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "Error creating user")
			return
		}

		c.IndentedJSON(http.StatusOK, "Created user")
	}
}
