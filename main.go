package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {

	// new 一个 Gin Engine 实例
	r := gin.New()

	bootstrap.SetupRoute(r)
	// 运行服务，默认为 8080，我们指定端口为 3000
	err := r.Run(":3001")
	if err != nil {
		fmt.Println("服务启动异常",err)
	}
}