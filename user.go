package main

import(
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	// "time"
	"log"
)

type user struct {
	ID int64 `json:"ID" bson:"ID"`
	Name string `json:"Name "bson:"Name"`
	Balance int64
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("could not connect mongoDB to a new client: %s\n", err)
	} 

	ctx := context.Background()
	// ctx1, close := context.WithTimeout(ctx, 1*time.Second)
	// defer close()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("could not initialise the client: %s\n", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("could not ping the client: %s\n", err)
	}

	collection := client.Database("tournament").Collection("user")

	user1 := user {
		ID: 1, 
		Name: "NameUser1", 
		Balance: 15,
	}
	insertResult, err := collection.InsertOne(ctx, user1)
	if err != nil {
	    log.Fatalf("could not insert a user into collection: %s\n", err)
	}
	fmt.Println("Inserted document: ", insertResult.InsertedID)

	filter := bson.D{{"Name", "NameUser1"}}

	var result user

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
	    log.Fatalf("could not find a user in the collection: %s\n", err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("could not delete a user from the collection: %s\n", err)
	}
	fmt.Printf("Deleted document: %v\n", deleteResult)

	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatalf("could not dsconnect: %s\n", err)
	}
}