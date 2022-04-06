package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
)

func UseJWT(secrets *model.Secrets) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(constants.SECRETS_JWT, secrets.JWTKey)
		c.Next()
	}
}