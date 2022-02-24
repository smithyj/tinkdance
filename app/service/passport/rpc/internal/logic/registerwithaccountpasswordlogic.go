package logic

import (
	"context"

	"tinkdance/service/passport/rpc/internal/svc"
	"tinkdance/service/passport/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterWithAccountPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterWithAccountPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterWithAccountPasswordLogic {
	return &RegisterWithAccountPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterWithAccountPasswordLogic) RegisterWithAccountPassword(in *pb.RegisterWithAccountPasswordRequest) (*pb.RegisterWithAccountPasswordResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.RegisterWithAccountPasswordResponse{}, nil
}
