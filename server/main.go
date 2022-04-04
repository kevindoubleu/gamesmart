package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/kevindoubleu/gamesmart/pkg/config"
	"github.com/kevindoubleu/gamesmart/pkg/model"
	"github.com/kevindoubleu/gamesmart/pkg/route"
)

func main() {
	config.LoadEnv()

	db, err := config.ConnectDb()
	if err != nil {
		panic("db init error:" + err.Error())
	}

	fmt.Println("db initializec")

	dbOperation, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()
	cursor, err := db.Collection("questions").Find(dbOperation, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(dbOperation)

	result := make([]model.Question, 0)
	for cursor.Next(dbOperation) {
		var doc model.Question
		err := cursor.Decode(&doc)
		if err != nil {
			log.Fatal(config.E_DBO_READ, err)
		}

		result = append(result, doc)
	}

	fmt.Println(result)

	fmt.Println("starting router")


	router := gin.Default()
	route.InitRouter(router)
	router.Run()
}
