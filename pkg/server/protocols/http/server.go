package http

import (
	"context"
	"encoding/json"
	"fmt"
	"my-store/internal/errcode"

	// "my-store/internal/log"
	"my-store/pkg/models/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func EchoServer() *echo.Echo {
	server := echo.New()

	server.HTTPErrorHandler = errorHandler
	server.Binder = &auditBinder{defaultBinder: new(echo.DefaultBinder)}

	server.Any("", HealthCheck)

	return server
}

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// APIResponse API response error code OK and Message of OK with data was attached
func APIResponse(c echo.Context, data interface{}) error {
	resp := &response.BaseResponse{
		Error:     0,
		ErrorCode: errcode.OK,
		Message:   errcode.OK.Error(),
		Data:      data,
	}
	return c.JSON(http.StatusOK, resp)
}

// APIResponseError response with internal error code
func APIResponseError(c echo.Context, err error) error {
	resp := &response.BaseResponse{
		Data: nil,
	}
	// thay thế lỗi lib echo bằng lỗi app
	if he, ok := err.(*echo.HTTPError); ok {
		if he.Code == http.StatusNotFound {
			err = errcode.Error(errcode.MethodNotFound)
		}
	}
	resp.ErrorCode = errcode.GetAppErrorCode(err)
	resp.Data = errcode.GetAppErrorData(err)
	resp.Message = errcode.GetAppErrorMessage(err)
	code, _ := strconv.ParseInt(string(resp.ErrorCode), 10, 64)
	resp.Error = int(code)
	return c.JSON(http.StatusOK, resp)
}

type auditBinder struct {
	defaultBinder echo.Binder
}

// Bind Implement echo#Binder
func (binder *auditBinder) Bind(i interface{}, ctx echo.Context) error {
	if jsonBin := ctx.Request().Context().Value(auditBinder{}); jsonBin != nil {
		return json.Unmarshal(jsonBin.([]byte), i)
	}

	if err := binder.defaultBinder.Bind(i, ctx); err != nil {
		return fmt.Errorf("bind: %w", err)
	}

	jsonBin, err := json.Marshal(i)
	if err != nil {
		return err
	}
	ctx.SetRequest(ctx.Request().WithContext(context.WithValue(ctx.Request().Context(), auditBinder{}, jsonBin)))
	return nil
}

func errorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}

		if _, ok := he.Message.(string); ok {
			// log.Logger(c.Request().Context()).Sugar().Errorf("got echo internal code %d error %s", he.Code, msg)
		}
		err = he
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = APIResponseError(c, err)
		}
	}
}

type BaseMessage struct {
	ErrorCode string      `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Debug     interface{} `json:"debug,omitempty"`
}
