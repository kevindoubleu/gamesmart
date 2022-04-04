package model

import (
	"fmt"
	"strings"
)

type Question struct {
	Question	string	`bson:"question"`
	Answer		string	`bson:"answer"`
	Difficulty	int32	`bson:"difficulty"`
	Grade		int32	`bson:"grade"`
}

func (q Question) String() string {
	result := strings.Builder{}

	result.WriteString("question: " + q.Question + "\n")
	result.WriteString("answer: " + q.Answer + "\n")
	result.WriteString("difficulty: " + fmt.Sprintf("%d", q.Difficulty) + "\n")
	result.WriteString("grade: " + fmt.Sprintf("%d", q.Grade) + "\n")

	return result.String()
}
