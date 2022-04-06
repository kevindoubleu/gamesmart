package heroku

import (
	"github.com/gin-gonic/gin"

	"github.com/kevindoubleu/gamesmart/pkg/config"
	"github.com/kevindoubleu/gamesmart/pkg/route"
)

func Start() {
	// load env vars
	config.LoadEnv()

	// server-wide initializations done, start the server
	router := gin.Default()
	route.InitRouter(router)
	router.Run()
}
