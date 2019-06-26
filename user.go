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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", createUser).Methods(http.MethodPost)
	r.HandleFunc("/users", printUsers).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", findUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", deleteUser).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", r)

}

func printUsers(w http.ResponseWriter, r *http.Request) {
	var users []user
	w.Header().Set("Content-Type", "application/json")
	client, _ := dbConnector()

	collection := client.Database("tournament").Collection("user")
	ctx := r.Context()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		errorHelper(w, err, "could not find users. Error: ")
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var oneUser user
		err = cursor.Decode(&oneUser)
		if err != nil {
			log.Fatalf("could not decode into oneUser in printUsers(): %s\n", err)
		}
		users = append(users, oneUser)
	}
	if err = cursor.Err(); err != nil {
		errorHelper(w, err, "cursor error message: ")
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatalf("could not encode users in printUsers(): %s\n", err)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, _ := dbConnector()
	var oneUser user
	err := json.NewDecoder(r.Body).Decode(&oneUser)
	if err != nil {
		errorHelper(w, err, "could not decode into oneUser in createUser(): ")
	}
	collection := client.Database("tournament").Collection("user")
	ctx := r.Context()
	result, err := collection.InsertOne(ctx, oneUser)
	if err != nil {
		errorHelper(w, err, "could not insert user: ")
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		errorHelper(w, err, "could not encode oneUser in createUser(): ")
	}
}

func findUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, _ := dbConnector()
	ctx := r.Context()
	collection := client.Database("tournament").Collection("user")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		errorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	var oneUser user
	idDoc := bson.M{"_id": id}
	res := collection.FindOne(ctx, idDoc)
	if res.Err() != nil {
		errorHelper(w, err, "could not find specific user: ")
		return
	}
	err = res.Decode(&oneUser)
	if err != nil {
		errorHelper(w, err, "could not decode specific user: ")
		return
	}
	err = json.NewEncoder(w).Encode(oneUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, _ := dbConnector()
	ctx := r.Context()
	collection := client.Database("tournament").Collection("user")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		errorHelper(w, err, "hex string is not valid ObjectID: ")
	}
	idDoc := bson.M{"_id": id}
	_, err = collection.DeleteOne(ctx, idDoc)
	if err != nil {
		errorHelper(w, err, "could not delete specific user: ")
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

func errorHelper(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{ ` + message + ` ` + err.Error() + `" }`))
	return
}
