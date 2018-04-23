// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	Credentials
	Token
	SessionCredentials
	Username
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/adamsanghera/blurber-protobufs/dist/common"

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

type Credentials struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
}

func (m *Credentials) Reset()                    { *m = Credentials{} }
func (m *Credentials) String() string            { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()               {}
func (*Credentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Credentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Token struct {
	Token string `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SessionCredentials struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=Token" json:"Token,omitempty"`
}

func (m *SessionCredentials) Reset()                    { *m = SessionCredentials{} }
func (m *SessionCredentials) String() string            { return proto.CompactTextString(m) }
func (*SessionCredentials) ProtoMessage()               {}
func (*SessionCredentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SessionCredentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SessionCredentials) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Username struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
}

func (m *Username) Reset()                    { *m = Username{} }
func (m *Username) String() string            { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()               {}
func (*Username) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Username) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*Credentials)(nil), "Credentials")
	proto.RegisterType((*Token)(nil), "Token")
	proto.RegisterType((*SessionCredentials)(nil), "SessionCredentials")
	proto.RegisterType((*Username)(nil), "Username")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserDB service

type UserDBClient interface {
	Add(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*common.Empty, error)
	Delete(ctx context.Context, in *Username, opts ...grpc.CallOption) (*common.Empty, error)
	GetID(ctx context.Context, in *Username, opts ...grpc.CallOption) (*common.UserID, error)
	LogIn(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Token, error)
	CheckIn(ctx context.Context, in *SessionCredentials, opts ...grpc.CallOption) (*Token, error)
	CheckOut(ctx context.Context, in *SessionCredentials, opts ...grpc.CallOption) (*common.Empty, error)
}

type userDBClient struct {
	cc *grpc.ClientConn
}

func NewUserDBClient(cc *grpc.ClientConn) UserDBClient {
	return &userDBClient{cc}
}

func (c *userDBClient) Add(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/UserDB/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDBClient) Delete(ctx context.Context, in *Username, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/UserDB/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDBClient) GetID(ctx context.Context, in *Username, opts ...grpc.CallOption) (*common.UserID, error) {
	out := new(common.UserID)
	err := grpc.Invoke(ctx, "/UserDB/GetID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDBClient) LogIn(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/UserDB/LogIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDBClient) CheckIn(ctx context.Context, in *SessionCredentials, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/UserDB/CheckIn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDBClient) CheckOut(ctx context.Context, in *SessionCredentials, opts ...grpc.CallOption) (*common.Empty, error) {
	out := new(common.Empty)
	err := grpc.Invoke(ctx, "/UserDB/CheckOut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserDB service

type UserDBServer interface {
	Add(context.Context, *Credentials) (*common.Empty, error)
	Delete(context.Context, *Username) (*common.Empty, error)
	GetID(context.Context, *Username) (*common.UserID, error)
	LogIn(context.Context, *Credentials) (*Token, error)
	CheckIn(context.Context, *SessionCredentials) (*Token, error)
	CheckOut(context.Context, *SessionCredentials) (*common.Empty, error)
}

func RegisterUserDBServer(s *grpc.Server, srv UserDBServer) {
	s.RegisterService(&_UserDB_serviceDesc, srv)
}

func _UserDB_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDBServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserDB/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDBServer).Add(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDB_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDBServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserDB/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDBServer).Delete(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDB_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Username)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDBServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserDB/GetID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDBServer).GetID(ctx, req.(*Username))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDB_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDBServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserDB/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDBServer).LogIn(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDB_CheckIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDBServer).CheckIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserDB/CheckIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDBServer).CheckIn(ctx, req.(*SessionCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDB_CheckOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDBServer).CheckOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserDB/CheckOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDBServer).CheckOut(ctx, req.(*SessionCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserDB",
	HandlerType: (*UserDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _UserDB_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserDB_Delete_Handler,
		},
		{
			MethodName: "GetID",
			Handler:    _UserDB_GetID_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _UserDB_LogIn_Handler,
		},
		{
			MethodName: "CheckIn",
			Handler:    _UserDB_CheckIn_Handler,
		},
		{
			MethodName: "CheckOut",
			Handler:    _UserDB_CheckOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0xf0,
	0x94, 0x5c, 0xb9, 0xb8, 0x9d, 0x8b, 0x52, 0x53, 0x52, 0xf3, 0x4a, 0x32, 0x13, 0x73, 0x8a, 0x85,
	0xa4, 0xb8, 0x38, 0x42, 0x8b, 0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xe0, 0x7c, 0x90, 0x5c, 0x40, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04, 0x13,
	0x44, 0x0e, 0xc6, 0x57, 0x92, 0xe5, 0x62, 0x0d, 0xc9, 0xcf, 0x4e, 0xcd, 0x13, 0x12, 0x81, 0x32,
	0xa0, 0xba, 0x21, 0x1c, 0x25, 0x37, 0x2e, 0xa1, 0xe0, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0x3c, 0x62,
	0x2d, 0x83, 0x9b, 0xc3, 0x84, 0x6c, 0x8e, 0x1a, 0x17, 0x8a, 0x73, 0x70, 0xe9, 0x36, 0x3a, 0xcf,
	0xc8, 0xc5, 0x06, 0xe2, 0xb8, 0x38, 0x09, 0x49, 0x73, 0x31, 0x3b, 0xa6, 0xa4, 0x08, 0xf1, 0xe8,
	0x21, 0xd9, 0x2c, 0xc5, 0xa6, 0xe7, 0x9a, 0x5b, 0x50, 0x52, 0x29, 0x24, 0xcd, 0xc5, 0xe6, 0x92,
	0x9a, 0x93, 0x5a, 0x92, 0x2a, 0xc4, 0xa9, 0x07, 0xd3, 0x8c, 0x24, 0xc9, 0xea, 0x9e, 0x5a, 0xe2,
	0xe9, 0x82, 0x2c, 0xc7, 0x0e, 0x66, 0x7a, 0xba, 0x08, 0xc9, 0x72, 0xb1, 0xfa, 0xe4, 0xa7, 0x7b,
	0xe6, 0x61, 0x18, 0x0c, 0x09, 0x06, 0x35, 0x2e, 0x76, 0xe7, 0x8c, 0xd4, 0xe4, 0x6c, 0xcf, 0x3c,
	0x21, 0x61, 0x3d, 0x4c, 0xaf, 0xc3, 0xd5, 0xa9, 0x73, 0x71, 0x80, 0xd5, 0xf9, 0x97, 0x96, 0xe0,
	0x52, 0x08, 0x76, 0x4c, 0x12, 0x1b, 0x38, 0xba, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc,
	0x29, 0x9f, 0x66, 0xca, 0x01, 0x00, 0x00,
}
