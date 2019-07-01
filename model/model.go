package model

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UseMongoDB struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMongoDB(db *mongo.Database, collection *mongo.Collection) *UseMongoDB {
	return &UseMongoDB{
		db:         db,
		collection: collection,
	}
}

//PrintUsers gets all users.
func (um *UseMongoDB) PrintUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	cursor, err := um.collection.Find(ctx, bson.M{})
	if err != nil {
		controller.ErrorHelper(w, err, "could not find users. Error: ")
	}
	defer cursor.Close(ctx)
	users := RetrieveUsers(ctx, cursor, w)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatalf("could not encode users in printUsers(): %s\n", err)
	}
}

//RetrieveUsers retrieves a users and return them.
func RetrieveUsers(ctx context.Context, cursor *mongo.Cursor, w http.ResponseWriter) Users {
	var users Users
	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			controller.ErrorHelper(w, err, "could not decode into oneUser in printUsers(): ")
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		controller.ErrorHelper(w, err, "cursor error message: ")
	}
	return users
}

//CreateUser creates a user.
func (um *UseMongoDB) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		controller.ErrorHelper(w, err, "could not decode into oneUser in createUser(): ")
		return
	}
	ctx := r.Context()
	result, err := um.collection.InsertOne(ctx, user)
	if err != nil {
		controller.ErrorHelper(w, err, "could not insert user: ")
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		controller.ErrorHelper(w, err, "could not encode oneUser in createUser(): ")
		return
	}
}

//FindUser gets a user from the database.
func (um *UseMongoDB) FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		controller.ErrorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	ctx := r.Context()
	idDoc := bson.M{"_id": id}
	res := um.collection.FindOne(ctx, idDoc)
	if res.Err() != nil {
		controller.ErrorHelper(w, err, "could not find specific user: ")
		return
	}
	var user User
	err = res.Decode(&user)
	if err != nil {
		controller.ErrorHelper(w, err, "could not decode specific user: ")
		return
	}
	err = json.NewEncoder(w).Encode(user)
}

//DeleteUser deletes a specific user from the database.
func (um *UseMongoDB) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		controller.ErrorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	ctx := r.Context()
	idDoc := bson.M{"_id": id}
	_, err = um.collection.DeleteOne(ctx, idDoc)
	if err != nil {
		controller.ErrorHelper(w, err, "could not delete specific user: ")
		return
	}
}
