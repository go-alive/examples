package web

import (
	"github.com/go-alive/go-micro"
	"github.com/go-alive/go-micro/web"
	"github.com/go-alive/go-plugins/registry/kubernetes/v2"
	// static selector offloads load balancing to k8s services
	// enable with MICRO_SELECTOR=static or --selector=static
	// requires user to create k8s services
	"github.com/go-alive/go-plugins/client/selector/static/v2"
)

// NewService returns a web service for kubernetes
func NewService(opts ...web.Option) web.Service {
	// setup
	k := kubernetes.NewRegistry()
	st := static.NewSelector()

	// create new service
	service := micro.NewService(
		micro.Registry(k),
		micro.Selector(st),
	)

	// prepend option
	options := []web.Option{
		web.MicroService(service),
	}

	options = append(options, opts...)

	// return new service
	return web.NewService(options...)
}
