package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/controller"
)

func InitRouter(router *gin.Engine) {
	router.GET("/", controller.Home)
	router.GET("/ping", controller.Pong)

	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.POST("/albums", controller.PostAlbums)
}
