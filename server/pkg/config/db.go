package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kevindoubleu/gamesmart/pkg/config/constants"
)

func ConnectDb(name string) *mongo.Database {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(os.Getenv("MONGO_URI")),
		options.Client().SetAuth(options.Credential{
			Username: os.Getenv("MONGO_USERNAME"),
			Password: os.Getenv("MONGO_PASSWORD"),
		}),
	)
	if err != nil {
		log.Fatal(constants.E_DBI_INIT, err)
	}

	dbInit, cancel := context.WithTimeout(context.Background(), constants.DBO_TIMEOUT)
	defer cancel()
	err = client.Connect(dbInit)
	if err != nil {
		log.Fatal(constants.E_DBI_CONNECT, err)
	}

	return client.Database(name)
}
