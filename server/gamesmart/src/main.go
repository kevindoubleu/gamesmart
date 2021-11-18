package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kevindoubleu/gamesmart/route"
)

func main() {
	router := gin.Default()

	route.InitRouter(router)

	router.Run()
}