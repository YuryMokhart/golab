package main

import (
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/handler"
	"github.com/YuryMokhart/golab/mongo"
)

func main() {
	db := mongo.DBConnect()
	controller := controller.New(db)
	httphandler := handler.New(controller)
	r, err := handler.Router(httphandler)
	if err != nil {
		log.Fatalf("could not register a new route %s", err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not listen and serve %s", err)
	}
}

// TODO: cover with unit tests all the code.
