// 获取历史股价

package eniu

import (
	"context"
	"errors"
	"fmt"
	"math"
	"smzdtz-server/pkg/http"
	"smzdtz-server/pkg/statistics"
	"strings"
	"time"
)

// RespHistoricalStockPrice 历史股价接口返回结构
type RespHistoricalStockPrice struct {
	Date  []string  `json:"date"`
	Price []float64 `json:"price"`
}

// LastYearFinalPrice 获取去年12月份最后一个交易日的股价
func (p RespHistoricalStockPrice) LastYearFinalPrice() float64 {
	if len(p.Date) == 0 {
		return 0
	}
	for i := len(p.Date) - 1; i > 0; i-- {
		prefix := fmt.Sprintf("%d-12-", time.Now().Year()-1)
		date := p.Date[i]
		if strings.Contains(date, prefix) {
			price := p.Price[i]
			return price
		}
	}
	return 0
}

// HistoricalVolatility 计算历史波动率
// 历史波动率计算方法：https://goodcalculators.com/historical-volatility-calculator/
// 1、从市场上获得标的股票在固定时间间隔(如每天DAY、每周WEEK或每月MONTH等)上的价格。
// 2、对于每个时间段，求出该时间段末的股价与该时段初的股价之比的自然对数。
// 3、求出这些对数值的标准差即为历史波动率的估计值
// 4、若将日、周等标准差转化为年标准差，需要再乘以一年中包含的时段数量的平方根(如，选取时间间隔为每天，则若扣除闭市，每年中有250个交易日，应乘以根号250)
func (p RespHistoricalStockPrice) HistoricalVolatility(ctx context.Context, period string) (float64, error) {
	priceLen := len(p.Price)
	if priceLen == 0 {
		return -1.0, errors.New("no historical price data")
	}
	// 求末初股价比自然对数
	logs := []float64{}
	for i := priceLen - 1; i >= 1; i-- {
		endPrice := p.Price[i]
		startPrice := p.Price[i-1]
		log := math.Log(endPrice / startPrice)
		logs = append(logs, log)
	}
	// 标准差
	stdev, err := statistics.StdDeviationFloat64(logs)
	if err != nil {
		return -1.0, err
	}

	periodValue := float64(250)
	period = strings.ToUpper(period)
	switch period {
	case "DAY":
		periodValue = 1
	case "WEEK":
		periodValue = 5
	case "MONTH":
		periodValue = 21.75
	case "YEAR":
		periodValue = 250
	}
	volatility := stdev * math.Sqrt(periodValue)
	// 数据异常时全部股价为 0 导致返回 NaN
	if math.IsNaN(volatility) {
		return -1, errors.New("volatility is NaN")
	}
	return volatility, nil
}

// QueryHistoricalStockPrice 获取历史股价，最新数据在最后，有一天的延迟
func (e Eniu) QueryHistoricalStockPrice(ctx context.Context, secuCode string) (RespHistoricalStockPrice, error) {
	apiurl := fmt.Sprintf("https://eniu.com/chart/pricea/%s/t/all", e.GetPathCode(ctx, secuCode))
	resp := RespHistoricalStockPrice{}
	err := http.HTTPGET(ctx, e.HTTPClient, apiurl, nil, &resp)
	return resp, err
}

// GetPathCode 返回接口 url path 中的股票代码
func (e Eniu) GetPathCode(ctx context.Context, secuCode string) string {
	s := strings.Split(secuCode, ".")
	if len(s) != 2 {
		return ""
	}
	return strings.ToLower(s[1]) + s[0]
}
