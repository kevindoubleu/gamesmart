package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService struct {
	UsersTable	*mongo.Collection
	JWTService	helper.JWTService
}

func (svc AuthService) Register(c *gin.Context) {
	// parse req
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// check existing username
	filter := bson.M{"username":newUser.Username}
	row := svc.UsersTable.FindOne(c.Request.Context(), filter)
	if row.Err() == nil {
		c.JSON(http.StatusConflict, constants.DUPLICATE_USERNAME)
		return
	}

	// create new user
	result, err := svc.UsersTable.InsertOne(c.Request.Context(), newUser)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.ErrDbInsert, err)
		return
	}
	
	c.JSON(http.StatusCreated, result)
}

func (svc AuthService) Login(c *gin.Context) {
	// parse req
	var enteredUser model.User
	if err := c.BindJSON(&enteredUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// check username
	filter := bson.M{"username":enteredUser.Username}
	row := svc.UsersTable.FindOne(c.Request.Context(), filter)
	if row.Err() == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, constants.INVALID_CREDENTIALS)
		return
	}

	// check password
	var userInDb model.User
	err := row.Decode(&userInDb)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.ErrDbRead)
		return
	}

	if enteredUser.Password != userInDb.Password {
		c.JSON(http.StatusUnauthorized, constants.INVALID_CREDENTIALS)
		return
	}

	// get jwt generated with helper
	token, err := svc.JWTService.GenerateSession(enteredUser)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.ErrJWTVerify)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
