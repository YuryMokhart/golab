package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/entity"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HTTPHandler represents HTTPHandler struct.
type HTTPHandler struct {
	// TODO: your http layer knows about controller. Oh my God!
	H controller.ControllerStruct
}

// Router registers a new route with a matcher.
func Router() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/users", printHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", postHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", findHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", deleteHandler).Methods(http.MethodDelete)
	// r.Handle("/users", httphandler)
	// r.Handle("/user/{id}", httphandler).Methods(http.MethodGet)
	// r.Handle("/user/{id}", httphandler).Methods(http.MethodDelete)
	// r.Handle("/user", httphandler)

	return r, nil
}

// // TODO: you don't need it.
// func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/users" && r.Method == http.MethodGet {
// 		h.printHandler(w, r)
// 	} else if r.URL.Path == "/user/{id}" && r.Method == http.MethodGet {
// 		h.findHandler(w, r)
// 	} else if r.URL.Path == "/user/{id}" && r.Method == http.MethodDelete {
// 		h.deleteHandler(w, r)
// 	} else if r.URL.Path == "/user" && r.Method == http.MethodPost {
// 		h.postHandler(w, r)
// 	}
// 	// TODO: what will be here?
// }

func printHandler(w http.ResponseWriter, r *http.Request) {
	var h HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	users, err := h.H.PrintUsers()
	if err != nil {
		errorHelper(w, err, "could not print a user")
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		errorHelper(w, err, "could not print users due to the problem with encoding")
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var h HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorHelper(w, err, "could not send a user due to the problem with decoding it")
		return
	}
	result, err := h.H.CreateUser(user)
	if err != nil {
		errorHelper(w, err, "could not create a user")
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		errorHelper(w, err, "could not print the result due to the problem with encoding")
		return
	}
}

func findHandler(w http.ResponseWriter, r *http.Request) {
	var h HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		errorHelper(w, err, "id is not valid")
		return
	}
	user, err := h.H.FindUser(id)
	if err != nil {
		errorHelper(w, err, "could not find a user")
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		errorHelper(w, err, "could not print a found user due to the problem with encoding ")
		return
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	var h HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		errorHelper(w, err, "id is not valid")
		return
	}
	err = h.H.DeleteUser(id)
	if err != nil {
		errorHelper(w, err, "could not delete a user")
		return
	}
}

func errorHelper(w http.ResponseWriter, err error, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	type customError struct {
		Msg string `json:"msg"`
		Err error  `json:"err"`
	}
	ce := customError{
		Msg: message,
		Err: err,
	}
	err = json.NewEncoder(w).Encode(ce)
	if err != nil {
		fmt.Printf("could not write an error due to the problem with encoding %s", err)
		return
	}
	return
}
