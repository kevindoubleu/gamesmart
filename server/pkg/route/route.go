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
	secrets	*model.Secrets

	dbConn	*mongo.Database
	jwtSvc	helper.JWTService
}

func defaultResources() resources {
	secrets := config.GenerateSecrets()

	return resources{
		secrets: secrets,
		dbConn: defaultDb(),
		jwtSvc: defaultJWTService(secrets),
	}
}

func defaultDb() *mongo.Database {
	return config.ConnectDb(constants.DB)
}

func defaultJWTService(secrets *model.Secrets) helper.JWTService {
	return helper.NewJWTService(secrets)
}

func InitRouter(router *gin.Engine) {
	// prepare resources needed by endpoints
	rsc := defaultResources()

	// fire up endpoints
	router.GET("/", controller.Home)
	authEndpoints(router, rsc)
	questionEndpoints(router, rsc)
}

func authEndpoints(router *gin.Engine, rsc resources) {
	authSvc := controller.AuthService{
		UsersTable: rsc.dbConn.Collection(constants.DB_TBL_USERS),
		JWTService: rsc.jwtSvc,
	}

	auth := router.Group("/auth", middleware.UnauthUser(authSvc.JWTService))
	{
		auth.POST("/register", authSvc.Register)
		auth.POST("/login", authSvc.Login)
	}
}

func questionEndpoints(router *gin.Engine, rsc resources) {
	questionSvc := controller.QuestionService{
		QuestionsTable: rsc.dbConn.Collection(constants.DB_TBL_QUESTIONS),
		JWTService: rsc.jwtSvc,
	}

	// TODO: add auth middleware here after testing
	question := router.Group("/question", middleware.AuthUser(questionSvc.JWTService))
	{
		question.GET("/", questionSvc.GetQuestions)
		question.GET("/:" + controller.QUESTION_ID, questionSvc.GetQuestionById)
		question.POST("/", questionSvc.AddQuestion)
		question.PATCH("/:" + controller.QUESTION_ID, questionSvc.UpdateQuestion)
		question.DELETE("/:" + controller.QUESTION_ID, questionSvc.DeleteQuestion)
	}
}
