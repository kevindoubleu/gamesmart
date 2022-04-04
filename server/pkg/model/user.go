package model

import "time"

type User struct {
	// client provided
	Username	string		`bson:"username"`
	Password	string		`bson:"password"`
	Grade		int32		`bson:"grade"`

	// server initialized
	Money		int64		`bson:"money"`
	JoinDate	time.Time	`bson:"join_date"`
}

func (u *User) Init() {
	u.Money = 0
	u.JoinDate = time.Now()
}
