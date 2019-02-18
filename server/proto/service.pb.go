// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	User
	Product
	Done
	Response
	RequestDiscount
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	ID        int64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Birthdate string `protobuf:"bytes,2,opt,name=Birthdate" json:"Birthdate,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetBirthdate() string {
	if m != nil {
		return m.Birthdate
	}
	return ""
}

type Product struct {
	ID    int64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Price int64 `protobuf:"varint,2,opt,name=Price" json:"Price,omitempty"`
}

func (m *Product) Reset()                    { *m = Product{} }
func (m *Product) String() string            { return proto1.CompactTextString(m) }
func (*Product) ProtoMessage()               {}
func (*Product) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Product) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Product) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type Done struct {
	Status bool `protobuf:"varint,1,opt,name=Status" json:"Status,omitempty"`
}

func (m *Done) Reset()                    { *m = Done{} }
func (m *Done) String() string            { return proto1.CompactTextString(m) }
func (*Done) ProtoMessage()               {}
func (*Done) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Done) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type Response struct {
	Pct          float32 `protobuf:"fixed32,1,opt,name=Pct" json:"Pct,omitempty"`
	ValueInCents int64   `protobuf:"varint,2,opt,name=ValueInCents" json:"ValueInCents,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetPct() float32 {
	if m != nil {
		return m.Pct
	}
	return 0
}

func (m *Response) GetValueInCents() int64 {
	if m != nil {
		return m.ValueInCents
	}
	return 0
}

type RequestDiscount struct {
	ProductID int64 `protobuf:"varint,1,opt,name=ProductID" json:"ProductID,omitempty"`
	UserID    int64 `protobuf:"varint,2,opt,name=UserID" json:"UserID,omitempty"`
}

func (m *RequestDiscount) Reset()                    { *m = RequestDiscount{} }
func (m *RequestDiscount) String() string            { return proto1.CompactTextString(m) }
func (*RequestDiscount) ProtoMessage()               {}
func (*RequestDiscount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestDiscount) GetProductID() int64 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

