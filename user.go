package main

// TODO: fix GOROOT and GOPATH.
import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "time"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	ID      string `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Balance int64  `json:"balance" bson:"balance"`
}

// TODO: you don't need global variables.
// TODO: rewrite everything using REST API architecture.
// TODO: don't write real long functions! Start readind Robert Martin: Clean code.
func main() {
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
		log.Fatalf("could not ping the client: %s\n", err)
	}

	collection := client.Database("tournament").Collection("user")

	user1 := user{
		ID:      "1",
		Name:    "NameUser1",
		Balance: 15,
	}
	insertResult, err := collection.InsertOne(ctx, user1)
	if err != nil {
		log.Fatalf("could not insert a user into collection: %s\n", err)
	}
	fmt.Println("Inserted document: ", insertResult.InsertedID)

	filter := bson.D{{"name", "NameUser1"}}

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
		log.Fatalf("could not disconnect: %s\n", err)
	}
	var users []user
	r := mux.NewRouter()
	users = append(users, user{
		ID:      "34",
		Name:    "Andrew",
		Balance: 10000,
	})
	// TODO: use http constans for GET, POST and etc
	r.HandleFunc("/users", getUsers).Methods(http.MethodGet)
	// r.HandleFunc("/users/{_id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	// r.HandleFunc("/users/{_id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// TODO: remove all the copy-paste code, find similar code and create new helper funcs.
// TODO: read about naming convention. https://golang.org/doc/effective_go.html#Getters
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []user
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		// TODO: remember about panics and Fatalf.
		log.Fatalf("could not connect mongoDB to a new client: %s\n", err)
	}
	collection := client.Database("tournament").Collection("user")
	// TODO: you don't need create new context. You have context in http.Request (r)
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatalf("could not find users: %s\n", err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var oneUser user
		// TODO: error check
		cursor.Decode(&oneUser)
		users = append(users, oneUser)
	}
	// TODO: remember to check all the errors.
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("could not connect mongoDB to a new client: %s\n", err)
	}
	var oneUser user
	_ = json.NewDecoder(r.Body).Decode(&oneUser)
	collection := client.Database("tournament").Collection("user")
	ctx := context.Background()
	result, _ := collection.InsertOne(ctx, oneUser)
	json.NewEncoder(w).Encode(result)
}

// func getUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// 	params := mux.Vars(r)
// 	for _, specUser := range users {
// 		if specUser.ID == params["_id"] {
// 			json.NewEncoder(w).Encode(specUser)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&user{})
// }

// func deleteUser(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     params := mux.Vars(r)
//     for index, specUser := range users {
//         if specUser.ID == params["_id"] {
//             users = append(users[:index], users[index+1:]...)
//             break
//         }
//     }
//     json.NewEncoder(w).Encode(users)
// }
func dbConnector()(*Client, *Collection, error){
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("could not connect mongoDB to a new client: %s\n", err)
	}
	collection := client.Database("tournament").Collection("user")
	return client, collection, err
}