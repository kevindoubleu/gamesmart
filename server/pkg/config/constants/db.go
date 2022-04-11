package constants

import "time"

const (
	DB = "gamesmart"

	DBO_TIMEOUT = 5 * time.Second

	// table names
	DB_TBL_USERS = "users"
	DB_TBL_QUESTIONS = "questions"

	// errors
	ErrDbInit    = "initializing database"
	ErrDbConnect = "connecting to database"

	ErrDbRead   = "reading from database"
	ErrDbInsert = "inserting into database"
	ErrDbUpdate = "updating database"
	ErrDbDelete = "deleting from database"
)
