package main

import (
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/controller"

	"github.com/YuryMokhart/golab/handlers"
	"github.com/YuryMokhart/golab/mongo"
)

func main() {
	var model mongo.ModelMongo
	model.Collection = mongo.DBConnect()
	var controller controller.Control
	controller.M = model
	var handler handlers.HTTPHandler
	handler.H = controller
	r, err := handlers.Router(handler)
	if err != nil {
		log.Fatalf("could not register a new route %s", err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not listen and serve %s", err)
	}
}
