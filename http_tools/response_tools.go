package http_tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS = 0
	FAILED  = 1
)

func MakeResponse(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	c.JSON(httpCode, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"body": data,
	})
}

func JsonHttpSuccess(c *gin.Context, data interface{}) {
	MakeResponse(c, http.StatusOK, SUCCESS, "", data)
}

func JsonHttpError(c *gin.Context, msg string, data interface{}) {
	MakeResponse(c, http.StatusOK, FAILED, msg, data)
}
