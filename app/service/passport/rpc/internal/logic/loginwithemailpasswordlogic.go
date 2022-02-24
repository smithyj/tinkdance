package logic

import (
	"context"

	"tinkdance/service/passport/rpc/internal/svc"
	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithEmailPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithEmailPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithEmailPasswordLogic {
	return &LoginWithEmailPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithEmailPasswordLogic) LoginWithEmailPassword(in *pb.LoginWithEmailPasswordRequest) (*pb.LoginWithEmailPasswordResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginWithEmailPasswordResponse{}, nil
}
