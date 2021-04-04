package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConnection() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb://tinyuser:tinypass@tiny-url-db:27017/tinydb")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to MongoDB!")
	return client
}
