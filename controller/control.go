package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//PrintUsers gets all users.
func PrintUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := model.DBConnect()
	ctx := r.Context()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		ErrorHelper(w, err, "could not find users. Error: ")
	}
	defer cursor.Close(ctx)
	users := RetrieveUsers(ctx, cursor, w)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatalf("could not encode users in printUsers(): %s\n", err)
	}
}

//RetrieveUsers retrieves a users and return them.
func RetrieveUsers(ctx context.Context, cursor *mongo.Cursor, w http.ResponseWriter) model.Users {
	var users model.Users
	for cursor.Next(ctx) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			ErrorHelper(w, err, "could not decode into oneUser in printUsers(): ")
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		ErrorHelper(w, err, "cursor error message: ")
	}
	return users
}

//CreateUser creates a user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := model.DBConnect()
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorHelper(w, err, "could not decode into oneUser in createUser(): ")
		return
	}
	ctx := r.Context()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		ErrorHelper(w, err, "could not insert user: ")
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		ErrorHelper(w, err, "could not encode oneUser in createUser(): ")
		return
	}
}

//FindUser gets a user from the database.
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := model.DBConnect()
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		ErrorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	ctx := r.Context()
	idDoc := bson.M{"_id": id}
	res := collection.FindOne(ctx, idDoc)
	if res.Err() != nil {
		ErrorHelper(w, err, "could not find specific user: ")
		return
	}
	var user model.User
	err = res.Decode(&user)
	if err != nil {
		ErrorHelper(w, err, "could not decode specific user: ")
		return
	}
	err = json.NewEncoder(w).Encode(user)
}

//DeleteUser deletes a specific user from the database.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection := model.DBConnect()
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		ErrorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	ctx := r.Context()
	idDoc := bson.M{"_id": id}
	_, err = collection.DeleteOne(ctx, idDoc)
	if err != nil {
		ErrorHelper(w, err, "could not delete specific user: ")
		return
	}
}
