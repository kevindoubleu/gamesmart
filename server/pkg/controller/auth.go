package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/controller/helper"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(c *gin.Context) {
	// get db conn fron context and make timeout db operation ctx
	db := c.MustGet(constants.DB).(*mongo.Database)
	dbOp, cancel := context.WithTimeout(context.Background(), constants.DBO_TIMEOUT)
	defer cancel()

	// parse req
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(newUser)

		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// check existing username
	usersTbl := db.Collection(constants.DB_TBL_USERS)
	row := usersTbl.FindOne(dbOp, bson.M{"username":newUser.Username})
	if row.Err() == nil {
		c.JSON(http.StatusConflict, constants.DUPLICATE_USERNAME)
		return
	}

	// create new user
	result, err := usersTbl.InsertOne(dbOp, newUser)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_DBO_INSERT, err)
		return
	}
	
	c.JSON(http.StatusCreated, result)
}

func Login(c *gin.Context) {
	// get db conn fron context and make timeout db operation ctx
	db := c.MustGet(constants.DB).(*mongo.Database)
	dbOp, cancel := context.WithTimeout(context.Background(), constants.DBO_TIMEOUT)
	defer cancel()

	// parse req
	var enteredUser model.User
	if err := c.BindJSON(&enteredUser); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// check username
	usersTbl := db.Collection(constants.DB_TBL_USERS)
	row := usersTbl.FindOne(dbOp, bson.M{"username":enteredUser.Username})
	if row.Err() == mongo.ErrNoDocuments {
		c.JSON(http.StatusUnauthorized, constants.INVALID_CREDENTIALS)
		return
	}

	// check password
	var userInDb model.User
	err := row.Decode(&userInDb)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_DBO_READ)
		return
	}

	if enteredUser.Password != userInDb.Password {
		c.JSON(http.StatusUnauthorized, constants.INVALID_CREDENTIALS)
		return
	}

	// get jwt generated with helper
	c.Set(helper.CurrUser, enteredUser)
	token, err := helper.GenerateJWTSession(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(constants.E_JWT_VERIFY)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
