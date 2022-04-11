package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
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
	JWTService		helper.JWTService
	UserService		helper.UserService
}

// basic single CRUD

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
		log.Println(constants.ErrDbInsert, err)
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
		log.Println(constants.ErrDbUpdate, err)
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
		log.Println(constants.ErrDbDelete, err)
		return
	}

	if deleted.DeletedCount == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK)
}

// READS

func (svc QuestionService) GetAllQuestions(c *gin.Context) {
	result := []model.Question{}

	cursor, err := svc.QuestionsTable.Find(c.Request.Context(), bson.M{})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.ErrDbRead, err)
		return
	}

	err = cursor.All(c.Request.Context(), &result)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.ErrDbRead, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (svc QuestionService) GetRandomQuestionByGrade(c *gin.Context) {
	// get username
	token, _ := svc.JWTService.GetValidToken(c)
	username, err := svc.UserService.GetUsernameFromSession(token)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// get user
	user := svc.UserService.GetUserByUsername(c.Request.Context(), username)
	grade := user.Grade
	log.Println(grade)

	// find random question in grade
	cursor, err := svc.QuestionsTable.Aggregate(c.Request.Context(), []bson.M{
		{"$match": bson.M{"grade": grade}},
		{"$sample": bson.M{"size": 1}},
	})

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.ErrDbAggregate)
		return
	}

	var question model.Question
	cursor.Next(c.Request.Context())
	cursor.Decode(&question)

	// json the picked question
	c.JSON(http.StatusOK, question)
}
