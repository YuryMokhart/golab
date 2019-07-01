package main

import (
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", controller.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users", controller.PrintUsers).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", controller.FindUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", controller.DeleteUser).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", r)

}
