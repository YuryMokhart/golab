package main

import (
	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/mongo"
)

func main() {
	db := mongo.DBConnect()
	controller := controller.New(db)
	_ = controller
	/*handler := &handlers.HTTPHandler{C: controller}
	r, err := handlers.Router(handler)
	if err != nil {
		log.Fatalf("could not register a new route %s", err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not listen and serve %s", err)
	}*/
}

// TODO: cover with unit tests all the code.
