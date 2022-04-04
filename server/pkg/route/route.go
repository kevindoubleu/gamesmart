package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config"
	"github.com/kevindoubleu/gamesmart/pkg/controller"
	"github.com/kevindoubleu/gamesmart/pkg/middleware"
)

func InitRouter(router *gin.Engine) {
	router.GET("/", controller.Home)
	router.GET("/ping", controller.Pong)

	db := config.ConnectDb()
	router.Use(middleware.UseDB(db))

	auth := router.Group("/auth")
	{
		auth.POST("/register", controller.Register)
	}

	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.POST("/albums", controller.PostAlbums)
}
