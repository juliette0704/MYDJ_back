package middleware

import (
	authtoken "mydj_server/src/authToken"
	"net/http"

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
