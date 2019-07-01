package main

import (
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/model"
	"github.com/gorilla/mux"
)

func main() {
	db, col := model.DBConnect()
	_ = model.NewMongoDB(db, col)
	r := mux.NewRouter()
	r.HandleFunc("/users", controller.PrintUsers).Methods(http.MethodGet)
	r.HandleFunc("/user", controller.CreateUser).Methods(http.MethodPost)
	// r.HandleFunc("/user/{id}", controller.FindUser).Methods(http.MethodGet)
	// r.HandleFunc("/user/{id}", controller.DeleteUser).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", r)

}
