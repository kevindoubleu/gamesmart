package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
)

// only allow authorized users
func AuthUser(c *gin.Context) {
	if helper.ValidJWTSession(c) {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// only allow unauthorized users
func UnauthUser(c *gin.Context) {
	if !helper.ValidJWTSession(c) {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
