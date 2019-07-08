package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct represents what fields user type should contain.
type User struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Balance string             `json:"balance" bson:"balance"`
}

// Users is a slice of type user.
type Users []User
