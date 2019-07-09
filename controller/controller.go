package controller

import (
	"fmt"

	"github.com/YuryMokhart/golab/entity"
	"github.com/YuryMokhart/golab/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

// TODO: you need that interface, but not in the controller.
type controller interface {
	// TODO: controller layer should do not know about mongo package.
	CreateUser(entity.User) *mongoDriver.InsertOneResult
	PrintUsers(entity.Users) entity.Users
	FindUser(primitive.ObjectID) entity.User
	DeleteUser(primitive.ObjectID)
}

// ControllerStruct struct.
type ControllerStruct struct {
	m mongo.ModelMongo
}

// CreateUser creates a user.
func (c ControllerStruct) CreateUser(user entity.User) (*mongoDriver.InsertOneResult, error) {
	result, err := c.m.CreateUser(&user)
	if err != nil {
		// TODO: do the same with error in other controllers.
		return nil, fmt.Errorf("could not create a new user: %s", err)
	}
	return result, nil
}

// PrintUsers returns all users from the database.
func (c ControllerStruct) PrintUsers(users entity.Users) entity.Users {
	users, _ = c.m.PrintUsers()
	return users
}

// FindUser finds a specific user by id.
func (c ControllerStruct) FindUser(id primitive.ObjectID) entity.User {
	user, _ := c.m.FindUser(id)
	return user
}

// DeleteUser deletes a specific user.
func (c ControllerStruct) DeleteUser(id primitive.ObjectID) {
	c.m.DeleteUser(id)
}
