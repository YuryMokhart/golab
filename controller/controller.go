package controller

import (
	"encoding/json"
	"net/http"

	"github.com/YuryMokhart/golab/entity"
	"github.com/YuryMokhart/golab/mongo"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Controller interface.
type Controller interface {
	CreateUser(http.ResponseWriter, *http.Request)
	PrintUsers(http.ResponseWriter, *http.Request)
	FindUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

// CreateUser creates a user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorHelper(w, err, "couldn't encode user in createUser")
		return
	}
	result, err := mongo.CreateUser(&user)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		ErrorHelper(w, err, "could not encode oneUser in createUser(): ")
		return
	}
}

// PrintUsers prints all users.
func PrintUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users entity.Users
	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		ErrorHelper(w, err, "couldn't encode users in printUsers")
		return
	}
	users, err = mongo.PrintUsers()
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		ErrorHelper(w, err, "couldn't encode users in printUsers")
		return
	}
}

// FindUser finds a specific user by id.
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		ErrorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	user, err := mongo.FindUser(id)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		ErrorHelper(w, err, "couldn't encode users in findUsers")
		return
	}
}

// DeleteUser deletes a specific user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		ErrorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	err = mongo.DeleteUser(id)
	if err != nil {
		return
	}
}
