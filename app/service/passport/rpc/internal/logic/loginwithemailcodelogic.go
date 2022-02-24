package logic

import (
	"context"

	"tinkdance/service/passport/rpc/internal/svc"
	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWithEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginWithEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWithEmailCodeLogic {
	return &LoginWithEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginWithEmailCodeLogic) LoginWithEmailCode(in *pb.LoginWithEmailCodeRequest) (*pb.LoginWithEmailCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.LoginWithEmailCodeResponse{}, nil
}
