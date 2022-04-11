package route

import (
	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/controller"
	"github.com/kevindoubleu/gamesmart/pkg/middleware"
)

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
		UserService: rsc.userSvc,
	}

	question := router.Group("/question", middleware.AuthUser(questionSvc.JWTService))
	{
		// single CRUD
		question.GET("/:" + controller.QUESTION_ID, questionSvc.GetQuestionById)
		question.POST("/", questionSvc.AddQuestion)
		question.PATCH("/:" + controller.QUESTION_ID, questionSvc.UpdateQuestion)
		question.DELETE("/:" + controller.QUESTION_ID, questionSvc.DeleteQuestion)

		// reads
		question.GET("/", questionSvc.GetAllQuestions)
		question.GET("/bygrade", questionSvc.GetRandomQuestionByGrade)
	}
}
