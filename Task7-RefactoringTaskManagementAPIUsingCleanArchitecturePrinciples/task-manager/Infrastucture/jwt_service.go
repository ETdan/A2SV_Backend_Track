package infrastucture

import (
	domain "TaskManager/task-manager/Domain"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service struct{}

var secrate_key = []byte(os.Getenv("SECRATE_KEY"))

func NewService() *Service {
	return &Service{}
}

type JwtService interface {
	GenerateJwt(user domain.User, expTime time.Duration) (string, error)
	ValidateJwt(tokenString string) (jwt.MapClaims, error)
}

// func (j *JwtService) NewJwtService() *JwtService {
// 	return JwtService.GenerateJwt()
// }

func (*Service) GenerateJwt(user domain.User, expTime time.Duration) (string, error) {
	fmt.Println("user.ID")
	fmt.Println(user.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.UserName,
		"role":     user.Role,
		"exp":      time.Now().Add(expTime).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secrate_key))
	if err == nil {
		return tokenString, nil
	} else {

		return "", err
	}
}

func (*Service) ValidateJwt(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secrate_key, nil
	})
	if t, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return t, nil
	} else {
		return nil, err
	}
}
