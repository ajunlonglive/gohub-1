package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//现在还是一个测试文件
func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}

}