package class

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"

	"edumaster/internal/logic/class"
	"edumaster/internal/svc"
	"edumaster/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClassAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if err := svcCtx.Validate.Struct(req); err != nil {
			if err != nil {
				var validationErrors []types.ValidationError
				for _, err := range err.(validator.ValidationErrors) {
					validationErrors = append(validationErrors, types.ValidationError{
						Field:   err.Field(),
						Message: err.Translate(svcCtx.Trans),
					})
				}
				err = fmt.Errorf("验证错误: %v", validationErrors)
			}
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
