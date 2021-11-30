package cron

import (
	"context"
	"fmt"
	"log"
	"smzdtz-server/internal/datacenter"
	"smzdtz-server/internal/datacenter/eastmoney"
	"smzdtz-server/internal/service"

	"github.com/gin-gonic/gin"
)

// SyncFund 同步基金数据
func SyncFund(c *gin.Context) (map[string]*service.Fund, error) {
	fmt.Printf("hello")
	// if !util.IsTradingDay() {
	// 	return
	// }
	ctx := context.Background()
	log.Printf("SyncFund request start...")

	// 获取全量列表
	efundlist, err := datacenter.EastMoney.QueryAllFundList(ctx, eastmoney.FundTypeALL)

	println(efundlist, err)
	fundCodes := []string{}
	for _, efund := range efundlist {
		fundCodes = append(fundCodes, efund.Fcode)
	}

	data, err := service.SearchFunds(ctx, fundCodes[0:2])

	return data, err
}
