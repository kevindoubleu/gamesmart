package constants

import (
	"errors"
	"time"
)

const (
	DB = "gamesmart"

	DBO_TIMEOUT = 5 * time.Second

	// table names
	DB_TBL_USERS = "users"
	DB_TBL_QUESTIONS = "questions"
)

// errors
var (
	ErrDbInit    = errors.New("initializing database")
	ErrDbConnect = errors.New("connecting to database")

	ErrDbRead   = errors.New("reading from database")
	ErrDbInsert = errors.New("inserting into database")
	ErrDbUpdate = errors.New("updating database")
	ErrDbDelete = errors.New("deleting from database")

	ErrDbAggregate = errors.New("aggregating from database")
)
