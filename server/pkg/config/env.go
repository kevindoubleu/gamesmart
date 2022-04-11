package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(constants.ErrMissingEnv)
	}
}
