package middleware

import (
	"gin-general-module/config"
	"gin-general-module/response"
	"github.com/gin-gonic/gin"
)

// 拦截器
func CasBin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.S.RunMod == "dev" {
			c.Next()
		}
		// TODO 鉴权
		response.Fail(response.AuthError, c)
		c.Abort()
		return
	}
}