func (m *RequestDiscount) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func init() {
	proto1.RegisterType((*User)(nil), "User")
	proto1.RegisterType((*Product)(nil), "Product")
	proto1.RegisterType((*Done)(nil), "Done")
	proto1.RegisterType((*Response)(nil), "Response")
	proto1.RegisterType((*RequestDiscount)(nil), "RequestDiscount")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	RegisterBirthdate(ctx context.Context, in *User, opts ...grpc.CallOption) (*Done, error)
	RegisterProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Done, error)
	GetDiscount(ctx context.Context, in *RequestDiscount, opts ...grpc.CallOption) (*Response, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) RegisterBirthdate(ctx context.Context, in *User, opts ...grpc.CallOption) (*Done, error) {
	out := new(Done)
	err := grpc.Invoke(ctx, "/Service/RegisterBirthdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) RegisterProduct(ctx context.Context, in *Product, opts ...grpc.CallOption) (*Done, error) {
	out := new(Done)
	err := grpc.Invoke(ctx, "/Service/RegisterProduct", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetDiscount(ctx context.Context, in *RequestDiscount, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/Service/GetDiscount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	RegisterBirthdate(context.Context, *User) (*Done, error)
	RegisterProduct(context.Context, *Product) (*Done, error)
	GetDiscount(context.Context, *RequestDiscount) (*Response, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_RegisterBirthdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RegisterBirthdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/RegisterBirthdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RegisterBirthdate(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_RegisterProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Product)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).RegisterProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/RegisterProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).RegisterProduct(ctx, req.(*Product))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDiscount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Service/GetDiscount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetDiscount(ctx, req.(*RequestDiscount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterBirthdate",
			Handler:    _Service_RegisterBirthdate_Handler,
		},
		{
			MethodName: "RegisterProduct",
			Handler:    _Service_RegisterProduct_Handler,
		},
		{
			MethodName: "GetDiscount",
			Handler:    _Service_GetDiscount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto1.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x6d, 0xd2, 0xa6, 0x4d, 0x46, 0x6b, 0xeb, 0x22, 0x52, 0x8a, 0x48, 0x59, 0x3d, 0xf4, 0x20,
	0x1b, 0x50, 0x3f, 0x40, 0x6a, 0xa0, 0xe4, 0x16, 0xb6, 0xe8, 0xc1, 0x5b, 0x4c, 0x07, 0x0d, 0x48,
	0xb6, 0xee, 0x4e, 0xfc, 0x00, 0xbf, 0x5c, 0x36, 0xd9, 0x36, 0xa8, 0xa7, 0x9d, 0x19, 0xde, 0x9b,
	0x7d, 0xef, 0x0d, 0x8c, 0x0d, 0xea, 0xaf, 0xb2, 0x40, 0xb1, 0xd3, 0x8a, 0x14, 0xbf, 0x87, 0xc1,
	0x93, 0x41, 0xcd, 0x4e, 0xc0, 0x4f, 0x93, 0x99, 0xb7, 0xf0, 0x96, 0x7d, 0xe9, 0xa7, 0x09, 0xbb,
	0x80, 0x68, 0x55, 0x6a, 0x7a, 0xdf, 0xe6, 0x84, 0x33, 0x7f, 0xe1, 0x2d, 0x23, 0xd9, 0x0d, 0x78,
	0x0c, 0xa3, 0x4c, 0xab, 0x6d, 0x5d, 0xd0, 0x3f, 0xe2, 0x19, 0x04, 0x99, 0x2e, 0x8b, 0x96, 0xd4,
	0x97, 0x6d, 0xc3, 0x2f, 0x61, 0x90, 0xa8, 0x0a, 0xd9, 0x39, 0x0c, 0x37, 0x94, 0x53, 0x6d, 0x1a,
	0x46, 0x28, 0x5d, 0xc7, 0x1f, 0x20, 0x94, 0x68, 0x76, 0xaa, 0x32, 0xc8, 0xa6, 0xd0, 0xcf, 0x0a,
	0x6a, 0x00, 0xbe, 0xb4, 0x25, 0xe3, 0x70, 0xfc, 0x9c, 0x7f, 0xd4, 0x98, 0x56, 0x8f, 0x58, 0x91,
	0x71, 0xab, 0x7f, 0xcd, 0xf8, 0x1a, 0x26, 0x12, 0x3f, 0x6b, 0x34, 0x94, 0x94, 0xa6, 0x50, 0x75,
	0x45, 0xd6, 0x83, 0x53, 0x79, 0x50, 0xd8, 0x0d, 0xac, 0x14, 0xeb, 0x3c, 0x4d, 0xdc, 0x3a, 0xd7,
	0xdd, 0x7e, 0x7b, 0x30, 0xda, 0xb4, 0x19, 0xb1, 0x2b, 0x38, 0x95, 0xf8, 0x56, 0x1a, 0x42, 0x7d,
	0x30, 0xcf, 0x02, 0x61, 0x91, 0xf3, 0x40, 0x58, 0x47, 0xbc, 0xc7, 0xae, 0xed, 0xcf, 0x2d, 0x68,
	0x1f, 0x4a, 0x28, 0x5c, 0xd5, 0xa1, 0x6e, 0xe0, 0x68, 0x8d, 0x9d, 0xb6, 0xa9, 0xf8, 0xa3, 0x76,
	0x1e, 0x89, 0x7d, 0x02, 0xbc, 0xb7, 0x9a, 0xbc, 0x8c, 0x63, 0x7b, 0x28, 0xd4, 0x71, 0x73, 0xa7,
	0xd7, 0x61, 0xf3, 0xdc, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xe6, 0xf9, 0xb0, 0xbf, 0x01,
	0x00, 0x00,
}