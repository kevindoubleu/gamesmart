package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
)

const (
	CurrUser = "current user in context"
)

type JWTService struct {
	JWTKey	[]byte
}

func (svc JWTService) GenerateSession(recipient model.User) (string, error) {
	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": recipient.Username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	// sign jwt
	signedToken, err := token.SignedString(svc.JWTKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (svc JWTService) ValidSession(c *gin.Context) bool {
	// take client provided jwt from http header
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 { return false }
	tokenString := authHeader[len("Bearer "):]

	// verify jwt
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		key, _ := c.Get(constants.SECRETS_JWT)
		fmt.Println("key is", key)
		return key, nil
	})

	if err != nil {
		log.Println(constants.E_JWT_VERIFY, err)
		return false
	}

	return parsedToken.Valid
}
