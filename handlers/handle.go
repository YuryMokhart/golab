package handlers

import (
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/gorilla/mux"
)

// HTTPController type.
type HTTPController struct {
	w http.ResponseWriter
	r http.Request
}

// Router registers a new route with a matcher.
func Router() (*mux.Router, error) {
	r := mux.NewRouter()
	r.HandleFunc("/users", controller.PrintUsers).Methods(http.MethodGet)
	r.HandleFunc("/user", controller.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", controller.FindUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods(http.MethodDelete)
	return r, nil
}
