package response

import "my-store/internal/errcode"

type BaseResponse struct {
	Error     int               `json:"error"`
	ErrorCode errcode.ErrorCode `json:"errorCode"`
	Message   string            `json:"message"`
	Data      interface{}       `json:"data"`
}
