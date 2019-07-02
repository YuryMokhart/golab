package model

import (
	"context"

	"github.com/YuryMokhart/golab/controller"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UseMongoDB struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

//PrintUsers gets all users.
func (c controller.HTTPController) PrintUsers(users Users) Users {
	ctx := context.Background()
	cursor, err := c.Collection.Find(ctx, bson.M{})
	if err != nil {
		//controller.ErrorHelper(w, err, "could not find users. Error: ")
	}
	defer cursor.Close(ctx)
	users = RetrieveUsers(ctx, cursor)
	return users
}

//RetrieveUsers retrieves users and return them.
func RetrieveUsers(ctx context.Context, cursor *mongo.Cursor) Users {
	var users Users
	for cursor.Next(ctx) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			//controller.ErrorHelper(w, err, "could not decode into oneUser in printUsers(): ")
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		//controller.ErrorHelper(w, err, "cursor error message: ")
	}
	return users
}

//CreateUser creates a user.
func (c HTTPController) CreateUser(user User) *mongo.InsertOneResult {
	ctx := context.Background()
	result, err := c.Collection.InsertOne(ctx, user)
	if err != nil {
		//controller.ErrorHelper(w, err, "could not insert user: ")
		//return
	}
	return result
}

// //FindUser gets a user from the database.
// func (um *UseMongoDB) FindUser(user User, id ObjectID) User {
// 	ctx := context.Background()
// 	idDoc := bson.M{"_id": id}
// 	res := um.collection.FindOne(ctx, idDoc)
// 	if res.Err() != nil {
// 		//controller.ErrorHelper(w, err, "could not find specific user: ")
// 		//return
// 	}
// 	err := res.Decode(&user)
// 	if err != nil {
// 		//controller.ErrorHelper(w, err, "could not decode specific user: ")
// 		//return
// 	}
// 	return user
// }

// //DeleteUser deletes a specific user from the database.
// func (um *UseMongoDB) DeleteUser(user User, id ObjectID) {
// 	ctx := context.Background()
// 	idDoc := bson.M{"_id": id}
// 	_, err := um.collection.DeleteOne(ctx, idDoc)
// 	if err != nil {
// 		//controller.ErrorHelper(w, err, "could not delete specific user: ")
// 		return
// 	}
// }
