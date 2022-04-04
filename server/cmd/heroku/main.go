package heroku

import (
	"github.com/gin-gonic/gin"

	"github.com/kevindoubleu/gamesmart/pkg/config"
	"github.com/kevindoubleu/gamesmart/pkg/route"
)

func Start() {
	config.LoadEnv()

	// db := config.ConnectDb()
	// defer db.CancelFunc()

	// cursor, err := db.Conn.Collection("questions").Find(db.Context, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(db.Context)

	// result := make([]model.Question, 0)
	// for cursor.Next(db.Context) {
	// 	var doc model.Question
	// 	err := cursor.Decode(&doc)
	// 	if err != nil {
	// 		log.Fatal(constants.E_DBO_READ, err)
	// 	}

	// 	result = append(result, doc)
	// }

	// fmt.Println(result)

	// all initializations done, start the server
	router := gin.Default()
	route.InitRouter(router)
	router.Run()
}
