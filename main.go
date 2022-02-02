package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// new 一个 Gin Engine 实例
	r := gin.New()

	// 注册中间件
	r.Use(gin.Logger(), gin.Recovery())

	r.NoRoute(func(c *gin.Context) {
		acceptString :=  c.Request.Header.Get("Accept")
		if strings.Contains(acceptString,"text/html") {
			c.String(http.StatusNotFound,"404您所访问的资源不存在")
		}else {
			c.JSON(http.StatusNotFound,gin.H{
				"error_code" : 404,
				"error_message" : "没有所指定的路由",
			})
		}

	})

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// 运行服务，默认为 8080，我们指定端口为 8000
	r.Run(":8000")
}