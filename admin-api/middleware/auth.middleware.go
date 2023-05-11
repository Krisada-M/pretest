package middleware

import (
	"admin-api/helper"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authenticate is JWT authorization
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.GetHeader("Authorization")

		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}
		clientToken = strings.Split(clientToken, "Bearer ")[1]
		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("name", claims.Data.Username)
		c.Next()
	}
}

// AccountAuthenticate for User authorization
func AccountAuthenticate(ut string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetString("name")
		if user != os.Getenv("ADMIN_USERNAME") {
			c.Abort()
			return
		}
	}
}
