package main

import (
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/model"
	"github.com/gorilla/mux"
)

func main() {
	db, col := model.DBConnect()
	c := controller.HTTPController(db, col)
	r := mux.NewRouter()
	r.HandleFunc("/users", c.PrintUsers).Methods(http.MethodGet)
	r.HandleFunc("/user", c.CreateUser).Methods(http.MethodPost)
	// r.HandleFunc("/user/{id}", controller.FindUser).Methods(http.MethodGet)
	// r.HandleFunc("/user/{id}", controller.DeleteUser).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", r)

}
