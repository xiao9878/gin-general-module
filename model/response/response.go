package response

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao9878/gin-general-module/utils"
	"net/http"
)

type Response struct {
	Code ResCode     `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const CodeSuccess ResCode = 0

func Result(code ResCode, data interface{}, c *gin.Context, msg ...string) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		utils.String(utils.IF(len(msg) > 0, msg[0], code.Msg())),
	})
}

// 成功的返回
func Ok(c *gin.Context) {
	Result(CodeSuccess, nil, c)
}

// 带消息的返回
func OkWithMessage(c *gin.Context, message string) {
	Result(CodeSuccess, nil, c, message)
}

// 带数据的返回
func OkWithData(data interface{}, c *gin.Context) {
	Result(CodeSuccess, data, c)
}

// 返回失败
func Fail(code ResCode, c *gin.Context) {
	Result(code, nil, c)
}

// 返回错误信息
func FailWithMessage(code ResCode, c *gin.Context, message string) {
	Result(code, nil, c, message)
}
