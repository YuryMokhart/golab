package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBConnect connects to the mongo database.
func DBConnect() *ModelMongo {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil //, fmt.Errorf("could not create a new client to connect to the database: %s", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil //, fmt.Errorf("could not create a new client to connect to the database: %s", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil //, fmt.Errorf("could not ping if a new client can connect to the database: %s", err)
	}
	collection := client.Database("tournament").Collection("user")
	return &ModelMongo{Collection: collection} //, nil
}
