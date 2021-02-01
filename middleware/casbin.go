package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiao9878/gin-general-module/global"
	"github.com/xiao9878/gin-general-module/model/request"
	"github.com/xiao9878/gin-general-module/model/response"
	service "github.com/xiao9878/gin-general-module/sys_service"
)

// 拦截器
func CasBin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := waitUse.AuthorityId
		e := service.Casbin()
		// 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if global.GVA_CONFIG.System.Env == "dev" || success {
			c.Next()
		} else {
			response.Fail(response.AuthError, c)
			c.Abort()
			return
		}
	}
}
