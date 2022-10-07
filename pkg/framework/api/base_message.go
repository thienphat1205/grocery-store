package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseMessage struct {
	Code    int         `json:"code"`
	Error   int         `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseOK API response error code OK and Message of OK with data was attached
func ResponseOK(c echo.Context, data interface{}) error {
	resp := &BaseMessage{
		Code:    200,
		Error:   0,
		Message: "Thao tác thành công",
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}
