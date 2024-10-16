package middleware

import (
	"backend/app"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ValidateAndUpdateSession(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "You are not authorized")
			return
		}

		sessionToken := strings.TrimPrefix(token, "Bearer ")

		session, err := app.DB.GetSessionByToken(sessionToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		if time.Now().Compare(session.Expiration) >= 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Token expired. Log in again")
			return
		}

		c.Next()
	}
}
