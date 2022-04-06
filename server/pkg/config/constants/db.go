package constants

import "time"

const (
	DB = "gamesmart"

	DBO_TIMEOUT = 5 * time.Second

	// table names
	DB_TBL_USERS = "users"
	DB_TBL_QUESTIONS = "questions"

	// errors
	E_DBI_INIT    = "initializing database"
	E_DBI_CONNECT = "connecting to database"

	E_DBO_READ   = "reading from database"
	E_DBO_INSERT = "inserting into database"
)
