package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/YuryMokhart/golab/entity"
	"github.com/gorilla/mux"
)

// Controller represents what methods it should contain.
type Controller interface {
	CreateUser(entity.User) error
	PrintUsers() (entity.Users, error)
	FindUser() (entity.User, error)
	DeleteUser() error
}

// HTTPHandler represents HTTPHandler struct.
type HTTPHandler struct {
	BLogic Controller
}

// New creates a new object of httphandler.
func New(c Controller) HTTPHandler {
	return HTTPHandler{BLogic: c}
}

// Router registers a new route with a matcher.
func Router(h HTTPHandler) (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/users", h.printHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", h.postHandler).Methods(http.MethodPost)
	// r.HandleFunc("/user/{id}", h.findHandler).Methods(http.MethodGet)
	// r.HandleFunc("/user/{id}", h.deleteHandler).Methods(http.MethodDelete)

	return r, nil
}

func (h HTTPHandler) printHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.PrintUsers()
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

// PrintUsers prints all users fom the database.
func (h HTTPHandler) PrintUsers() (entity.Users, error) {
	users, err := h.BLogic.PrintUsers()
	if err != nil {
		return nil, fmt.Errorf("could not print a user: %s", err)
	}
	return users, err
}

func (h HTTPHandler) postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorHelper(w, err, "could not send a user due to the problem with decoding it")
		return
	}
	err = h.CreateUser(user)
	if err != nil {
		errorHelper(w, err, "could not create a user")
		return
	}
}

// CreateUser creates a user and adds it into the database.
func (h HTTPHandler) CreateUser(user entity.User) error {
	err := h.BLogic.CreateUser(user)
	if err != nil {
		return fmt.Errorf("could not create a user: %s", err)
	}
	return err
}

func (h HTTPHandler) findHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// vars := mux.Vars(r)
	// id, err := primitive.ObjectIDFromHex(vars["id"])
	// h.BLogic.
	// h.C.M.ID = id
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorHelper(w, err, "could not read from the request")
	}
	err = json.Unmarshal()
	if err != nil {
		errorHelper(w, err, "id is not valid")
		return
	}
	user, err := h.FindUser()
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

// FindUser finds a specific user in the database.
func (h HTTPHandler) FindUser() (entity.User, error) {
	user, err := h.BLogic.FindUser()
	return user, err
}

// func (h HTTPHandler) deleteHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	vars := mux.Vars(r)
// 	id, err := primitive.ObjectIDFromHex(vars["id"])
// 	h.C.M.ID = id
// 	if err != nil {
// 		errorHelper(w, err, "id is not valid")
// 		return
// 	}
// 	err = h.C.DeleteUser()
// 	if err != nil {
// 		errorHelper(w, err, "could not delete a user")
// 		return
// 	}
// }

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
