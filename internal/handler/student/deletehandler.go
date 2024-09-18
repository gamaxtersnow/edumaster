package student

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"

	"edumaster/internal/logic/student"
	"edumaster/internal/svc"
	"edumaster/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StudentDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if err := svcCtx.Validate.Struct(req); err != nil {
			var translatedErrors []string
			for _, e := range err.(validator.ValidationErrors) {
				ns := e.Namespace()
				if ns != "" {
					firstDotIndex := strings.Index(ns, ".")
					if firstDotIndex != -1 {
						ns = ns[firstDotIndex+1:]
					}
				}
				translatedErrors = append(translatedErrors, fmt.Sprintf("%s:%s", ns, strings.TrimLeft(e.Translate(svcCtx.Trans), e.Field())))
			}
			httpx.Error(w, fmt.Errorf("%v", translatedErrors))
			return
		}

		l := student.NewDeleteLogic(r.Context(), svcCtx)
		resp, err := l.Delete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
