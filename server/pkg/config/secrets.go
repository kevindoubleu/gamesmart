package config

import (
	"crypto/rand"
	"log"
	"os"

	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
)

func GenerateSecrets() *model.Secrets {
	secrets := model.NewSecrets()

	if os.Getenv("GIN_MODE") != "release" {
		log.Println("NOT IN RELEASE MODE, ALL SECRETS ARE STATIC")
		return secrets
	}

	// jwt key
	_, err := rand.Read(secrets.JWTKey)
	if err != nil {
		log.Panic(constants.ErrSecretsInit)
	}

	return secrets
}
