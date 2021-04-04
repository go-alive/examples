package subscriber

import (
	"context"

	"github.com/go-alive/go-micro/util/log"

	example "github.com/go-alive/examples/template/fnc/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
