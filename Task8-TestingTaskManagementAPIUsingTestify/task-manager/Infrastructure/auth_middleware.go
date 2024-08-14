package Infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	service JwtService
}

func NewMiddleware(jwtservice JwtService) *Middleware {
	return &Middleware{
		service: jwtservice,
	}
}

type MiddlewareInterface interface {
	Auth_middleware()
}

func (s *Middleware) Auth_middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headers := strings.Split(header, " ")
		if len(headers) != 2 || headers[0] != "bearer" {
			c.JSON(http.StatusBadRequest, "Authorization is requered")
			c.Abort()
			return
		} else {

			claims, err := s.service.ValidateJwt(headers[1])
			if err == nil {
				c.Set("claims", claims)
				c.Next()
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "could not validate claims", "error": err.Error()})
				c.Abort()
				return
			}
		}
	}
}
