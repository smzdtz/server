package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"smzdtz-server/internal/cron"
	"smzdtz-server/internal/router"

	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func initViper(configPath, configName, configType string, onConfigChangeRun func(fsnotify.Event)) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.SetDefault("viper.inited", true)
	viper.WatchConfig()
	if onConfigChangeRun != nil {
		viper.OnConfigChange(onConfigChangeRun)
	}
	return nil
}

func main() {
	// 结束时关闭db连接
	// defer utils.CloseGormInstances()
	// 加载配置文件内容到 viper 中以便使用
	if err := initViper(".", "config", "toml", func(e fsnotify.Event) {
		log.Println("Config file changed")
	}); err != nil {
		// 文件不存在时 1 使用默认配置，其他 err 直接 panic
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(err)
		}
		log.Println(nil, "Init viper error:"+err.Error())
	}

	// 执行定时任务
	cron.RunCronJobs(true)
	var router = router.InitRouter()

	srv := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	// Shutdown 时关闭 db 和 redis 连接
	srv.RegisterOnShutdown(func() {
		fmt.Println("shut down...")
		// utils.CloseGormInstances()
		// utils.CloseRedisInstances()
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
