package middleware

import (
	"mp-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token", "info": "you should authorize"})
			c.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		isExpired, err := utils.IsTokenExpired(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		if isExpired {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "token was expired"})
			c.Abort()
			return
		}
		c.Next()
	}
}
