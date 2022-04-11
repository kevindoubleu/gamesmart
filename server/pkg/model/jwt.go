package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type SessionJWT struct {
	Username	string	`json:"username"`
	
	jwt.StandardClaims
}

// SessionJWT implements the jwt.Claims interface
func (s SessionJWT) Valid() error {
	// for now is only based on standard claims
	return s.StandardClaims.Valid()
}

func NewSessionJWT(recipient User) SessionJWT {
	return SessionJWT{
		Username: recipient.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
}
