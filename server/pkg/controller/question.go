package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	QUESTION_ID = "id"
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

func (svc QuestionService) GetQuestionById(c *gin.Context) {
	// get id param from url
	hexId := c.Param(QUESTION_ID)
	objId, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// query
	row := svc.QuestionsTable.FindOne(c.Request.Context(), bson.M{"_id":objId})
	if row.Err() == mongo.ErrNoDocuments {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var result model.Question
	row.Decode(&result)
	c.JSON(http.StatusOK, result)
}

func (svc QuestionService) AddQuestion(c *gin.Context) {
	// parse req
	var newQuestion model.Question
	if err := c.BindJSON(&newQuestion);err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// insert
	result, err := svc.QuestionsTable.InsertOne(c.Request.Context(), newQuestion)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_DBO_INSERT, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (svc QuestionService) UpdateQuestion(c *gin.Context) {
	// get id param from url
	hexId := c.Param(QUESTION_ID)
	objId, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// parse req
	var newQuestion model.Question
	if err := c.BindJSON(&newQuestion);err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// update in db
	result, err := svc.QuestionsTable.UpdateByID(c.Request.Context(), objId, bson.M{
		"$set": newQuestion,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_DBO_UPDATE, err)
		return
	}

	if result.MatchedCount == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// query the updated doc
	svc.GetQuestionById(c)
}

func (svc QuestionService) DeleteQuestion(c *gin.Context) {
	// get id param from url
	hexId := c.Param(QUESTION_ID)
	objId, err := primitive.ObjectIDFromHex(hexId)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// delete from db
	deleted, err := svc.QuestionsTable.DeleteOne(c.Request.Context(), bson.M{"_id":objId})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_DBO_DELETE, err)
		return
	}

	if deleted.DeletedCount == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}
