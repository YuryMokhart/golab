package controller

import (
	"net/http"

	"github.com/YuryMokhart/golab/entity"
	mongodb "github.com/YuryMokhart/golab/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Controller interface.
type Controller interface {
	CreateUser(http.ResponseWriter, *http.Request)
	PrintUsers(http.ResponseWriter, *http.Request)
	FindUser(http.ResponseWriter, *http.Request)
	DeleteUser(http.ResponseWriter, *http.Request)
}

// CreateUser creates a user.
func CreateUser(user *entity.User) *mongo.InsertOneResult {
	result, err := mongodb.CreateUser(&user)
	if err != nil {
		// return nil
	}
	return result
}

// PrintUsers prints all users.
func PrintUsers() *entity.Users {
	users, err := mongodb.PrintUsers()
	if err != nil {
		// return
	}
	return &users
}

// FindUser finds a specific user by id.
func FindUser(id primitive.ObjectID) entity.User {
	user, err := mongodb.FindUser(&id)
	if err != nil {
		// return nil
	}
	return user
}

// DeleteUser deletes a specific user.
func DeleteUser(id primitive.ObjectID) {
	err := mongodb.DeleteUser(id)
	if err != nil {
		// return
	}
}
