// 获取选股器中的行业列表数据

package eastmoney

import (
	"context"
	"fmt"
	"smzdtz-server/utils"
)

// RespIndustryList 接口返回的 json 结构
type RespIndustryList struct {
	Result struct {
		Count int `json:"count"`
		Pages int `json:"pages"`
		Data  []struct {
			Industry    string `json:"INDUSTRY"`
			FirstLetter string `json:"FIRST_LETTER"`
		} `json:"data"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// QueryIndustryList 获取行业列表
func (e EastMoney) QueryIndustryList(ctx context.Context) ([]string, error) {
	apiurl := "https://datacenter.eastmoney.com/stock/selection/api/data/get/"
	reqData := map[string]string{
		"source": "SELECT_SECURITIES",
		"client": "APP",
		"type":   "RPTA_APP_INDUSTRY",
		"sty":    "ALL",
	}
	req, err := utils.NewHTTPMultipartReq(ctx, apiurl, reqData)
	if err != nil {
		return nil, err
	}
	resp := RespIndustryList{}
	err = utils.HTTPPOST(ctx, e.HTTPClient, req, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, fmt.Errorf("%#v", resp)
	}
	result := []string{}
	for _, i := range resp.Result.Data {
		result = append(result, i.Industry)
	}
	return result, nil
}
