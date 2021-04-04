package main

import (
	"os"

	"github.com/go-alive/go-micro/broker"
	"github.com/go-alive/go-micro/client"
	cli "github.com/go-alive/go-micro/client/grpc"
	"github.com/go-alive/go-micro/config/cmd"
	"github.com/go-alive/go-micro/server"
	srv "github.com/go-alive/go-micro/server/grpc"
	bkr "github.com/go-alive/go-plugins/broker/grpc/v2"
	_ "github.com/go-alive/go-plugins/registry/kubernetes/v2"

	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	_ "github.com/go-alive/go-plugins/client/selector/static/v2"

	// disable namespace by default
	_ "github.com/go-alive/go-micro/api"
)

func main() {

	// set values for registry/selector
	os.Setenv("MICRO_REGISTRY", "kubernetes")
	os.Setenv("MICRO_SELECTOR", "static")

	// setup broker/client/server
	broker.DefaultBroker = bkr.NewBroker()
	client.DefaultClient = cli.NewClient()
	server.DefaultServer = srv.NewServer()

	// init command
	cmd.Init()
}
