package constants

const (
	SECRETS = "secrets"
	
	// ENV
	ErrMissingEnv = "couldn't find .env file"
	ErrSecretsInit  = "couldn't generate secrets"

	// JWT
	SECRETS_JWT = "secrets_jwt"
	
	ErrJWTInvalid = "invalid JWT"
	ErrJWTVerify = "couldn't verify JWT"
)