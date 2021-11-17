package datacenter

import (
	"smzdtz-server/datacenter/cninfo"
	"smzdtz-server/datacenter/eastmoney"
	"smzdtz-server/datacenter/sina"
	"smzdtz-server/datacenter/zsxg"
)

var (
	// 东方财富
	EastMoney eastmoney.EastMoney
	// 新浪财经
	Sina sina.Sina
	// 芝士财富
	Zsxg zsxg.Zsxg
	// 巨潮资讯
	CnInfo cninfo.CnInfo
)

func init() {
	Sina = sina.NewSina()
	EastMoney = eastmoney.NewEastMoney()
	Zsxg = zsxg.NewZsxg()
	CnInfo = cninfo.NewCninfo()
}
