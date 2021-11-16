package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"smzdtz-server/cron"
	"smzdtz-server/routes"
	"smzdtz-server/utils"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 结束时关闭db连接
	defer utils.CloseGormInstances()

	// 加载配置文件内容到 viper 中以便使用
	if err := utils.InitViper(".", "config", "toml",
		func(e fsnotify.Event) {
			fmt.Println("Config file changed:" + e.Name)
		}); err != nil {
		// 文件不存在时 1 使用默认配置，其他 err 直接 panic
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(err)
		}
		fmt.Println("Init viper error:" + err.Error())
	}

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
		Addr:    ":5000",
		Handler: router,
	}

	// Shutdown 时关闭 db 和 redis 连接
	srv.RegisterOnShutdown(func() {
		fmt.Println("shut down...")
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
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
