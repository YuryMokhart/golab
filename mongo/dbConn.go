package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConnect connects to the database.
func DBConnect() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("could not connect mongoDB to a new client: %s\n", err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("could not initialise the client: %s\n", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("could not ping: %s\n", err)
	}
	collection := client.Database("tournament").Collection("user")
	return collection
}
