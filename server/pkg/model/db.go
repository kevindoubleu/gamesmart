package model

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Conn		*mongo.Database
	Context		context.Context
	CancelFunc	context.CancelFunc
}
