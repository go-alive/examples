// Package micro implements a go-micro service for k8s
package micro

import (
	"github.com/go-alive/go-micro"
	"github.com/go-alive/go-plugins/registry/kubernetes/v2"

	// static selector offloads load balancing to k8s services
	// note: requires user to create k8s services
	"github.com/go-alive/go-plugins/client/selector/static/v2"
)

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...micro.Option) micro.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()
	s := static.NewSelector()

	// set the registry and selector
	options := []micro.Option{
		micro.Registry(r),
		micro.Selector(s),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return micro.NewService(options...)
}
