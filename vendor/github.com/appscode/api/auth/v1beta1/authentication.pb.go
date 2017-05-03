// Code generated by protoc-gen-go.
// source: authentication.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	authentication.proto
	conduit.proto
	project.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	CSRFTokenResponse
	ConduitWhoAmIResponse
	ConduitUsersResponse
	ConduitUser
	Preferences
	ProjectListRequest
	ProjectListResponse
	ProjectMemberListRequest
	ProjectMemberListResponse
	Project
	Member
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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

type LoginRequest struct {
	Namespace  string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Token      string `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	IssueToken bool   `protobuf:"varint,5,opt,name=issue_token,json=issueToken" json:"issue_token,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRequest) GetIssueToken() bool {
	if m != nil {
		return m.IssueToken
	}
	return false
}

type LoginResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Token  string                  `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CSRFTokenResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	CsrfToken string                  `protobuf:"bytes,2,opt,name=csrf_token,json=csrfToken" json:"csrf_token,omitempty"`
}

func (m *CSRFTokenResponse) Reset()                    { *m = CSRFTokenResponse{} }
func (m *CSRFTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CSRFTokenResponse) ProtoMessage()               {}
func (*CSRFTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CSRFTokenResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CSRFTokenResponse) GetCsrfToken() string {
	if m != nil {
		return m.CsrfToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "appscode.auth.v1beta1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "appscode.auth.v1beta1.LoginResponse")
	proto.RegisterType((*CSRFTokenResponse)(nil), "appscode.auth.v1beta1.CSRFTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authentication service

type AuthenticationClient interface {
	// This rpc is used to check a valid user from other applications.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	CSRFToken(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*CSRFTokenResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Authentication/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Logout(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Authentication/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) CSRFToken(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*CSRFTokenResponse, error) {
	out := new(CSRFTokenResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Authentication/CSRFToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authentication service

type AuthenticationServer interface {
	// This rpc is used to check a valid user from other applications.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *appscode_dtypes.VoidRequest) (*appscode_dtypes.VoidResponse, error)
	CSRFToken(context.Context, *appscode_dtypes.VoidRequest) (*CSRFTokenResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Authentication/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Logout(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_CSRFToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).CSRFToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Authentication/CSRFToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).CSRFToken(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.auth.v1beta1.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Authentication_Logout_Handler,
		},
		{
			MethodName: "CSRFToken",
			Handler:    _Authentication_CSRFToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xe5, 0x94, 0x84, 0x64, 0x0a, 0x48, 0x98, 0x22, 0x56, 0x51, 0x42, 0xab, 0x05, 0x41,
	0x44, 0xe9, 0x5a, 0x2d, 0x07, 0xfe, 0x5c, 0x10, 0x45, 0xe2, 0xd4, 0x43, 0xb5, 0x45, 0x3d, 0x70,
	0x89, 0xdc, 0x8d, 0x49, 0x17, 0x5a, 0x8f, 0xc9, 0xd8, 0x20, 0x4e, 0x48, 0xbd, 0x22, 0x71, 0xe1,
	0xc8, 0x4b, 0x70, 0xe1, 0x49, 0x78, 0x05, 0x1e, 0x04, 0xd9, 0xbb, 0xd9, 0xee, 0xaa, 0xa4, 0x95,
	0xb8, 0xac, 0xe4, 0xf9, 0x79, 0xfc, 0x7d, 0xe3, 0x6f, 0x0d, 0x2b, 0xd2, 0xd9, 0x43, 0xa5, 0x6d,
	0x9e, 0x49, 0x9b, 0xa3, 0x4e, 0xcc, 0x0c, 0x2d, 0xf2, 0x9b, 0xd2, 0x18, 0xca, 0x70, 0xa2, 0x12,
	0x8f, 0x93, 0x8f, 0x9b, 0x07, 0xca, 0xca, 0xcd, 0xfe, 0x60, 0x8a, 0x38, 0x3d, 0x52, 0x42, 0x9a,
	0x5c, 0x48, 0xad, 0xd1, 0x86, 0x1e, 0x2a, 0x9a, 0xfa, 0xb7, 0xe7, 0x4d, 0x0b, 0xf8, 0x6a, 0x83,
	0x4f, 0xec, 0x67, 0xa3, 0x48, 0x84, 0x6f, 0xb1, 0x21, 0xfe, 0xc1, 0xe0, 0xca, 0x0e, 0x4e, 0x73,
	0x9d, 0xaa, 0x0f, 0x4e, 0x91, 0xe5, 0x03, 0xe8, 0x69, 0x79, 0xac, 0xc8, 0xc8, 0x4c, 0x45, 0x6c,
	0x8d, 0x8d, 0x7a, 0xe9, 0x69, 0x81, 0xf7, 0xa1, 0xeb, 0x48, 0xcd, 0x7c, 0x21, 0x6a, 0x05, 0x58,
	0xad, 0x3d, 0x33, 0x92, 0xe8, 0x13, 0xce, 0x26, 0xd1, 0x52, 0xc1, 0xe6, 0x6b, 0xbe, 0x02, 0x6d,
	0x8b, 0xef, 0x95, 0x8e, 0x2e, 0x05, 0x50, 0x2c, 0xf8, 0x2a, 0x2c, 0xe7, 0x44, 0x4e, 0x8d, 0x0b,
	0xd6, 0x5e, 0x63, 0xa3, 0x6e, 0x0a, 0xa1, 0xf4, 0xda, 0x57, 0xe2, 0x7d, 0xb8, 0x5a, 0x9a, 0x23,
	0x83, 0x9a, 0x14, 0x17, 0xd0, 0x21, 0x2b, 0xad, 0xa3, 0x60, 0x6d, 0x79, 0xeb, 0x56, 0x52, 0xdd,
	0x5a, 0x31, 0x5c, 0xb2, 0x17, 0x70, 0x5a, 0x6e, 0x3b, 0x15, 0x6e, 0xd5, 0x84, 0xe3, 0x0c, 0xae,
	0xbf, 0xdc, 0x4b, 0x5f, 0x05, 0x91, 0xff, 0x3f, 0x7b, 0x08, 0x90, 0xd1, 0xec, 0xed, 0xb8, 0x2e,
	0xd0, 0xf3, 0x95, 0x70, 0xee, 0xd6, 0xcf, 0x25, 0xb8, 0xf6, 0xa2, 0x91, 0x34, 0xff, 0xca, 0xa0,
	0x1d, 0x06, 0xe2, 0x77, 0x92, 0x7f, 0xc6, 0x9d, 0xd4, 0xb3, 0xe8, 0xdf, 0x3d, 0x7f, 0x53, 0xe1,
	0x3b, 0x7e, 0x7c, 0xf2, 0x2b, 0x6a, 0x75, 0xd9, 0xc9, 0xef, 0x3f, 0xdf, 0x5b, 0xeb, 0xf1, 0x3d,
	0x31, 0x6e, 0xfe, 0x13, 0xce, 0x1e, 0x8a, 0xb2, 0x51, 0x1c, 0xf9, 0x46, 0xf1, 0x8e, 0x50, 0x3f,
	0x63, 0x0f, 0xf8, 0x17, 0xe8, 0xec, 0xe0, 0x14, 0x9d, 0xe5, 0x83, 0x33, 0xa3, 0xee, 0x63, 0x3e,
	0x99, 0xdb, 0x18, 0x2e, 0xa0, 0xa5, 0xfe, 0x93, 0x9a, 0xfe, 0xc3, 0xf8, 0xfe, 0x05, 0xfa, 0xe8,
	0x6c, 0x65, 0xe0, 0x1b, 0x83, 0x5e, 0x95, 0xc3, 0x05, 0x26, 0x46, 0x0b, 0xee, 0xe2, 0x4c, 0x8e,
	0xf1, 0xd3, 0x9a, 0x9f, 0x0d, 0xbe, 0x7e, 0x9e, 0x1f, 0x1f, 0xd6, 0x46, 0x88, 0x2f, 0x78, 0xda,
	0x7e, 0x0e, 0xc3, 0x0c, 0x8f, 0x6b, 0x4a, 0x26, 0x6f, 0xa8, 0x6d, 0xdf, 0x68, 0x06, 0xba, 0xeb,
	0xdf, 0xd0, 0x2e, 0x7b, 0x73, 0xb9, 0xe4, 0x07, 0x9d, 0xf0, 0xaa, 0x1e, 0xfd, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xb3, 0xe5, 0x79, 0x41, 0xe3, 0x03, 0x00, 0x00,
}
