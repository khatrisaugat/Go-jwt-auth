package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/JWTPractise/helpers"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
