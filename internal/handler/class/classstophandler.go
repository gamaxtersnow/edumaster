package class

import (
	"net/http"

	"edumaster/internal/logic/class"
	"edumaster/internal/svc"
	"edumaster/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClassStopHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClassStopReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewClassStopLogic(r.Context(), svcCtx)
		resp, err := l.ClassStop(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
