package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nanfeng/ginchat/internal/pkg/xerr"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data any, msg ...string) {
	message := "success"

	if len(msg) >= 0 {
		message = msg[0]
	}

	c.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  message,
		Data: data,
	})
}

// Fail 通用响应失败
func Fail(c *gin.Context, code int, msg string, httpCode ...int) {
	status := http.StatusBadRequest
	if len(httpCode) > 0 {
		status = httpCode[0]
	}

	c.JSON(status, Response{
		Code: code,
		Msg:  msg,
	})
}

// BadRequest 参数错误快捷方法
func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Code: xerr.ErrInvalidParam.Code,
		Msg:  msg,
	})
}

// UnauthorizedRequest 未授权错误快捷方法
func UnauthorizedRequest(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: xerr.ErrUnanthorizad.Code,
		Msg:  xerr.ErrUnanthorizad.Error(),
	})
}
