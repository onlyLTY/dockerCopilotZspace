package container

import (
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/types"
	"net/http"

	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/logic/container"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
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
			httpx.WriteJson(w, resp.Code, resp)
		} else {
			httpx.WriteJson(w, resp.Code, resp)
		}
	}
}
