package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/entity"
	"github.com/gorilla/mux"
)

// HTTPHandler type.
type HTTPHandler struct {
	h controller.ControllerStruct
}

// Router registers a new route with a matcher.
func Router() (*mux.Router, error) {
	r := mux.NewRouter()
	// r.HandleFunc("/users", controller.PrintUsers).Methods(http.MethodGet)
	// r.HandleFunc("/user", controller.CreateUser).Methods(http.MethodPost)
	// r.HandleFunc("/user/{id}", controller.FindUser).Methods(http.MethodGet)
	// r.HandleFunc("/user/{id}", controller.DeleteUser).Methods(http.MethodDelete)
	r.Handle("/", http.HandlerFunc(handler))
	return r, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var hh HTTPHandler
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path == "/users" && r.Method == http.MethodGet {
		var users entity.Users
		err := json.NewEncoder(w).Encode(&users)
		if err != nil {
			controller.ErrorHelper(w, err, "couldn't encode users in printUsers")
			return
		}
		users = hh.h.PrintUsers(users)
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			controller.ErrorHelper(w, err, "couldn't encode users in printUsers")
			return
		}
	}
}
