package middleware

import (
	"backend/app"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidateAndUpdateSession(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("auth_cookie")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		session, err := app.DB.GetSessionByToken(cookie)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Not authorized")
			return
		}

		if time.Now().Compare(session.Expiration) >= 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Token expired. Log in again")
			return
		}

		err = app.DB.UpdateSessionToken(*session)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Next()
	}
}
