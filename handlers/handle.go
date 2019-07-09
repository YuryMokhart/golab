package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/entity"
	"github.com/YuryMokhart/golab/helpers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HTTPHandler type.
type HTTPHandler struct {
	h controller.ControllerStruct
}

// Router registers a new route with a matcher.
func Router() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlerPrint).Methods(http.MethodGet)
	r.HandleFunc("/user", handlerPost).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handlerFind).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handlerDelete).Methods(http.MethodDelete)
	return r, nil
}

func handlerPrint(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	var users entity.Users
	users = hh.h.PrintUsers(users)
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		helpers.ErrorHelper(w, err, "couldn't encode users in handler.")
		return
	}
}

func handlerPost(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.ErrorHelper(w, err, "couldn't encode user in createUser")
		return
	}
	result := hh.h.CreateUser(user)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		helpers.ErrorHelper(w, err, "could not encode oneUser in createUser(): ")
		return
	}
}

func handlerFind(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		helpers.ErrorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	user := hh.h.FindUser(id)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		helpers.ErrorHelper(w, err, "couldn't encode users in findUsers")
		return
	}
}

func handlerDelete(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		helpers.ErrorHelper(w, err, "hex string is not valid ObjectID: ")
		return
	}
	hh.h.DeleteUser(id)
}
