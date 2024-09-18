package customerrelationshipgroup

import (
	"fmt"
	"net/http"
	"strings"

	"edumaster/internal/logic/customerrelationshipgroup"
	"edumaster/internal/svc"
	"edumaster/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/go-playground/validator/v10"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CustomerRelationshipGroupInfoReq
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

		l := customerrelationshipgroup.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
