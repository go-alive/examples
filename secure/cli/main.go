package main

import (
	"fmt"

	"github.com/go-alive/go-micro/client"
	"github.com/go-alive/go-micro/transport"

	hello "github.com/go-alive/examples/greeter/srv/proto/hello"

	"context"
)

func init() {
	client.DefaultClient.Init(
		client.Transport(
			transport.NewTransport(transport.Secure(true)),
		),
	)
}

func main() {
	cl := hello.NewSayService("go.micro.srv.greeter", client.DefaultClient)

	rsp, err := cl.Hello(context.TODO(), &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
