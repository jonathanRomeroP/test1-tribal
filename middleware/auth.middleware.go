package middleware

import (
	"net/http"
	"test1-tribal/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token incorrecto"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token invalido"})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)

		c.Next()
	}
}
