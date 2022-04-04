package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func UseDB(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context)  {
		c.Set(constants.DB, db)
		c.Next()
	}
}
