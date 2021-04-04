// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/go-alive/examples/service/proto/greeter.proto

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/go-alive/examples/service/proto/greeter.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/Greeter/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Greeter/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Greeter_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/go-alive/examples/service/proto/greeter.proto",
}

func init() {
	proto.RegisterFile("github.com/go-alive/examples/service/proto/greeter.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xcd, 0x0a, 0xc2, 0x30,
	0x10, 0x84, 0xad, 0xf8, 0xbb, 0x58, 0x0f, 0x39, 0x95, 0x9e, 0x24, 0xa7, 0x82, 0x90, 0x80, 0xc5,
	0x67, 0xd0, 0x73, 0xdf, 0xa0, 0x2d, 0x43, 0x0d, 0x34, 0x4d, 0x4d, 0x52, 0xf1, 0xf1, 0x85, 0x55,
	0x44, 0x6f, 0x3b, 0x30, 0xfb, 0x7d, 0x43, 0xe7, 0xce, 0xc4, 0xdb, 0xd4, 0xa8, 0xd6, 0x59, 0x6d,
	0x4d, 0xeb, 0x9d, 0xc6, 0xb3, 0xb6, 0x63, 0x8f, 0xa0, 0x03, 0xfc, 0xc3, 0xb4, 0xd0, 0xa3, 0x77,
	0xd1, 0xe9, 0xce, 0x03, 0x11, 0x5e, 0x71, 0x92, 0x92, 0x76, 0x57, 0xf4, 0xbd, 0xab, 0x70, 0x9f,
	0x10, 0xa2, 0x10, 0xb4, 0x18, 0x6a, 0x8b, 0x2c, 0x39, 0x24, 0xc5, 0xb6, 0xe2, 0x5b, 0x1e, 0x29,
	0xfd, 0x74, 0xc2, 0xe8, 0x86, 0x00, 0x91, 0xd3, 0x86, 0x29, 0x66, 0xe8, 0xb2, 0x39, 0x17, 0xbf,
	0xf9, 0x54, 0xd2, 0xfa, 0xf2, 0x36, 0x88, 0x82, 0x96, 0xfc, 0x27, 0x52, 0xf5, 0xeb, 0xc8, 0xf7,
	0xea, 0x0f, 0x27, 0x67, 0xcd, 0x8a, 0xc7, 0x94, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xa3,
	0x46, 0x4c, 0xc5, 0x00, 0x00, 0x00,
}
