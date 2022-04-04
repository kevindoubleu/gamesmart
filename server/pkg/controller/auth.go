package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(c *gin.Context) {
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
	fmt.Println("good", newUser)

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
