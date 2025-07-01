package main

import (
	"fmt"
	"log"
	"go-mall/config"
	"go-mall/router"
	"go-mall/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatal("加载配置失败:", err)
	}

	// 初始化数据库
	if err := utils.InitDB(); err != nil {
		log.Fatal("初始化数据库失败:", err)
	}

	// 设置Gin模式
	gin.SetMode(config.AppConfig.Server.Mode)

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	port := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	log.Printf("服务器启动在端口 %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
} 