package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionService struct {
	QuestionsTable	*mongo.Collection
}

func (svc QuestionService) GetQuestions(c *gin.Context) {
	result := []model.Question{}

	cursor, err := svc.QuestionsTable.Find(c.Request.Context(), bson.M{})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_DBO_READ, err)
		return
	}

	err = cursor.All(c.Request.Context(), &result)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_DBO_READ, err)
		return
	}

	c.JSON(http.StatusOK, result)
}


