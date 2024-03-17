package middleware

import (
	authtoken "mydj_server/src/authToken"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.String(http.StatusUnauthorized, "Token manquant")
			c.Abort()
			return
		}

		// Vérifie si le token commence par "Bearer "
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.String(http.StatusUnauthorized, "Format de token invalide")
			c.Abort()
			return
		}

		// Supprime le préfixe "Bearer " du token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := authtoken.ValidateToken(tokenString)
		if err != nil {
			c.String(http.StatusUnauthorized, "Token invalide")
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}

func GuardMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
