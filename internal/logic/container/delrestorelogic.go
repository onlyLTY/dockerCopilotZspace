package container

import (
	"context"

	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelRestoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelRestoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelRestoreLogic {
	return &DelRestoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelRestoreLogic) DelRestore(req *types.ContainerRestoreReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	return
}
