// A股股票列表
package cninfo

import (
	"context"
	"fmt"
	"smzdtz-server/pkg/http"
	// "smzdtz-server/utils"
)

type RespStocks struct {
	StockList []Stock `json:"stockList"`
}

type Result struct {
	Secucode string `json:"Secucode"`
	Name     string `json:"Name"`
	Market   string `json:"Market"`
	Pinyin   string `json:"Pinyin"`
}

type Stock struct {
	OrgID    string `json:"orgId"`
	Category string `json:"category"`
	Code     string `json:"code"`
	Pinyin   string `json:"pinyin"`
	Zwjc     string `json:"zwjc"`
}

func (c CnInfo) StockList(ctx context.Context) (stocks []Result, err error) {
	apiurl := "http://www.cninfo.com.cn/new/data/szse_stock.json"
	resp := RespStocks{}
	fmt.Println(apiurl)
	err = http.HTTPGET(ctx, c.HTTPClient, apiurl, nil, &resp)
	if err != nil {
		return nil, err
	}
	results := []Result{}
	for _, stock := range resp.StockList {
		result := Result{
			Secucode: stock.Code,
			Name:     stock.Zwjc,
			Pinyin:   stock.Pinyin,
			Market:   stock.Category,
		}
		results = append(results, result)
	}
	return results, nil
}
