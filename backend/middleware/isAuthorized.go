package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mbient/todo-api/utils"
	"net/http"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the JWT token, validate it, and identify the user
		// If valid, call c.Next()
		// If invalid, call c.AbortWithStatusJSON(http.StatusUnauthorized)
		tokenString, err := utils.ExtractTokenFromHeader(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
