package logic

import (
	"context"

	"tinkdance/service/passport/rpc/internal/svc"
	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterWithMobileCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterWithMobileCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterWithMobileCodeLogic {
	return &RegisterWithMobileCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterWithMobileCodeLogic) RegisterWithMobileCode(in *pb.RegisterWithMobileCodeRequest) (*pb.RegisterWithMobileCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.RegisterWithMobileCodeResponse{}, nil
}
