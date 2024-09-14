package classschedule

import (
	"net/http"

	"edumaster/internal/logic/classschedule"
	"edumaster/internal/svc"
	"edumaster/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClassScheduleModifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := classschedule.NewModifyLogic(r.Context(), svcCtx)
		resp, err := l.Modify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
