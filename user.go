package main

import(
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"time"
)

type user struct {
	id int64
	name string
	balance int64
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	collection := client.Database("tournament").Collection("user")

	user1 := user{1, "NameUser1", 15}
	insertResult, err := collection.InsertOne(context.TODO(), user1)
	if err != nil {
	    panic(err)
	}
	fmt.Println("Inserted document: ", insertResult.InsertedID)

	filter := bson.D{{"name", "NameUser1"}}

	var result user

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
	    panic(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted document %v\n", deleteResult)

	err = client.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}