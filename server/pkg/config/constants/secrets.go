package constants

import "errors"

const (
	SECRETS = "secrets"
	
	// JWT
	SECRETS_JWT = "secrets_jwt"
)

// errors
var (
	// ENV
	ErrMissingEnv  = errors.New("couldn't find .env file")
	ErrSecretsInit = errors.New("couldn't generate secrets")

	// JWT
	ErrJWTInvalid = errors.New("invalid JWT")
	ErrJWTVerify  = errors.New("couldn't verify JWT")
	ErrJWTNil     = errors.New("JWT is nil")
)
