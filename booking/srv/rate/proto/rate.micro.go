// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/go-alive/examples/booking/srv/rate/proto/rate.proto

/*
Package rate is a generated protocol buffer package.

It is generated from these files:
	github.com/go-alive/examples/booking/srv/rate/proto/rate.proto

It has these top-level messages:
	Request
	Result
	RatePlan
	RoomType
*/
package rate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/go-alive/go-micro/client"
	server "github.com/go-alive/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Rate service

type RateService interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error)
}

type rateService struct {
	c           client.Client
	serviceName string
}

func NewRateService(serviceName string, c client.Client) RateService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "rate"
	}
	return &rateService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *rateService) GetRates(ctx context.Context, in *Request, opts ...client.CallOption) (*Result, error) {
	req := c.c.NewRequest(c.serviceName, "Rate.GetRates", in)
	out := new(Result)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rate service

type RateHandler interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(context.Context, *Request, *Result) error
}

func RegisterRateHandler(s server.Server, hdlr RateHandler, opts ...server.HandlerOption) {
	type rate interface {
		GetRates(ctx context.Context, in *Request, out *Result) error
	}
	type Rate struct {
		rate
	}
	h := &rateHandler{hdlr}
	s.Handle(s.NewHandler(&Rate{h}, opts...))
}

type rateHandler struct {
	RateHandler
}

func (h *rateHandler) GetRates(ctx context.Context, in *Request, out *Result) error {
	return h.RateHandler.GetRates(ctx, in, out)
}
