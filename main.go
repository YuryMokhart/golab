package main

import (
	"log"
	"net/http"

	"github.com/YuryMokhart/golab/controller"
	"github.com/YuryMokhart/golab/handler"
	"github.com/YuryMokhart/golab/mongo"
)

func main() {
	db, err := mongo.DBConnect()
	if err != nil {
		log.Fatal(err)
	}
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

// TODO: cover with unit tests all the code with mock testing. >95% coverage.
// TODO: you can watch video about unit testing by justForFunc.
// TODO: add Swagger file as HTTP documentation for your project. Use OpenAPI version of Swagger (3.0). Write the documentation yourserlf, don't use libraties.
// TODO: add go modules as dependence manager of you packages and read about the reason why we need it.
// TODO: add Makefile, README file and .gitignore file. Read about format and the reasons why it's needed.
// TODO: read about linters and start using golangci-lint for your project, add it to Makefile. Fix all warnings by golangci-lint.
// TODO: read about Dockerfiles and create Docker container for your project.
