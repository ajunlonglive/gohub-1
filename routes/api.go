package routes

import "github.com/gin-gonic/gin"

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(200,"hello")

		})
	}

}
