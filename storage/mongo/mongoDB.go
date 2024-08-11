package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Hour)

	defer cancel()

	clientOPtion := options.Client().ApplyURI("mongodb://localhost:27017")


	client, err := mongo.Connect(ctx, clientOPtion)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Println("salodfkemdsfmiewksdvofmiosm kmdvc")
	return client.Database("car"), nil

}