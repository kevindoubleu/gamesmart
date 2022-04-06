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

	// prepare resources neede by endpoints
	// db
	db := config.ConnectDb()
	router.Use(middleware.UseDB(db))
	// secrets
	secrets := config.GenerateSecrets()

	auth := router.Group("/auth", middleware.UseJWT(secrets))
	{
		auth.POST("/register", middleware.UnauthUser, controller.Register)
		auth.POST("/login", middleware.UnauthUser, controller.Login)
	}

	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.POST("/albums", controller.PostAlbums)
}
