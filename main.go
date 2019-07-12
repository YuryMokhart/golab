package main

import (
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/handlers"
	"github.com/YuryMokhart/golab/mongo"
)

func main() {
	collection := mongo.DBConnect()
	mod := mongo.ModelMongo{Collection: collection}
	contr := controller.ControllerStruct{M: mod}
	httphandler := handlers.HTTPHandler{H: contr}
	// var modelInterface controller.Modeller
	// var controllerInterface controller.Controller
	// trying, ok := modelInterface.(controllerInterface)
	r, err := handlers.Router(httphandler)
	if err != nil {
		log.Fatalf("could not register a new route %s", err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not listen and serve %s", err)
	}
}
