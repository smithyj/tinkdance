// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: passport.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PassportClient is the client API for Passport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PassportClient interface {
	LoginWithAccountPassword(ctx context.Context, in *LoginWithAccountPasswordRequest, opts ...grpc.CallOption) (*LoginWithAccountPasswordResponse, error)
	LoginWithEmailPassword(ctx context.Context, in *LoginWithEmailPasswordRequest, opts ...grpc.CallOption) (*LoginWithEmailPasswordResponse, error)
	LoginWithEmailCode(ctx context.Context, in *LoginWithEmailCodeRequest, opts ...grpc.CallOption) (*LoginWithEmailCodeResponse, error)
	LoginWithMobilePassword(ctx context.Context, in *LoginWithMobilePasswordRequest, opts ...grpc.CallOption) (*LoginWithMobilePasswordResponse, error)
	LoginWithMobileCode(ctx context.Context, in *LoginWithMobileCodeRequest, opts ...grpc.CallOption) (*LoginWithMobileCodeResponse, error)
	RegisterWithAccountPassword(ctx context.Context, in *RegisterWithAccountPasswordRequest, opts ...grpc.CallOption) (*RegisterWithAccountPasswordResponse, error)
	RegisterWithEmailCode(ctx context.Context, in *RegisterWithEmailCodeRequest, opts ...grpc.CallOption) (*RegisterWithEmailCodeResponse, error)
	RegisterWithMobileCode(ctx context.Context, in *RegisterWithMobileCodeRequest, opts ...grpc.CallOption) (*RegisterWithMobileCodeResponse, error)
}

type passportClient struct {
	cc grpc.ClientConnInterface
}

func NewPassportClient(cc grpc.ClientConnInterface) PassportClient {
	return &passportClient{cc}
}

