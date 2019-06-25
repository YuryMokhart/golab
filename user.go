package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Balance string             `json:"balance" bson:"balance"`
}

// TODO: rewrite everything using REST API architecture.
// TODO: don't write real long functions! Start readind Robert Martin: Clean code.
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", createUser).Methods(http.MethodPost)
	r.HandleFunc("/users", getUsers).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}

// TODO: remove all the copy-paste code, find similar code and create new helper funcs.
// TODO: read about naming convention. https://golang.org/doc/effective_go.html#Getters
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []user
	w.Header().Set("Content-Type", "application/json")
	client, _ := dbConnector()

	collection := client.Database("tournament").Collection("user")
	ctx := r.Context()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "could not find users. Error: "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var oneUser user
		err = cursor.Decode(&oneUser)
		if err != nil {
			log.Fatalf("could not decode into oneUser in getUsers(): %s\n", err)
		}
		users = append(users, oneUser)
	}
	if err = cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "cursor error message: "` + err.Error() + `" }`))
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatalf("could not encode users in getUsers(): %s\n", err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, _ := dbConnector()
	var oneUser user
	err := json.NewDecoder(r.Body).Decode(&oneUser)
	if err != nil {
		log.Fatalf("could not decode into oneUser in createUser(): %s\n", err)
	}
	collection := client.Database("tournament").Collection("user")
	ctx := context.Background()
	result, _ := collection.InsertOne(ctx, oneUser)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatalf("could not decode oneUser into val in createUser(): %s\n", err)
	}
}

func dbConnector() (*mongo.Client, error) {
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
	return client, nil
}
