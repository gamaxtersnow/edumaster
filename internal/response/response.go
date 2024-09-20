package response

import (
	"edumaster/internal/consts/errorcode"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ParseErrorResponse(w http.ResponseWriter, err error) {
	var body Body
	body.Code = errorcode.ParamParseError
	body.Msg = "不支持的参数传递格式"
	body.Data = nil
	httpx.OkJson(w, body)
}
func ValidateErrorResponse(w http.ResponseWriter, resp interface{}) {
	var body Body
	body.Code = errorcode.ParamValidateError
	body.Msg = "参数解析错误"
	body.Data = resp
	httpx.OkJson(w, body)
}
