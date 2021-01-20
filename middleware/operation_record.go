package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := time.Now()
		// TODO 日志收集
		c.Next()
		t := time.Since(s).Milliseconds()
		fmt.Println(t)
	}
}
