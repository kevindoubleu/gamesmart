package model

type Question struct {
	Id			string		`bson:"_id,omitempty"`
	Question	string		`bson:"question"`
	Answer		string		`bson:"answer"`
	Difficulty	int32		`bson:"difficulty"`
	Grade		int32		`bson:"grade"`
	Hints		[]string	`bson:"hints"`
}
