package main

import (
	"encoding/json"
	"log"
	"strings"

	hello "github.com/go-alive/examples/greeter/srv/proto/hello"
	"github.com/go-alive/go-micro"
	"github.com/go-alive/go-micro/errors"
	shard "github.com/go-alive/go-plugins/wrapper/select/shard/v2"
	api "github.com/go-alive/micro/api/proto"

	"context"
)

type Say struct {
	Client hello.SayService
}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.greeter", "Name cannot be blank")
	}

	response, err := s.Client.Hello(ctx, &hello.Request{
		Name: strings.Join(name.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Msg,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	wrapper := shard.NewClientWrapper("X-From-User")

	service := micro.NewService(
		micro.Name("go.micro.api.greeter"),
		micro.WrapClient(wrapper),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(service.Server().NewHandler(&Say{
		Client: hello.NewSayService("go.micro.srv.greeter", service.Client()),
	}))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
