package config

import (
	"crypto/rand"
	"log"

	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
	"github.com/kevindoubleu/gamesmart/pkg/model"
)

func GenerateSecrets() *model.Secrets {
	secrets := model.Secrets{}

	// jwt key
	secrets.JWTKey = make([]byte, 32) // 256 bit / 8 = 32 byte
	_, err := rand.Read(secrets.JWTKey)
	if err != nil {
		log.Panic(constants.E_SECRETS_INIT)
	}

	return &secrets
}
