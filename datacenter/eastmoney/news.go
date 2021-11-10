package eastmoney

import (
	"context"
	"smzdtz-server/utils"
)

type Article struct {
	ArtUniqueUrl  string `json:"Art_UniqueUrl"`
	ArtTitle      string `json:"Art_Title"`
	ArtUrl        string `json:"Art_Url"`
	ArtCreateTime string `json:"Art_CreateTime"`
	ArtContent    string `json:"Art_Content"`
}
type RespStockNews struct {
	IsSuccess   bool      `json:"IsSuccess"`
	Code        int       `json:"Code"`
	Message     string    `json:"Message"`
	TotalPage   int       `json:"TotalPage"`
	TotalCount  int       `json:"TotalCount"`
	Keyword     string    `json:"Keyword"`
	Data        []Article `json:"Data"`
	RelatedWord string    `json:"RelatedWord"`
	StillSearch []string  `json:"StillSearch"`
	StockModel  struct {
		Name string `json:"Name"`
		Code string `json:"Code"`
	} `json:"StockModel"`
}

func (e EastMoney) GetStockNews(ctx context.Context, secuCode string) ([]Article, error) {
	apiUrl := "http://searchapi.eastmoney.com//bussiness/Web/GetCMSSearchList"
	params := map[string]string{
		"type":      "8196",
		"pageindex": "1",
		"pagesize":  "20",
		"keyword":   secuCode,
		"name":      "zixun",
		"_":         "1608800267874",
	}
	header := map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
		"Accept-Charset":  "utf-8",
		"Accept":          "application/json",
		"Accept-Encoding": "utf-8",
		"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
		"Cache-Control":   "no-cache",
		"Connection":      "keep-alive",
		"Host":            "searchapi.eastmoney.com",
		"Pragma":          "no-cache",
		"Referer":         "http://so.eastmoney.com/",
	}
	newsList := []Article{}
	resp := RespStockNews{}
	apiurl, err := utils.NewHTTPGetURLWithQueryString(ctx, apiUrl, params)
	if err != nil {
		return newsList, err
	}
	err = utils.HTTPGET(ctx, e.HTTPClient, apiurl, header, &resp)
	newsList = append(newsList, resp.Data...)
	return newsList, err
}
