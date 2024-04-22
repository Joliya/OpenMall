package middleware

import (
	"OpenMall/logger"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]interface{}
		if c.Request.Method == "POST" {
			body, _ := io.ReadAll(c.Request.Body)
			// 等于拷贝一份往下传递,否则下接口的方法中拿不到请求体数据
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			_ = json.Unmarshal(body, &data)
		}
		start := time.Now()
		// 处理请求
		c.Next()
		// 计算请求处理时间
		latency := time.Since(start)
		var params interface{}
		if c.Request.Method == "POST" {
			// 判断请求头是否是json格式
			contentType := c.Request.Header.Get("Content-Type")
			if contentType == "application/json" {
				params = data
			} else {
				params = c.Request.PostForm
			}
		} else {
			params = c.Request.URL.Query()
		}
		logger.RouterLogger.Info("request completion",
			zap.Any("ip", c.ClientIP()),
			zap.Any("code", c.Writer.Status()),
			zap.Any("method", c.Request.Method),
			zap.Any("path", c.Request.URL.Path),
			zap.Any("params", params),
			zap.Any("proto", c.Request.Proto),
			zap.Any("cost", latency),
		)
	}
}
