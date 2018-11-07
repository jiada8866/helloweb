package errset

import (
	"fmt"
)

type (
	Response struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data,omitempty"`
	}

	ResponseError struct {
		Response
		// Error Types are empty when encoded to JSON
		// https://stackoverflow.com/questions/44989924/golang-error-types-are-empty-when-encoded-to-json
		Internal error  `json:"-"`
		ErrorMsg string `json:"error_msg,omitempty"`
	}
)

func (re *ResponseError) Error() string {
	return fmt.Sprintf("ResponseError: code=%d, msg=%s, internalErr: %v", re.Code, re.Msg, re.Internal)
}

// 一般如果 logic 层报 error 可以直接将该 error 在 handler 中返回，
// 由 HTTPErrorHandler 统一按某一个错误码(ErrServer)返回 response。
// 但有的时候需要将 logic 层的错误转换成不同的 ResponseError，
// 此时又想在 HTTPErrorHandler 统一处理真实的报错，
// 就需要用 WithInternal 将真实的错误包在 ResponseError 里。
func (re *ResponseError) WithInternal(err error) *ResponseError {
	return &ResponseError{
		Response: Response{
			Code: re.Code,
			Msg:  re.Msg,
		},
		Internal: err,
		ErrorMsg: err.Error(),
	}
}

func NewResponseError(code int, msg string) *ResponseError {
	return &ResponseError{Response: Response{Code: code, Msg: msg}}
}

func (r Response) WithData(data interface{}) Response {
	return Response{
		Code: r.Code,
		Msg:  r.Msg,
		Data: data,
	}
}
