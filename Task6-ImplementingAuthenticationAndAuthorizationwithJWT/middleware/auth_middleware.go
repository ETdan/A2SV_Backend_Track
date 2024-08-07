package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secrate_key = []byte(os.Getenv("SECRATE_KEY"))

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headers := strings.Split(header, " ")
		if len(headers) != 2 || headers[0] != "bearer" {
			c.JSON(http.StatusBadRequest, "Authorization is requered")
			c.Abort()
			return
		} else {

			token, err := jwt.Parse(headers[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return secrate_key, nil
			})
			if err != nil || !token.Valid {
				c.JSON(401, gin.H{"error": "Invalid JWT"})
				c.Abort()
				return
			} else {
				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid claims"})
					c.Abort()
					return
				}

				c.Set("claims", claims)
				c.Next()
			}

		}
	}
}
