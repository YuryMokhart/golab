package main

import (
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/handler"
)

func main() {
	r, err := handler.Handle()
	if err != nil {
		log.Fatal(err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
