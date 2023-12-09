package response

import (
	"gin-blog/common/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseData 定义响应返回规范
type ResponseData struct {
	Code    config.ResCode `json:"code"`
	Message interface{}    `json:"msg"`
	Data    interface{}    `json:"data,omitempty"`
}

// ResponseError定义返回错误信息
func ResponseError(c *gin.Context, code config.ResCode) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	c.JSON(http.StatusOK, rd)
}

// ResponseSuccess 返回成功响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    config.SUCCESS,
		Message: config.SUCCESS.Msg(),
		Data:    data,
	}
	c.JSON(http.StatusOK, rd)
}
