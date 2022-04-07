package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is the backend API for gamesmart, the frontend is at https://gamesmart.netlify.app/",
	})
}
