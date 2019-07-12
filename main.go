package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/handlers"
	"github.com/YuryMokhart/golab/mongo"
)

func main() {
	collection := mongo.DBConnect()
	var mi mongo.Modeller
	model, ok := mi.(mongo.ModelMongo)
	model.Collection = collection
	fmt.Println(model, ok)
	var ci controller.Controller
	control, ok := ci.(controller.ControllerStruct)
	control.M = model
	fmt.Println(control, ok)
	httphandler := handlers.HTTPHandler{H: control}
	fmt.Println(httphandler)
	r, err := handlers.Router(httphandler)
	if err != nil {
		log.Fatalf("could not register a new route %s", err)
	}
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("could not listen and serve %s", err)
	}
}
