package mongo

import (
	"context"

	"github.com/YuryMokhart/golab/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoModeller interface {
	createUser(entity.User) *mongo.InsertOneResult
	printUsers() entity.Users
	// findUser()
	// deleteUser()
}

// CreateUser creates a user.
func CreateUser(user *entity.User) *mongo.InsertOneResult {
	collection := DBConnect()
	ctx := context.Background()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		// controller.ErrorHelper(w, err, "could not insert user: ")
		// return
	}
	return result
}

// PrintUsers prints users from the database.
func PrintUsers() entity.Users {
	collection := DBConnect()
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		//controller.ErrorHelper(w, err, "could not find users. Error: ")
	}
	defer cursor.Close(ctx)
	users := retrieveUsers(ctx, cursor)
	return users
}

// retrieveUsers retrieves users and return them.
func retrieveUsers(ctx context.Context, cursor *mongo.Cursor) entity.Users {
	var users entity.Users
	for cursor.Next(ctx) {
		var user entity.User
		err := cursor.Decode(&user)
		if err != nil {
			// controller.ErrorHelper(w, err, "could not decode into oneUser in printUsers(): ")
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		// controller.ErrorHelper(w, err, "cursor error message: ")
	}
	return users
}

// FindUser gets a user from the database.
func FindUser(id primitive.ObjectID) entity.User {
	collection := DBConnect()
	ctx := context.Background()
	idDoc := bson.M{"_id": id}
	res := collection.FindOne(ctx, idDoc)
	if res.Err() != nil {
		//controller.ErrorHelper(w, err, "could not find specific user: ")
		//return
	}
	var user entity.User
	err := res.Decode(&user)
	if err != nil {
		//controller.ErrorHelper(w, err, "could not decode specific user: ")
		//return
	}
	return user
}

// DeleteUser deletes a specific user from the database.
func DeleteUser(id primitive.ObjectID) {
	collection := DBConnect()
	ctx := context.Background()
	idDoc := bson.M{"_id": id}
	_, err := collection.DeleteOne(ctx, idDoc)
	if err != nil {
		//controller.ErrorHelper(w, err, "could not delete specific user: ")
		return
	}
}
