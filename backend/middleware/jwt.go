package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetToken(username string, key []byte) (string, error) {
	claims := CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
			NotBefore: time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(key)
}

func WithValidation(key []byte) func(c *gin.Context) {
	return func(c *gin.Context) {
		header := c.GetHeader("Authentication")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		tokenString := header[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		claims := token.Claims.(*CustomClaims)

		c.Set("user", claims.Username)
		c.Next()
	}
}
