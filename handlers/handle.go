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

// TODO: all comments should be a sentence.
// TODO: think about exporting.
// HTTPHandler type.
type HTTPHandler struct {
	// TODO: your http layer knows about controller. Oh my God!
	h controller.ControllerStruct
}

// Router registers a new route with a matcher.
func Router() (*mux.Router, error) {
	r := mux.NewRouter()
	// TODO: change the naming chain: printHandler...
	r.HandleFunc("/users", printHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", postHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", findHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", deleteHandler).Methods(http.MethodDelete)

	return r, nil
}

func printHandler(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	// TODO: think about content type for errors.
	w.Header().Set("Content-Type", "application/json")
	// TODO: where is your error, man?
	users, err := hh.h.PrintUsers()
	if err != nil {
		helpers.ErrorHelper(w, err, "could not print a user")

		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		// TODO: keep the same error design.
		// TODO: you return errors for user, not for a programmer.
		helpers.ErrorHelper(w, err, "couldn't encode users in printHandler")
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.ErrorHelper(w, err, "couldn't decode a user in postHandler")
		return
	}
	result, err := hh.h.CreateUser(user)
	if err != nil {
		helpers.ErrorHelper(w, err, "could not create a user")
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		helpers.ErrorHelper(w, err, "could not encode result in postHandler")
		return
	}
}

func findHandler(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		helpers.ErrorHelper(w, err, "hex string is not valid ObjectID in findHandler")
		return
	}
	user, err := hh.h.FindUser(id)
	if err != nil {
		helpers.ErrorHelper(w, err, "could not find a user")
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		helpers.ErrorHelper(w, err, "couldn't encode users in findHandler")
		return
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		helpers.ErrorHelper(w, err, "hex string is not valid ObjectID in deleteHandler")
		return
	}
	err = hh.h.DeleteUser(id)
	if err != nil {
		helpers.ErrorHelper(w, err, "could not delete a user")
		return
	}
}
