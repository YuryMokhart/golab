package mongo

import (
	"context"
	"time"

	"github.com/YuryMokhart/golab/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoModeller interface {
	CreateUser(entity.User) (*mongo.InsertOneResult, error)
	PrintUsers() (entity.Users, error)
	FindUser(primitive.ObjectID) (entity.User, error)
	DeleteUser(primitive.ObjectID) error
}

// CreateUser creates a user.
func CreateUser(user *entity.User) (*mongo.InsertOneResult, error) {
	collection := DBConnect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PrintUsers prints users from the database.
func PrintUsers() (entity.Users, error) {
	var users entity.Users
	collection := DBConnect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}
	defer cursor.Close(ctx)
	users, err = retrieveUsers(ctx, cursor)
	return users, err
}

// retrieveUsers retrieves users and return them.
func retrieveUsers(ctx context.Context, cursor *mongo.Cursor) (entity.Users, error) {
	var users entity.Users
	for cursor.Next(ctx) {
		var user entity.User
		err := cursor.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return users, err
	}
	return users, nil
}

// FindUser gets a user from the database.
func FindUser(id primitive.ObjectID) (entity.User, error) {
	var user entity.User
	collection := DBConnect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	idDoc := bson.M{"_id": id}
	res := collection.FindOne(ctx, idDoc)
	if res.Err() != nil {
		return user, res.Err()
	}
	err := res.Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// DeleteUser deletes a specific user from the database.
func DeleteUser(id primitive.ObjectID) error {
	collection := DBConnect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	idDoc := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, idDoc)
	if err != nil {
		return err
	}
	return nil
}
