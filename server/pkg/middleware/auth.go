package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
)

// only allow authorized users
func AuthUser(svc helper.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if svc.ValidSession(c) {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

// only allow unauthorized users
func UnauthUser(svc helper.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !svc.ValidSession(c) {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
