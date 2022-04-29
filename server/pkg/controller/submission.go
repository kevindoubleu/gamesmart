package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubmissionService struct {
	UsersTable		*mongo.Collection
	QuestionsTable	*mongo.Collection
}

func (svc SubmissionService) Submit(c *gin.Context) {

}
