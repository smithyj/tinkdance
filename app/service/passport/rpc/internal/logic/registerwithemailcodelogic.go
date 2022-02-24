package logic

import (
	"context"

	"tinkdance/service/passport/rpc/internal/svc"
	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterWithEmailCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterWithEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterWithEmailCodeLogic {
	return &RegisterWithEmailCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterWithEmailCodeLogic) RegisterWithEmailCode(in *pb.RegisterWithEmailCodeRequest) (*pb.RegisterWithEmailCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.RegisterWithEmailCodeResponse{}, nil
}
