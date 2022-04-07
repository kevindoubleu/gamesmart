package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/controller"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
	"github.com/kevindoubleu/gamesmart/pkg/middleware"
)

func InitRouter(router *gin.Engine) {
	router.GET("/", controller.Home)
	router.GET("/ping", controller.Pong)

	// prepare resources needed by endpoints
	// db
	db := config.ConnectDb()
	// secrets
	secrets := config.GenerateSecrets()



	jwtSvc := helper.JWTService{
		JWTKey: secrets.JWTKey,
	}
	authSvc := controller.AuthService{
		UsersTable: db.Collection(constants.DB_TBL_USERS),
		JWTHelper: jwtSvc,
	}
	auth := router.Group("/auth")
	{
		auth.POST("/register", middleware.UnauthUser(jwtSvc), authSvc.Register)
		auth.POST("/login", middleware.UnauthUser(jwtSvc), authSvc.Login)
	}



	questionSvc := controller.QuestionService{
		QuestionsTable: db.Collection(constants.DB_TBL_QUESTIONS),
	}
	question := router.Group("/question")
	{
		question.GET("/", questionSvc.GetQuestions)
	}



	router.GET("/albums", controller.GetAlbums)
	router.GET("/albums/:id", controller.GetAlbumByID)
	router.POST("/albums", controller.PostAlbums)
}
