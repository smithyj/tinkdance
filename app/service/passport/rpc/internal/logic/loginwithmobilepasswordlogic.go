package logic

import (
	"context"

	"tinkdance/service/passport/rpc/internal/svc"
	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithMobilePasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithMobilePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithMobilePasswordLogic {
	return &LoginWithMobilePasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithMobilePasswordLogic) LoginWithMobilePassword(in *pb.LoginWithMobilePasswordRequest) (*pb.LoginWithMobilePasswordResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginWithMobilePasswordResponse{}, nil
}
