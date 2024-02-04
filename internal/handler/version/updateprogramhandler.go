package version

import (
	"net/http"

	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/logic/version"
	"github.com/onlyLTY/dockerCopilotZspace/zspace/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateProgramHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := version.NewUpdateProgramLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProgram()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
