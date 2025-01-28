package middlewares

import (
	"Spam_Span/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Remove the "Bearer " prefix if it's present
		token = strings.TrimPrefix(token, "Bearer ")

		// Validate token and extract the claims
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user_id in the context for further use
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
