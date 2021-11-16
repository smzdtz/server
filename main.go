package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	cron "smzdtz-server/corn"
	"smzdtz-server/routes"
	"smzdtz-server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 结束时关闭db连接
	defer utils.CloseGormInstances()

	// 判断是否加载viper配置
	if !utils.IsInitedViper() {
		panic("Running server must init viper by config file first!")
	}

	// 执行定时任务
	cron.RunCronJobs(true)

	var router = gin.Default()
	// 注册路由
	v1 := router.Group("/v1")
	// 测试
	routes.AddPingRoutes(v1)
	// 股票
	routes.AddStockRoutes(v1)
	// 基金
	routes.AddFundRoutes(v1)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Shutdown 时关闭 db 和 redis 连接
	srv.RegisterOnShutdown(func() {
		utils.CloseGormInstances()
		utils.CloseRedisInstances()
	})

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	router.Run(":5000")
}
