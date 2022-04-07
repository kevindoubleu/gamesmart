package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/controller"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
	"github.com/kevindoubleu/gamesmart/pkg/middleware"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type resources struct {
	db		*mongo.Database
	secrets	*model.Secrets
}

func defaultResources() resources {
	return resources{
		db: config.ConnectDb("gamesmart"),
		secrets: config.GenerateSecrets(),
	}
}

func InitRouter(router *gin.Engine) {
	// prepare resources needed by endpoints
	res := defaultResources()

	// fire up endpoints
	router.GET("/", controller.Home)
	authEndpoints(router, res)
	questionEndpoints(router, res)
}

func authEndpoints(router *gin.Engine, res resources) {
	jwtSvc := helper.JWTService{
		JWTKey: res.secrets.JWTKey,
	}
	authSvc := controller.AuthService{
		UsersTable: res.db.Collection(constants.DB_TBL_USERS),
		JWTHelper: jwtSvc,
	}

	auth := router.Group("/auth")
	{
		auth.POST("/register", middleware.UnauthUser(jwtSvc), authSvc.Register)
		auth.POST("/login", middleware.UnauthUser(jwtSvc), authSvc.Login)
	}
}

func questionEndpoints(router *gin.Engine, res resources) {
	questionSvc := controller.QuestionService{
		QuestionsTable: res.db.Collection(constants.DB_TBL_QUESTIONS),
	}

	question := router.Group("/question")
	{
		question.GET("/", questionSvc.GetQuestions)
		question.GET("/:" + controller.QUESTION_ID, questionSvc.GetQuestionById)
		question.POST("/", questionSvc.AddQuestion)
		question.PATCH("/:" + controller.QUESTION_ID, questionSvc.UpdateQuestion)
		question.DELETE("/:" + controller.QUESTION_ID, questionSvc.DeleteQuestion)
	}
}
