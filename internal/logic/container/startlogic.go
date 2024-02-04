package container

import (
	"context"

	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartLogic {
	return &StartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartLogic) Start(req *types.IdReq) (resp *types.Resp, err error) {
	// todo: add your logic here and delete this line

	return
}
