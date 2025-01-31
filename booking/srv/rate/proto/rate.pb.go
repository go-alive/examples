// Code generated by protoc-gen-go. DO NOT EDIT.
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

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	InDate   string   `protobuf:"bytes,2,opt,name=inDate" json:"inDate,omitempty"`
	OutDate  string   `protobuf:"bytes,3,opt,name=outDate" json:"outDate,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type Result struct {
	RatePlans []*RatePlan `protobuf:"bytes,1,rep,name=ratePlans" json:"ratePlans,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetRatePlans() []*RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

type RatePlan struct {
	HotelId  string    `protobuf:"bytes,1,opt,name=hotelId" json:"hotelId,omitempty"`
	Code     string    `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	InDate   string    `protobuf:"bytes,3,opt,name=inDate" json:"inDate,omitempty"`
	OutDate  string    `protobuf:"bytes,4,opt,name=outDate" json:"outDate,omitempty"`
	RoomType *RoomType `protobuf:"bytes,5,opt,name=roomType" json:"roomType,omitempty"`
}

func (m *RatePlan) Reset()                    { *m = RatePlan{} }
func (m *RatePlan) String() string            { return proto.CompactTextString(m) }
func (*RatePlan) ProtoMessage()               {}
func (*RatePlan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RatePlan) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *RatePlan) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RatePlan) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *RatePlan) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

func (m *RatePlan) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

type RoomType struct {
	BookableRate       float64 `protobuf:"fixed64,1,opt,name=bookableRate" json:"bookableRate,omitempty"`
	TotalRate          float64 `protobuf:"fixed64,2,opt,name=totalRate" json:"totalRate,omitempty"`
	TotalRateInclusive float64 `protobuf:"fixed64,3,opt,name=totalRateInclusive" json:"totalRateInclusive,omitempty"`
	Code               string  `protobuf:"bytes,4,opt,name=code" json:"code,omitempty"`
	Currency           string  `protobuf:"bytes,5,opt,name=currency" json:"currency,omitempty"`
	RoomDescription    string  `protobuf:"bytes,6,opt,name=roomDescription" json:"roomDescription,omitempty"`
}

func (m *RoomType) Reset()                    { *m = RoomType{} }
func (m *RoomType) String() string            { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()               {}
func (*RoomType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RoomType) GetBookableRate() float64 {
	if m != nil {
		return m.BookableRate
	}
	return 0
}

func (m *RoomType) GetTotalRate() float64 {
	if m != nil {
		return m.TotalRate
	}
	return 0
}

func (m *RoomType) GetTotalRateInclusive() float64 {
	if m != nil {
		return m.TotalRateInclusive
	}
	return 0
}

func (m *RoomType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RoomType) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *RoomType) GetRoomDescription() string {
	if m != nil {
		return m.RoomDescription
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rate.Request")
	proto.RegisterType((*Result)(nil), "rate.Result")
	proto.RegisterType((*RatePlan)(nil), "rate.RatePlan")
	proto.RegisterType((*RoomType)(nil), "rate.RoomType")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rate service

type RateClient interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type rateClient struct {
	cc *grpc.ClientConn
}

func NewRateClient(cc *grpc.ClientConn) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/rate.Rate/GetRates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rate service

type RateServer interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(context.Context, *Request) (*Result, error)
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate.Rate/GetRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServer).GetRates(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rate.Rate",
	HandlerType: (*RateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _Rate_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/go-alive/examples/booking/srv/rate/proto/rate.proto",
}

func init() {
	proto.RegisterFile("github.com/go-alive/examples/booking/srv/rate/proto/rate.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0xc9, 0x6f, 0xfd, 0x75, 0xed, 0xe3, 0x54, 0x78, 0x0e, 0x52, 0x86, 0x87, 0xd1, 0x8b,
	0x43, 0xa4, 0x85, 0x09, 0x5e, 0xbc, 0x0e, 0x64, 0x37, 0x09, 0x82, 0xe7, 0xb6, 0x0b, 0x5b, 0x31,
	0x6d, 0x6a, 0x92, 0x0e, 0xf7, 0x46, 0x7c, 0x69, 0xbe, 0x1e, 0x49, 0x9a, 0x76, 0x9b, 0xe8, 0xed,
	0xfb, 0x27, 0x3c, 0xfd, 0xe4, 0x69, 0xe0, 0x71, 0x53, 0xea, 0x6d, 0x9b, 0x27, 0x85, 0xa8, 0xd2,
	0xaa, 0x2c, 0xa4, 0x48, 0xd9, 0x47, 0x56, 0x35, 0x9c, 0xa9, 0x34, 0x17, 0xe2, 0xad, 0xac, 0x37,
	0xa9, 0x92, 0xbb, 0x54, 0x66, 0x9a, 0xa5, 0x8d, 0x14, 0x5a, 0x58, 0x99, 0x58, 0x89, 0x9e, 0xd1,
	0xf1, 0x2b, 0x8c, 0x29, 0x7b, 0x6f, 0x99, 0xd2, 0x38, 0x85, 0x60, 0x2b, 0x34, 0xe3, 0xab, 0xb5,
	0x8a, 0xc8, 0x6c, 0x34, 0x0f, 0xe9, 0xe0, 0xf1, 0x0a, 0xfc, 0xb2, 0x5e, 0x66, 0x9a, 0x45, 0xff,
	0x66, 0x64, 0x1e, 0x52, 0xe7, 0x30, 0x82, 0xb1, 0x68, 0xb5, 0x2d, 0x46, 0xb6, 0xe8, 0x6d, 0xfc,
	0x00, 0x3e, 0x65, 0xaa, 0xe5, 0x1a, 0xef, 0x20, 0x34, 0x9f, 0x7a, 0xe6, 0x59, 0xdd, 0x0d, 0x3e,
	0x5b, 0x5c, 0x24, 0x16, 0x84, 0xba, 0x98, 0x1e, 0x0e, 0xc4, 0x9f, 0x04, 0x82, 0x3e, 0x37, 0xe3,
	0x1d, 0x42, 0x44, 0xba, 0xf1, 0xce, 0x22, 0x82, 0x57, 0x88, 0x75, 0x8f, 0x63, 0xf5, 0x11, 0xe4,
	0xe8, 0x2f, 0x48, 0xef, 0x04, 0x12, 0x6f, 0x21, 0x90, 0x42, 0x54, 0x2f, 0xfb, 0x86, 0x45, 0xff,
	0x67, 0xe4, 0x88, 0xcc, 0xa5, 0x74, 0xe8, 0xe3, 0x2f, 0x03, 0xe6, 0x0c, 0xc6, 0x30, 0x31, 0x1b,
	0xce, 0x72, 0xce, 0x0c, 0xac, 0xa5, 0x23, 0xf4, 0x24, 0xc3, 0x6b, 0x08, 0xb5, 0xd0, 0x19, 0xa7,
	0xfd, 0xda, 0x08, 0x3d, 0x04, 0x98, 0x00, 0x0e, 0x66, 0x55, 0x17, 0xbc, 0x55, 0xe5, 0xae, 0x03,
	0x27, 0xf4, 0x97, 0x66, 0xb8, 0xb0, 0x77, 0x74, 0xe1, 0x29, 0x04, 0x45, 0x2b, 0x25, 0xab, 0x8b,
	0xbd, 0xc5, 0x0f, 0xe9, 0xe0, 0x71, 0x0e, 0x97, 0x06, 0x7d, 0xc9, 0x54, 0x21, 0xcb, 0x46, 0x97,
	0xa2, 0x8e, 0x7c, 0x7b, 0xe4, 0x67, 0xbc, 0x48, 0xc1, 0xb3, 0x44, 0x37, 0x10, 0x3c, 0x31, 0x6d,
	0xa4, 0xc2, 0x73, 0xb7, 0x86, 0xee, 0x69, 0x4c, 0x27, 0xbd, 0x35, 0x3f, 0x34, 0xf7, 0xed, 0x03,
	0xba, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xc6, 0x91, 0x7f, 0x02, 0x00, 0x00,
}
