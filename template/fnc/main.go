package main

import (
	"github.com/go-alive/examples/template/fnc/handler"
	"github.com/go-alive/examples/template/fnc/subscriber"
	"github.com/go-alive/go-micro"
	"github.com/go-alive/go-micro/util/log"
)

func main() {
	// New Service
	function := micro.NewFunction(
		micro.Name("go.micro.fnc.template"),
		micro.Version("latest"),
	)

	// Register Handler
	function.Handle(new(handler.Example))

	// Register Struct as Subscriber
	function.Subscribe("go.micro.fnc.template", new(subscriber.Example))

	// Initialise function
	function.Init()

	// Run service
	if err := function.Run(); err != nil {
		log.Fatal(err)
	}
}
