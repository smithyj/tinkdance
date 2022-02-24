package logic

import (
	"context"

	"tinkdance/service/passport/rpc/internal/svc"
	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithAccountPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithAccountPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithAccountPasswordLogic {
	return &LoginWithAccountPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithAccountPasswordLogic) LoginWithAccountPassword(in *pb.LoginWithAccountPasswordRequest) (*pb.LoginWithAccountPasswordResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginWithAccountPasswordResponse{}, nil
}
