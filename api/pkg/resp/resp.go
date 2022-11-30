package resp

import (
	"net/http"
	"order/api/pkg/errorx"
)

// Response 响应结构体
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// List 结构体
type List struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Size  int   `json:"size"`
	List  any   `json:"list"`
}

// Success 返回格式
func Success(obj any) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    obj,
	}
}

// SuccessList 成功列表返回格式
func SuccessList(total int64, page, size int, list any) Response {
	return Success(List{
		Total: total,
		Page:  page,
		Size:  size,
		List:  list,
	})
}

// Notice 返回格式
func Notice(code int, msg string) Response {
	response := Response{
		Code:    code,
		Message: msg,
	}
	return response
}

func Error(err *errorx.Error) Response {
	response := Response{
		Code:    err.Code,
		Message: err.String(),
	}

	return response
}
