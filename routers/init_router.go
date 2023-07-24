package routers

import (
	"OpenMall/routers/sku"
	"github.com/gin-gonic/gin"
)

// InitRouter 将每个模块的路由路径进行拆分，在这儿统一初始化，gin框架的路由启动
func InitRouter(server *gin.Engine) {
	sku.InitSkuRouters(server)
}
