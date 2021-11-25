// Package cron 定时任务
package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	promSyncLabels = []string{
		"jobname",
	}
	promSyncError = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "cron",
			Name:      "sync_error",
			Help:      "cron sync job error",
		}, promSyncLabels,
	)
)

// RunCronJobs 启动定时任务
func RunCronJobs(async bool) {
	timezone, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	sched := gocron.NewScheduler(timezone)
	// 测试定时任务
	sched.Every(5).Seconds().Do(SayHello)
	// sched.Cron("*/5 * * * * ?").Do(SayHello)
	// 同步基金净值列表和4433列表
	// sched.Cron("0 18 * * 1-5").Do(SyncFund)
	// 同步东方财富行业列表
	if async {
		sched.StartAsync()
	} else {
		sched.StartBlocking()
	}
}
