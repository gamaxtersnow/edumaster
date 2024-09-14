package customer

import (
	"net/http"

	"edumaster/internal/logic/customer"
	"edumaster/internal/svc"
	"edumaster/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CustomerAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := customer.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
