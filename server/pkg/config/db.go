package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb() (*mongo.Database, error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(os.Getenv("MONGO_URI")),
		options.Client().SetAuth(options.Credential{
			Username: os.Getenv("MONGO_USERNAME"),
			Password: os.Getenv("MONGO_PASSWORD"),
		}),
	)
	if err != nil {
		log.Fatal(E_DBI_INIT, err)
		return nil, err
	}

	dbInit, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err = client.Connect(dbInit)
	if err != nil {
		log.Fatal(E_DBI_CONNECT, err)
		return nil, err
	}

	return client.Database("gamesmart"), nil
}
