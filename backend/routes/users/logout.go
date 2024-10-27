package routes

import (
	"backend/app"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleLogout(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("auth_cookie")

		if err != nil {
			fmt.Print("Error getting auth_cookie")
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return
		}

		if cookie == "" {
			fmt.Printf("Not logged in. Can not log out")
			c.IndentedJSON(http.StatusForbidden, "Not logged in")
			return
		}

		err = app.DB.InvalidateSessionToken(cookie)

		if err != nil {
			fmt.Printf("Error invalidating session token %s", err.Error())
		}

		c.SetCookie("auth_cookie", "", 0, "/", "localhost", false, true)

		c.IndentedJSON(http.StatusOK, "Logged out")
	}
}
