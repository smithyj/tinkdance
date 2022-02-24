// Code generated by goctl. DO NOT EDIT!
// Source: passport.proto

package passport

import (
	"context"

	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginWithAccountPasswordRequest     = pb.LoginWithAccountPasswordRequest
	LoginWithAccountPasswordResponse    = pb.LoginWithAccountPasswordResponse
	LoginWithEmailCodeRequest           = pb.LoginWithEmailCodeRequest
	LoginWithEmailCodeResponse          = pb.LoginWithEmailCodeResponse
	LoginWithEmailPasswordRequest       = pb.LoginWithEmailPasswordRequest
	LoginWithEmailPasswordResponse      = pb.LoginWithEmailPasswordResponse
	LoginWithMobileCodeRequest          = pb.LoginWithMobileCodeRequest
	LoginWithMobileCodeResponse         = pb.LoginWithMobileCodeResponse
	LoginWithMobilePasswordRequest      = pb.LoginWithMobilePasswordRequest
	LoginWithMobilePasswordResponse     = pb.LoginWithMobilePasswordResponse
	RegisterWithAccountPasswordRequest  = pb.RegisterWithAccountPasswordRequest
	RegisterWithAccountPasswordResponse = pb.RegisterWithAccountPasswordResponse
	RegisterWithEmailCodeRequest        = pb.RegisterWithEmailCodeRequest
	RegisterWithEmailCodeResponse       = pb.RegisterWithEmailCodeResponse
	RegisterWithMobileCodeRequest       = pb.RegisterWithMobileCodeRequest
	RegisterWithMobileCodeResponse      = pb.RegisterWithMobileCodeResponse

	Passport interface {
		LoginWithAccountPassword(ctx context.Context, in *LoginWithAccountPasswordRequest, opts ...grpc.CallOption) (*LoginWithAccountPasswordResponse, error)
		LoginWithEmailPassword(ctx context.Context, in *LoginWithEmailPasswordRequest, opts ...grpc.CallOption) (*LoginWithEmailPasswordResponse, error)
		LoginWithEmailCode(ctx context.Context, in *LoginWithEmailCodeRequest, opts ...grpc.CallOption) (*LoginWithEmailCodeResponse, error)
		LoginWithMobilePassword(ctx context.Context, in *LoginWithMobilePasswordRequest, opts ...grpc.CallOption) (*LoginWithMobilePasswordResponse, error)
		LoginWithMobileCode(ctx context.Context, in *LoginWithMobileCodeRequest, opts ...grpc.CallOption) (*LoginWithMobileCodeResponse, error)
		RegisterWithAccountPassword(ctx context.Context, in *RegisterWithAccountPasswordRequest, opts ...grpc.CallOption) (*RegisterWithAccountPasswordResponse, error)
		RegisterWithEmailCode(ctx context.Context, in *RegisterWithEmailCodeRequest, opts ...grpc.CallOption) (*RegisterWithEmailCodeResponse, error)
		RegisterWithMobileCode(ctx context.Context, in *RegisterWithMobileCodeRequest, opts ...grpc.CallOption) (*RegisterWithMobileCodeResponse, error)
	}

	defaultPassport struct {
		cli zrpc.Client
	}
)

func NewPassport(cli zrpc.Client) Passport {
	return &defaultPassport{
		cli: cli,
	}
}

func (m *defaultPassport) LoginWithAccountPassword(ctx context.Context, in *LoginWithAccountPasswordRequest, opts ...grpc.CallOption) (*LoginWithAccountPasswordResponse, error) {
	client := pb.NewPassportClient(m.cli.Conn())
	return client.LoginWithAccountPassword(ctx, in, opts...)
}

func (m *defaultPassport) LoginWithEmailPassword(ctx context.Context, in *LoginWithEmailPasswordRequest, opts ...grpc.CallOption) (*LoginWithEmailPasswordResponse, error) {
	client := pb.NewPassportClient(m.cli.Conn())
	return client.LoginWithEmailPassword(ctx, in, opts...)
}

func (m *defaultPassport) LoginWithEmailCode(ctx context.Context, in *LoginWithEmailCodeRequest, opts ...grpc.CallOption) (*LoginWithEmailCodeResponse, error) {
	client := pb.NewPassportClient(m.cli.Conn())
	return client.LoginWithEmailCode(ctx, in, opts...)
}

func (m *defaultPassport) LoginWithMobilePassword(ctx context.Context, in *LoginWithMobilePasswordRequest, opts ...grpc.CallOption) (*LoginWithMobilePasswordResponse, error) {
	client := pb.NewPassportClient(m.cli.Conn())
	return client.LoginWithMobilePassword(ctx, in, opts...)
}

func (m *defaultPassport) LoginWithMobileCode(ctx context.Context, in *LoginWithMobileCodeRequest, opts ...grpc.CallOption) (*LoginWithMobileCodeResponse, error) {
	client := pb.NewPassportClient(m.cli.Conn())
	return client.LoginWithMobileCode(ctx, in, opts...)
}

func (m *defaultPassport) RegisterWithAccountPassword(ctx context.Context, in *RegisterWithAccountPasswordRequest, opts ...grpc.CallOption) (*RegisterWithAccountPasswordResponse, error) {
	client := pb.NewPassportClient(m.cli.Conn())
	return client.RegisterWithAccountPassword(ctx, in, opts...)
}

func (m *defaultPassport) RegisterWithEmailCode(ctx context.Context, in *RegisterWithEmailCodeRequest, opts ...grpc.CallOption) (*RegisterWithEmailCodeResponse, error) {
	client := pb.NewPassportClient(m.cli.Conn())
	return client.RegisterWithEmailCode(ctx, in, opts...)
}

func (m *defaultPassport) RegisterWithMobileCode(ctx context.Context, in *RegisterWithMobileCodeRequest, opts ...grpc.CallOption) (*RegisterWithMobileCodeResponse, error) {
	client := pb.NewPassportClient(m.cli.Conn())
	return client.RegisterWithMobileCode(ctx, in, opts...)
}
