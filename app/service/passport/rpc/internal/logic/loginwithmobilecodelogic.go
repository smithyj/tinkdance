package logic

import (
	"context"

	"tinkdance/service/passport/rpc/internal/svc"
	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithMobileCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithMobileCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithMobileCodeLogic {
	return &LoginWithMobileCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithMobileCodeLogic) LoginWithMobileCode(in *pb.LoginWithMobileCodeRequest) (*pb.LoginWithMobileCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginWithMobileCodeResponse{}, nil
}
