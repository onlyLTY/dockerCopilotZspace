package image

import (
	"net/http"

	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/logic/image"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveImageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := image.NewRemoveLogic(r.Context(), svcCtx)
		resp, err := l.Remove(&req)
		if err != nil {
			httpx.WriteJson(w, resp.Code, resp)
		} else {
			httpx.WriteJson(w, resp.Code, resp)
		}
	}
}
