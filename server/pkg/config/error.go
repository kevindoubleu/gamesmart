package config

const (
	E      = "Error "
	SUFFIX = ": "

	// ENV
	E_ENV_FILE = E + "couldn't find .env file"

	// DATABASE
	E_DBI_INIT    = E + "initializing database" + SUFFIX
	E_DBI_CONNECT = E + "connecting to database" + SUFFIX

	E_DBO_READ = E + "reading from database" + SUFFIX
)