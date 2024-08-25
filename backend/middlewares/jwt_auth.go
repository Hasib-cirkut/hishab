package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	return func(c *gin.Context) {
		authCookie, err := c.Cookie("Authorization")

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if authCookie == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Auth header missing",
			})

			return
		}

		token, err := jwt.Parse(authCookie, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtSecret, nil
		})

		if err != nil || token.Valid == false {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Token",
			})

			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["sub"])
			c.Set("exp", claims["exp"])
		}

		c.Next()
	}
}
