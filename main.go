package main

import (
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/handlers"
)

func main() {
	// collection := mongo.DBConnect()
	r, err := handlers.Router()
	if err != nil {
		log.Fatalf("could not register a new route %s", err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not listen and serve %s", err)
	}
}
