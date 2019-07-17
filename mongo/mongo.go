package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/YuryMokhart/golab/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ModelMongo represnts a model struct.
type ModelMongo struct {
	Collection *mongo.Collection
	ID         primitive.ObjectID
}

// CreateUser creates a user.
func (mm *ModelMongo) CreateUser(user entity.User) error {
	// TODO: you have context in http layer.
	// var mm ModelMongo
	// mm.Collection = DBConnect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := mm.Collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// PrintUsers prints users from the database.
func (mm *ModelMongo) PrintUsers() (entity.Users, error) {
	var users entity.Users
	// var mm ModelMongo
	// mm.Collection = DBConnect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cursor, err := mm.Collection.Find(ctx, bson.M{})
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
			return users, fmt.Errorf("could not decode current document from the database into user during retrieving users: %s", err)
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return users, fmt.Errorf("cursor error occured during retrieving users: %s", err)
	}
	return users, nil
}

// FindUser gets a user from the database.
func (mm *ModelMongo) FindUser() (entity.User, error) {
	var user entity.User
	// var mm ModelMongo
	// mm.Collection = DBConnect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	idDoc := bson.M{"_id": mm.ID}
	res := mm.Collection.FindOne(ctx, idDoc)
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
func (mm *ModelMongo) DeleteUser() error {
	// var mm ModelMongo
	// mm.Collection = DBConnect()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	idDoc := bson.M{"_id": mm.ID}
	_, err := mm.Collection.DeleteOne(ctx, idDoc)
	if err != nil {
		return err
	}
	return nil
}
