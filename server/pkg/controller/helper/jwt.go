package helper

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
)

const (
	CurrUser = "current user in context"
)

func GenerateJWTSession(c *gin.Context) (string, error) {
	recipient := c.MustGet(CurrUser).(model.User)

	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": recipient.Username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	// sign jwt
	key, _ := c.Get(constants.SECRETS_JWT)
	signedToken, err := token.SignedString(key)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_JWT_VERIFY, err)
		return "", err
	}

	return signedToken, nil
}

func ValidJWTSession(c *gin.Context) bool {
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
