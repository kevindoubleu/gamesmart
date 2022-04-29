package model

type Question struct {
	// metadata
	Id			string		`bson:"_id,omitempty"`
	Difficulty	int32		`bson:"difficulty"`
	Grade		int32		`bson:"grade"`

	// main content
	Question	string		`bson:"question"`
	Answer		string		`bson:"answer"`
	Explanation	string		`bson:"explanation"`

	// supporting content
	Choices		[]string	`bson:"choices"`
	Hints		[]string	`bson:"hints"`
}
