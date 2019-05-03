// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package models

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

type UserID struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_32c7b97edc90262b, []int{0}
}
func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (dst *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(dst, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type InfoUser struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PhotoUUID            string   `protobuf:"bytes,3,opt,name=photoUUID,proto3" json:"photoUUID,omitempty"`
	Active               bool     `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoUser) Reset()         { *m = InfoUser{} }
func (m *InfoUser) String() string { return proto.CompactTextString(m) }
func (*InfoUser) ProtoMessage()    {}
func (*InfoUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_32c7b97edc90262b, []int{1}
}
func (m *InfoUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoUser.Unmarshal(m, b)
}
func (m *InfoUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoUser.Marshal(b, m, deterministic)
}
func (dst *InfoUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoUser.Merge(dst, src)
}
func (m *InfoUser) XXX_Size() int {
	return xxx_messageInfo_InfoUser.Size(m)
}
func (m *InfoUser) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoUser.DiscardUnknown(m)
}

var xxx_messageInfo_InfoUser proto.InternalMessageInfo

func (m *InfoUser) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *InfoUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *InfoUser) GetPhotoUUID() string {
	if m != nil {
		return m.PhotoUUID
	}
	return ""
}

func (m *InfoUser) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type SessionToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionToken) Reset()         { *m = SessionToken{} }
func (m *SessionToken) String() string { return proto.CompactTextString(m) }
func (*SessionToken) ProtoMessage()    {}
func (*SessionToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_32c7b97edc90262b, []int{2}
}
func (m *SessionToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionToken.Unmarshal(m, b)
}
func (m *SessionToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionToken.Marshal(b, m, deterministic)
}
func (dst *SessionToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionToken.Merge(dst, src)
}
func (m *SessionToken) XXX_Size() int {
	return xxx_messageInfo_SessionToken.Size(m)
}
func (m *SessionToken) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionToken.DiscardUnknown(m)
}

var xxx_messageInfo_SessionToken proto.InternalMessageInfo

func (m *SessionToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SessionPayload struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionPayload) Reset()         { *m = SessionPayload{} }
func (m *SessionPayload) String() string { return proto.CompactTextString(m) }
func (*SessionPayload) ProtoMessage()    {}
func (*SessionPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_32c7b97edc90262b, []int{3}
}
func (m *SessionPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionPayload.Unmarshal(m, b)
}
func (m *SessionPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionPayload.Marshal(b, m, deterministic)
}
func (dst *SessionPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionPayload.Merge(dst, src)
}
func (m *SessionPayload) XXX_Size() int {
	return xxx_messageInfo_SessionPayload.Size(m)
}
func (m *SessionPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionPayload.DiscardUnknown(m)
}

var xxx_messageInfo_SessionPayload proto.InternalMessageInfo

func (m *SessionPayload) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *SessionPayload) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type UserIDs struct {
	IDs                  []*UserID `protobuf:"bytes,1,rep,name=IDs,proto3" json:"IDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserIDs) Reset()         { *m = UserIDs{} }
func (m *UserIDs) String() string { return proto.CompactTextString(m) }
func (*UserIDs) ProtoMessage()    {}
func (*UserIDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_32c7b97edc90262b, []int{4}
}
func (m *UserIDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIDs.Unmarshal(m, b)
}
func (m *UserIDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIDs.Marshal(b, m, deterministic)
}
func (dst *UserIDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIDs.Merge(dst, src)
}
func (m *UserIDs) XXX_Size() int {
	return xxx_messageInfo_UserIDs.Size(m)
}
func (m *UserIDs) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIDs.DiscardUnknown(m)
}

var xxx_messageInfo_UserIDs proto.InternalMessageInfo

func (m *UserIDs) GetIDs() []*UserID {
	if m != nil {
		return m.IDs
	}
	return nil
}

type InfoUsers struct {
	Users                []*InfoUser `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *InfoUsers) Reset()         { *m = InfoUsers{} }
func (m *InfoUsers) String() string { return proto.CompactTextString(m) }
func (*InfoUsers) ProtoMessage()    {}
func (*InfoUsers) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_32c7b97edc90262b, []int{5}
}
func (m *InfoUsers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoUsers.Unmarshal(m, b)
}
func (m *InfoUsers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoUsers.Marshal(b, m, deterministic)
}
func (dst *InfoUsers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoUsers.Merge(dst, src)
}
func (m *InfoUsers) XXX_Size() int {
	return xxx_messageInfo_InfoUsers.Size(m)
}
func (m *InfoUsers) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoUsers.DiscardUnknown(m)
}

var xxx_messageInfo_InfoUsers proto.InternalMessageInfo

func (m *InfoUsers) GetUsers() []*InfoUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*UserID)(nil), "models.UserID")
	proto.RegisterType((*InfoUser)(nil), "models.InfoUser")
	proto.RegisterType((*SessionToken)(nil), "models.SessionToken")
	proto.RegisterType((*SessionPayload)(nil), "models.SessionPayload")
	proto.RegisterType((*UserIDs)(nil), "models.UserIDs")
	proto.RegisterType((*InfoUsers)(nil), "models.InfoUsers")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	GetUserByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*InfoUser, error)
	GetSessionInfo(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*SessionPayload, error)
	GetUsersByIDs(ctx context.Context, in *UserIDs, opts ...grpc.CallOption) (*InfoUsers, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetUserByID(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*InfoUser, error) {
	out := new(InfoUser)
	err := c.cc.Invoke(ctx, "/models.Auth/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetSessionInfo(ctx context.Context, in *SessionToken, opts ...grpc.CallOption) (*SessionPayload, error) {
	out := new(SessionPayload)
	err := c.cc.Invoke(ctx, "/models.Auth/GetSessionInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUsersByIDs(ctx context.Context, in *UserIDs, opts ...grpc.CallOption) (*InfoUsers, error) {
	out := new(InfoUsers)
	err := c.cc.Invoke(ctx, "/models.Auth/GetUsersByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	GetUserByID(context.Context, *UserID) (*InfoUser, error)
	GetSessionInfo(context.Context, *SessionToken) (*SessionPayload, error)
	GetUsersByIDs(context.Context, *UserIDs) (*InfoUsers, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Auth/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserByID(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetSessionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetSessionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Auth/GetSessionInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetSessionInfo(ctx, req.(*SessionToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUsersByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUsersByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.Auth/GetUsersByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUsersByIDs(ctx, req.(*UserIDs))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByID",
			Handler:    _Auth_GetUserByID_Handler,
		},
		{
			MethodName: "GetSessionInfo",
			Handler:    _Auth_GetSessionInfo_Handler,
		},
		{
			MethodName: "GetUsersByIDs",
			Handler:    _Auth_GetUsersByIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_32c7b97edc90262b) }

var fileDescriptor_auth_32c7b97edc90262b = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0xa5, 0xeb, 0xd6, 0xaf, 0xbd, 0xfb, 0xac, 0x7a, 0x19, 0xa3, 0x14, 0x1f, 0x46, 0x10, 0x19,
	0x08, 0x15, 0xd6, 0x57, 0x11, 0x94, 0x81, 0xf4, 0x4d, 0xa2, 0xfd, 0x01, 0xd5, 0x45, 0x36, 0xec,
	0x9a, 0xb1, 0x74, 0xc2, 0xfe, 0x93, 0x3f, 0xd2, 0x9b, 0xa6, 0x51, 0x57, 0x7d, 0x6a, 0xee, 0x39,
	0xf7, 0xf4, 0xe4, 0x1c, 0x02, 0x50, 0xec, 0xea, 0x65, 0xb2, 0xd9, 0xca, 0x5a, 0xa2, 0xb7, 0x96,
	0x0b, 0x51, 0x2a, 0x16, 0x81, 0x97, 0x2b, 0xb1, 0xcd, 0xe6, 0x18, 0x42, 0x2f, 0x9b, 0x47, 0xce,
	0xc4, 0x99, 0xba, 0x9c, 0x4e, 0xac, 0x04, 0x3f, 0xab, 0x5e, 0xa5, 0x66, 0xbb, 0x1c, 0xc6, 0xe0,
	0xef, 0x08, 0xaf, 0x8a, 0xb5, 0x88, 0x7a, 0x84, 0x06, 0xfc, 0x6b, 0xc6, 0x33, 0x08, 0x36, 0x4b,
	0xb2, 0xc8, 0x73, 0x92, 0xb8, 0x0d, 0xf9, 0x0d, 0xe0, 0x18, 0xbc, 0xe2, 0xa5, 0x5e, 0xbd, 0x8b,
	0xa8, 0x4f, 0x94, 0xcf, 0xdb, 0x89, 0x9d, 0xc3, 0xff, 0x47, 0xa1, 0xd4, 0x4a, 0x56, 0x4f, 0xf2,
	0x4d, 0x54, 0x38, 0x82, 0x41, 0xad, 0x0f, 0x8d, 0x69, 0xc0, 0xcd, 0xc0, 0xae, 0x21, 0x6c, 0xb7,
	0x1e, 0x8a, 0x7d, 0x29, 0x8b, 0xc5, 0x5f, 0x37, 0xcb, 0x3b, 0x37, 0xb3, 0x33, 0xbb, 0x84, 0x7f,
	0x26, 0xab, 0xc2, 0x09, 0xb8, 0xf4, 0x21, 0x9d, 0x3b, 0x1d, 0xce, 0xc2, 0xc4, 0x94, 0x91, 0x18,
	0x96, 0x6b, 0x8a, 0xa5, 0x10, 0xd8, 0xf8, 0x0a, 0x2f, 0x60, 0xa0, 0xf3, 0x59, 0xc1, 0x89, 0x15,
	0xd8, 0x0d, 0x6e, 0xe8, 0xd9, 0x87, 0x03, 0xfd, 0x5b, 0x2a, 0x19, 0xaf, 0x60, 0x78, 0x2f, 0x6a,
	0x4d, 0xdd, 0xed, 0x75, 0xb7, 0x87, 0x0e, 0xf1, 0xaf, 0x1f, 0xe0, 0x0d, 0x84, 0x24, 0x68, 0xc3,
	0x69, 0x14, 0x47, 0x76, 0xe7, 0x67, 0x2f, 0xf1, 0xb8, 0x83, 0xda, 0x1e, 0x52, 0x38, 0x6a, 0x0d,
	0x95, 0x76, 0x54, 0x78, 0x7c, 0x68, 0xa9, 0xe2, 0xd3, 0xae, 0xa7, 0x7a, 0xf6, 0x9a, 0xb7, 0x90,
	0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x5c, 0xd3, 0x85, 0x19, 0x02, 0x00, 0x00,
}
