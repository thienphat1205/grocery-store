package errcode

import (
	"fmt"
	"reflect"
)

const errorPrefix = "#"
const errorSuffix = " - "

func Error(c ErrorCode, args ...interface{}) error {
	return &AppError{
		Code:    c,
		Message: fmt.Sprintf(fmt.Sprintf("#%s - ", string(c))+c.String(), args...),
	}
}

type AppError struct {
	Status  string      `json:"status"`
	Code    ErrorCode   `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (err *AppError) Error() string {
	return fmt.Sprintf("%s%s", codeFormat(err.Code), err.Message)
}

func IsAppError(objectPtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(&AppError{})
}

func codeFormat(code ErrorCode) string {
	return fmt.Sprintf("%s%s%s", errorPrefix, string(code), errorSuffix)
}

func GetAppErrorCode(err interface{}) ErrorCode {
	if IsAppError(err) {
		status := err.(*AppError).Code
		if status != "" {
			return status
		}
	}
	return Unknown
}

func GetAppErrorMessage(err interface{}) string {
	if IsAppError(err) {
		status := err.(*AppError).Message
		if status != "" {
			return status
		}
	}
	return Unknown.Error()
}

func GetAppErrorData(err interface{}) interface{} {
	if IsAppError(err) {
		return err.(*AppError).Data
	}
	return nil
}
