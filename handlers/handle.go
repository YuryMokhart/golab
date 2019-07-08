package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/entity"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HTTPController type.
type HTTPController struct {
	w http.ResponseWriter
	r *http.Request
}

// Router registers a new route with a matcher.
func Router() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/users", handler).Methods(http.MethodGet)
	r.HandleFunc("/user", handler).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handler).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", handler).Methods(http.MethodDelete)
	return r, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		var user entity.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			controller.ErrorHelper(w, err, "couldn't encode user in createUser")
			return
		}
		result := controller.CreateUser(&user)
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			controller.ErrorHelper(w, err, "could not encode oneUser in createUser(): ")
			return
		}
	} else if r.Method == http.MethodGet && r.URL.Path == "/users" {
		users := controller.PrintUsers()
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			controller.ErrorHelper(w, err, "couldn't encode users in printUsers")
			return
		}
	} else if r.Method == http.MethodGet && r.URL.Path == "/user/{id}" {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			controller.ErrorHelper(w, err, "hex string is not valid ObjectID: ")
			return
		}
		user := controller.FindUser(id)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			controller.ErrorHelper(w, err, "couldn't encode users in findUsers")
			return
		}
	} else if r.Method == http.MethodDelete {
		vars := mux.Vars(r)
		id, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			controller.ErrorHelper(w, err, "hex string is not valid ObjectID: ")
			return
		}
		controller.DeleteUser(id)
	}
}
