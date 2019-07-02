package controller

import (
	"encoding/json"
	"net/http"

	"github.com/YuryMokhart/golab/model"
)

type HTTPController struct {
	model model.UseMongoDB
}

func (c HTTPController) PrintUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users model.Users
	err := json.NewEncoder(w).Encode(&users)
	if err != nil {
		ErrorHelper(w, err, "couldn't encode users in printUsers")
	}
	users = c.model.PrintUsers(users)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		ErrorHelper(w, err, "couldn't encode users in printUsers")
	}
}

func (c HTTPController) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	err := json.NewEncoder(w).Encode(&user)
	if err != nil {
		ErrorHelper(w, err, "couldn't encode user in createUser")
	}
	result := c.model.CreateUser(user)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		ErrorHelper(w, err, "could not encode oneUser in createUser(): ")
		return
	}
}

// func FindUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user model.User
// 	vars := mux.Vars(r)
// 	id, err := primitive.ObjectIDFromHex(vars["id"])
// 	if err != nil {
// 		ErrorHelper(w, err, "hex string is not valid ObjectID: ")
// 		return
// 	}
// 	// user.ID := bson.NewObjectId()
// 	user = model.FindUser(user, id)
// 	err = json.NewEncoder(w).Encode(user)
// }

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user model.User
// 	vars := mux.Vars(r)
// 	id, err := primitive.ObjectIDFromHex(vars["id"])
// 	if err != nil {
// 		ErrorHelper(w, err, "hex string is not valid ObjectID: ")
// 		return
// 	}
// 	// user.ID := bson.NewObjectId()
// 	model.DeleteUser(user, id)
// }
