package container

import (
	"net/http"

	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/logic/container"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DelRestoreHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ContainerRestoreReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := container.NewDelRestoreLogic(r.Context(), svcCtx)
		resp, err := l.DelRestore(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
