package main

import (
	"github.com/gin-gonic/gin"
	"gohub/bootstrap"
)

func main() {
	r := gin.New()
	bootstrap.SetupRouter(r)
	// 运行服务
	r.Run()
}