func (c *passportClient) LoginWithAccountPassword(ctx context.Context, in *LoginWithAccountPasswordRequest, opts ...grpc.CallOption) (*LoginWithAccountPasswordResponse, error) {
	out := new(LoginWithAccountPasswordResponse)
	err := c.cc.Invoke(ctx, "/pb.Passport/LoginWithAccountPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportClient) LoginWithEmailPassword(ctx context.Context, in *LoginWithEmailPasswordRequest, opts ...grpc.CallOption) (*LoginWithEmailPasswordResponse, error) {
	out := new(LoginWithEmailPasswordResponse)
	err := c.cc.Invoke(ctx, "/pb.Passport/LoginWithEmailPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportClient) LoginWithEmailCode(ctx context.Context, in *LoginWithEmailCodeRequest, opts ...grpc.CallOption) (*LoginWithEmailCodeResponse, error) {
	out := new(LoginWithEmailCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.Passport/LoginWithEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportClient) LoginWithMobilePassword(ctx context.Context, in *LoginWithMobilePasswordRequest, opts ...grpc.CallOption) (*LoginWithMobilePasswordResponse, error) {
	out := new(LoginWithMobilePasswordResponse)
	err := c.cc.Invoke(ctx, "/pb.Passport/LoginWithMobilePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportClient) LoginWithMobileCode(ctx context.Context, in *LoginWithMobileCodeRequest, opts ...grpc.CallOption) (*LoginWithMobileCodeResponse, error) {
	out := new(LoginWithMobileCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.Passport/LoginWithMobileCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportClient) RegisterWithAccountPassword(ctx context.Context, in *RegisterWithAccountPasswordRequest, opts ...grpc.CallOption) (*RegisterWithAccountPasswordResponse, error) {
	out := new(RegisterWithAccountPasswordResponse)
	err := c.cc.Invoke(ctx, "/pb.Passport/RegisterWithAccountPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportClient) RegisterWithEmailCode(ctx context.Context, in *RegisterWithEmailCodeRequest, opts ...grpc.CallOption) (*RegisterWithEmailCodeResponse, error) {
	out := new(RegisterWithEmailCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.Passport/RegisterWithEmailCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *passportClient) RegisterWithMobileCode(ctx context.Context, in *RegisterWithMobileCodeRequest, opts ...grpc.CallOption) (*RegisterWithMobileCodeResponse, error) {
	out := new(RegisterWithMobileCodeResponse)
	err := c.cc.Invoke(ctx, "/pb.Passport/RegisterWithMobileCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PassportServer is the server API for Passport service.
// All implementations must embed UnimplementedPassportServer
// for forward compatibility
type PassportServer interface {
	LoginWithAccountPassword(context.Context, *LoginWithAccountPasswordRequest) (*LoginWithAccountPasswordResponse, error)
	LoginWithEmailPassword(context.Context, *LoginWithEmailPasswordRequest) (*LoginWithEmailPasswordResponse, error)
	LoginWithEmailCode(context.Context, *LoginWithEmailCodeRequest) (*LoginWithEmailCodeResponse, error)
	LoginWithMobilePassword(context.Context, *LoginWithMobilePasswordRequest) (*LoginWithMobilePasswordResponse, error)
	LoginWithMobileCode(context.Context, *LoginWithMobileCodeRequest) (*LoginWithMobileCodeResponse, error)
	RegisterWithAccountPassword(context.Context, *RegisterWithAccountPasswordRequest) (*RegisterWithAccountPasswordResponse, error)
	RegisterWithEmailCode(context.Context, *RegisterWithEmailCodeRequest) (*RegisterWithEmailCodeResponse, error)
	RegisterWithMobileCode(context.Context, *RegisterWithMobileCodeRequest) (*RegisterWithMobileCodeResponse, error)
	mustEmbedUnimplementedPassportServer()
}

// UnimplementedPassportServer must be embedded to have forward compatible implementations.
type UnimplementedPassportServer struct {
}

func (UnimplementedPassportServer) LoginWithAccountPassword(context.Context, *LoginWithAccountPasswordRequest) (*LoginWithAccountPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithAccountPassword not implemented")
}
func (UnimplementedPassportServer) LoginWithEmailPassword(context.Context, *LoginWithEmailPasswordRequest) (*LoginWithEmailPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithEmailPassword not implemented")
}
func (UnimplementedPassportServer) LoginWithEmailCode(context.Context, *LoginWithEmailCodeRequest) (*LoginWithEmailCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithEmailCode not implemented")
}
func (UnimplementedPassportServer) LoginWithMobilePassword(context.Context, *LoginWithMobilePasswordRequest) (*LoginWithMobilePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithMobilePassword not implemented")
}
func (UnimplementedPassportServer) LoginWithMobileCode(context.Context, *LoginWithMobileCodeRequest) (*LoginWithMobileCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithMobileCode not implemented")
}
func (UnimplementedPassportServer) RegisterWithAccountPassword(context.Context, *RegisterWithAccountPasswordRequest) (*RegisterWithAccountPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWithAccountPassword not implemented")
}
func (UnimplementedPassportServer) RegisterWithEmailCode(context.Context, *RegisterWithEmailCodeRequest) (*RegisterWithEmailCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWithEmailCode not implemented")
}
func (UnimplementedPassportServer) RegisterWithMobileCode(context.Context, *RegisterWithMobileCodeRequest) (*RegisterWithMobileCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterWithMobileCode not implemented")
}
func (UnimplementedPassportServer) mustEmbedUnimplementedPassportServer() {}

// UnsafePassportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PassportServer will
// result in compilation errors.
type UnsafePassportServer interface {
	mustEmbedUnimplementedPassportServer()
}

func RegisterPassportServer(s grpc.ServiceRegistrar, srv PassportServer) {
	s.RegisterService(&Passport_ServiceDesc, srv)
}

func _Passport_LoginWithAccountPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithAccountPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServer).LoginWithAccountPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passport/LoginWithAccountPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServer).LoginWithAccountPassword(ctx, req.(*LoginWithAccountPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passport_LoginWithEmailPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithEmailPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServer).LoginWithEmailPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passport/LoginWithEmailPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServer).LoginWithEmailPassword(ctx, req.(*LoginWithEmailPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passport_LoginWithEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServer).LoginWithEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passport/LoginWithEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServer).LoginWithEmailCode(ctx, req.(*LoginWithEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passport_LoginWithMobilePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithMobilePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServer).LoginWithMobilePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passport/LoginWithMobilePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServer).LoginWithMobilePassword(ctx, req.(*LoginWithMobilePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passport_LoginWithMobileCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithMobileCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServer).LoginWithMobileCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passport/LoginWithMobileCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServer).LoginWithMobileCode(ctx, req.(*LoginWithMobileCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passport_RegisterWithAccountPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWithAccountPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServer).RegisterWithAccountPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passport/RegisterWithAccountPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServer).RegisterWithAccountPassword(ctx, req.(*RegisterWithAccountPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passport_RegisterWithEmailCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWithEmailCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServer).RegisterWithEmailCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passport/RegisterWithEmailCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServer).RegisterWithEmailCode(ctx, req.(*RegisterWithEmailCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Passport_RegisterWithMobileCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterWithMobileCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PassportServer).RegisterWithMobileCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Passport/RegisterWithMobileCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PassportServer).RegisterWithMobileCode(ctx, req.(*RegisterWithMobileCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Passport_ServiceDesc is the grpc.ServiceDesc for Passport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Passport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Passport",
	HandlerType: (*PassportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginWithAccountPassword",
			Handler:    _Passport_LoginWithAccountPassword_Handler,
		},
		{
			MethodName: "LoginWithEmailPassword",
			Handler:    _Passport_LoginWithEmailPassword_Handler,
		},
		{
			MethodName: "LoginWithEmailCode",
			Handler:    _Passport_LoginWithEmailCode_Handler,
		},
		{
			MethodName: "LoginWithMobilePassword",
			Handler:    _Passport_LoginWithMobilePassword_Handler,
		},
		{
			MethodName: "LoginWithMobileCode",
			Handler:    _Passport_LoginWithMobileCode_Handler,
		},
		{
			MethodName: "RegisterWithAccountPassword",
			Handler:    _Passport_RegisterWithAccountPassword_Handler,
		},
		{
			MethodName: "RegisterWithEmailCode",
			Handler:    _Passport_RegisterWithEmailCode_Handler,
		},
		{
			MethodName: "RegisterWithMobileCode",
			Handler:    _Passport_RegisterWithMobileCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "passport.proto",
}
