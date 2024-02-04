package container

import (
	"net/http"

	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/logic/container"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RenameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ContainerRenameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := container.NewRenameLogic(r.Context(), svcCtx)
		resp, err := l.Rename(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
