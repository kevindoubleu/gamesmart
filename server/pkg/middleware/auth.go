package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
)

// only allow authorized users
func AuthUser(jwtSvc helper.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwtSvc.GetValidToken(c)
		if err == nil && token != nil && token.Valid {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// only allow unauthorized users
func UnauthUser(jwtSvc helper.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwtSvc.GetValidToken(c)
		if err != nil && token == nil {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
