package helper

import (
	"errors"
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

var (
	ErrInvalidJWT = errors.New("JWT is not valid")
	ErrMissingJWTHeader = errors.New("JWT session not found in HTTP header")
)

type JWTService interface {
	// get the provided JWT key from the secrets provided
	GetJWTKey() []byte

	// generate a new session for the user
	// returns the JWT in signed string format
	GenerateSession(model.User) (string, error)

	// tries to get the token in context
	// if token is invalid then a nil token will be returned
	GetValidToken(*gin.Context) (*jwt.Token, error)
}

func NewJWTService(secrets *model.Secrets) JWTService {
	return myJWTService{
		key: secrets.JWTKey,
	}
}

type myJWTService struct {
	key	[]byte
}

func (svc myJWTService) GetJWTKey() []byte {
	return svc.key
}

func (svc myJWTService) GenerateSession(recipient model.User) (string, error) {
	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": recipient.Username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	// sign jwt
	signedToken, err := token.SignedString(svc.key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (svc myJWTService) GetValidToken(c *gin.Context) (*jwt.Token, error) {
	// take client provided jwt from http header
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 { return nil, ErrMissingJWTHeader }
	tokenString := authHeader[len("Bearer "):]

	// verify jwt
	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return svc.key, nil
	})

	if err != nil {
		log.Println(constants.ErrJWTVerify, err)
		return nil, err
	}

	if !parsedToken.Valid {
		log.Println(ErrInvalidJWT)
		return nil, err
	}

	return parsedToken, nil
}
