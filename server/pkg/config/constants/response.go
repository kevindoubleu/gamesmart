package constants

import "github.com/gin-gonic/gin"

var (
	DUPLICATE_USERNAME = errMsg("Username is taken")

	INVALID_CREDENTIALS = errMsg("Invalid username or password")
)

func errMsg(msg string) map[string]interface{} {
	return gin.H{
		"message": msg,
	}
}
