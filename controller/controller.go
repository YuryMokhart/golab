package controller

import (
	"fmt"

	"github.com/YuryMokhart/golab/entity"
	"github.com/YuryMokhart/golab/mongo"
)

// Modeller is a model interface.
type Modeller interface {
	CreateUser(*entity.User) error
	PrintUsers() (entity.Users, error)
	FindUser() (entity.User, error)
	DeleteUser() error
}

// Control represents a controller struct.
type Control struct {
	M mongo.ModelMongo
}

// CreateUser creates a user.
func (c Control) CreateUser(user entity.User) error {
	err := c.M.CreateUser(user)
	if err != nil {
		return fmt.Errorf("could not create a new user: %s", err)
	}
	return nil
}

// PrintUsers returns all users from the database.
func (c Control) PrintUsers() (entity.Users, error) {
	users, err := c.M.PrintUsers()
	if err != nil {
		return nil, fmt.Errorf("could not print users: %s", err)
	}
	return users, nil
}

// FindUser finds a specific user by id.
func (c Control) FindUser() (entity.User, error) {
	user, err := c.M.FindUser()
	if err != nil {
		return user, fmt.Errorf("could not find a users: %s", err)
	}
	return user, nil
}

// DeleteUser deletes a specific user.
func (c Control) DeleteUser() error {
	err := c.M.DeleteUser()
	if err != nil {
		return fmt.Errorf("could not delete a users: %s", err)
	}
	return nil
}
