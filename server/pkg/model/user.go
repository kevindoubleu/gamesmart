package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	// client provided
	Username	string					`bson:"username"`
	Password	string					`bson:"password"`
	Grade		int32					`bson:"grade"`

	// server initialized
	Money		int64					`bson:"money"`
	JoinDate	time.Time				`bson:"join_date"`
	Wrongs		[]primitive.ObjectID	`bson:"wrongs"`
	Corrects	[]primitive.ObjectID	`bson:"corrects"`
}

func (u *User) Init() {
	u.Money    = 0
	u.JoinDate = time.Now()
	u.Wrongs   = make([]primitive.ObjectID, 0)
	u.Corrects = make([]primitive.ObjectID, 0)
}
