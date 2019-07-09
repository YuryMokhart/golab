package main

import (
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/handlers"
)

func main() {
	r, err := handlers.Router()
	if err != nil {
		// TODO: don't forget to wrap your error
		log.Fatal(err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
