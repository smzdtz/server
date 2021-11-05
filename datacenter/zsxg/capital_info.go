package zsxg

import (
	"context"
	"smzdtz-server/utils"
	"strings"
)

// RespZSCF 芝士财富接口返回结构
type RespZSCF struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Datas   struct {
		ListTime   int64 `json:"listTime"`
		CommentNew struct {
			PositiveNew []struct {
				Value string `json:"Value"`
				Tag   string `json:"Tag"`
			} `json:"positive_new"`
			UnPositiveNew []struct {
				Value string `json:"Value"`
				Tag   string `json:"Tag"`
			} `json:"unpositive_new"`
		} `json:"comment_new"`
		Close string `json:"close"`
	} `json:"datas"`
}

type CapitalInfo struct {
	// 股票上市时间
	ListTime int64 `json:"listTime"`
	// 收盘价
	Close string `json:"close"`
	//	芝士财富投资亮点
	PositiveNew []struct {
		Value string `json:"Value"`
		Tag   string `json:"Tag"`
	}
	// 芝士财富投资风险提示
	UnPositiveNew []struct {
		Value string `json:"Value"`
		Tag   string `json:"Tag"`
	}
}

func (e Zsxg) QueryCommentNew(ctx context.Context, secuCode string) (CapitalInfo, error) {
	capitalInfo := CapitalInfo{}

	// 芝士财富股票接口
	apiurl := "https://zsxg.cn/api/v2/capital/info?code=" + strings.ToUpper(secuCode) + "&yearNum=6"
	resp := RespZSCF{}
	err := utils.HTTPGET(ctx, e.HTTPClient, apiurl, nil, &resp)
	if err != nil {
		return capitalInfo, err
	}
	if resp.Code != "000" {
		return capitalInfo, err
	}
	capitalInfo.Close = resp.Datas.Close
	capitalInfo.ListTime = resp.Datas.ListTime
	capitalInfo.PositiveNew = resp.Datas.CommentNew.PositiveNew
	capitalInfo.UnPositiveNew = resp.Datas.CommentNew.UnPositiveNew
	return capitalInfo, nil
}
