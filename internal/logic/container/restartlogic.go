package container

import (
	"context"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/types"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/utiles"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestartLogic {
	return &RestartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestartLogic) Restart(req *types.IdReq) (resp *types.Resp, err error) {
	resp = &types.Resp{}
	err = utiles.RestartContainer(l.svcCtx, req.Id)
	if err != nil {
		resp.Code = 400
		resp.Msg = err.Error()
		resp.Data = map[string]interface{}{}
		return resp, err
	}
	resp.Code = 200
	resp.Msg = "success"
	resp.Data = map[string]interface{}{}
	return resp, nil
}
