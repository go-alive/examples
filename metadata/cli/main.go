package main

import (
	"fmt"

	hello "github.com/go-alive/examples/greeter/srv/proto/hello"
	"github.com/go-alive/go-micro"
	"github.com/go-alive/go-micro/metadata"

	"context"
)

func main() {
	service := micro.NewService()
	service.Init()

	cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"User": "john",
		"ID":   "1",
	})

	rsp, err := cl.Hello(ctx, &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
