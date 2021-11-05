package datacenter

import (
	"smzdtz-server/datacenter/eastmoney"
	"smzdtz-server/datacenter/sina"
)

var (
	// 东方财富
	EastMoney eastmoney.EastMoney
	// 新浪财经
	Sina sina.Sina
)

func init() {
	Sina = sina.NewSina()
	EastMoney = eastmoney.NewEastMoney()
}
