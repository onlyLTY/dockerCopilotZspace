package container

import (
	"net/http"

	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/logic/container"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ContainersListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := container.NewContainersListLogic(r.Context(), svcCtx)
		resp, err := l.ContainersList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
