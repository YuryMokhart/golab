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
	CreateUser(entity.User) (*mongoDriver.InsertOneResult, error)
	PrintUsers() (entity.Users, error)
	FindUser(primitive.ObjectID) (entity.User, error)
	DeleteUser(primitive.ObjectID) error
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
func (c ControllerStruct) PrintUsers() (entity.Users, error) {
	users, err := c.m.PrintUsers()
	if err != nil {
		return nil, fmt.Errorf("could not print users: %s", err)
	}
	return users, nil
}

// FindUser finds a specific user by id.
func (c ControllerStruct) FindUser(id primitive.ObjectID) (entity.User, error) {
	user, err := c.m.FindUser(id)
	if err != nil {
		return user, fmt.Errorf("could not find a users: %s", err)
	}
	return user, nil
}

// DeleteUser deletes a specific user.
func (c ControllerStruct) DeleteUser(id primitive.ObjectID) error {
	err := c.m.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("could not delete a users: %s", err)
	}
	return nil
}
