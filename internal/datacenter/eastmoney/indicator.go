package eastmoney

import (
	"context"
	"fmt"
	"smzdtz-server/pkg/http"
	"strings"
)

// 获取最新指标

type RespIndicator struct {
	Version string `json:"version"`
	Result  struct {
		Pages int64   `json:"pages"`
		Data  []Datum `json:"data"`
		Count int64   `json:"count"`
	}
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

type Datum struct {
	Secucode                string      `json:"SECUCODE"`
	SecurityCode            string      `json:"SECURITY_CODE"`
	OrgCode                 string      `json:"ORG_CODE"`
	SecurityNameAbbr        string      `json:"SECURITY_NAME_ABBR"`
	Eps                     float64     `json:"EPS"`
	Bvps                    float64     `json:"BVPS"`
	TotalOperateIncome      float64     `json:"TOTAL_OPERATE_INCOME"`
	OperateIncomeRatio      float64     `json:"OPERATE_INCOME_RATIO"`
	Netprofit               float64     `json:"NETPROFIT"`
	NetprofitRatio          float64     `json:"NETPROFIT_RATIO"`
	GrossProfitRatio        float64     `json:"GROSS_PROFIT_RATIO"`
	Npr                     float64     `json:"NPR"`
	Roe                     float64     `json:"ROE"`
	Debt                    float64     `json:"DEBT"`
	CapitalAdequacyRatio    interface{} `json:"CAPITAL_ADEQUACY_RATIO"`
	NPL                     interface{} `json:"NPL"`
	AllowanceNPL            interface{} `json:"ALLOWANCE_NPL"`
	Commnreve               float64     `json:"COMMNREVE"`
	CommnreveYoy            float64     `json:"COMMNREVE_YOY"`
	EarnedPremium           interface{} `json:"EARNED_PREMIUM"`
	CompensateExpense       interface{} `json:"COMPENSATE_EXPENSE"`
	SurrenderRateLife       interface{} `json:"SURRENDER_RATE_LIFE"`
	SolvencyAr              interface{} `json:"SOLVENCY_AR"`
	ResearchExpense         interface{} `json:"RESEARCH_EXPENSE"`
	RsexpenseRatio          interface{} `json:"RSEXPENSE_RATIO"`
	ResearchNum             interface{} `json:"RESEARCH_NUM"`
	ResearchNumRatio        interface{} `json:"RESEARCH_NUM_RATIO"`
	TotalShares             int64       `json:"TOTAL_SHARES"`
	ASharesEquity           interface{} `json:"A_SHARES_EQUITY"`
	FreeAShares             int64       `json:"FREE_A_SHARES"`
	PledgeRatio             float64     `json:"PLEDGE_RATIO"`
	Goodwill                float64     `json:"GOODWILL"`
	CdrShare                interface{} `json:"CDR_SHARE"`
	CdrConvertRatio         interface{} `json:"CDR_CONVERT_RATIO"`
	MarketcapA              interface{} `json:"MARKETCAP_A"`
	BSharesEquity           interface{} `json:"B_SHARES_EQUITY"`
	MarketcapB              interface{} `json:"MARKETCAP_B"`
	FreeBShares             interface{} `json:"FREE_B_SHARES"`
	BUnit                   interface{} `json:"B_UNIT"`
	Securitytype            string      `json:"SECURITYTYPE"`
	Trademarket             string      `json:"TRADEMARKET"`
	DateType                string      `json:"DATE_TYPE"`
	IsProfit                string      `json:"IS_PROFIT"`
	OrgType                 string      `json:"ORG_TYPE"`
	IsVoteDiff              string      `json:"IS_VOTE_DIFF"`
	ListingState            string      `json:"LISTING_STATE"`
	PEDynamicSource         string      `json:"PE_DYNAMIC_SOURCE"`
	PbNoticeSource          string      `json:"PB_NOTICE_SOURCE"`
	EpsSource               string      `json:"EPS_SOURCE"`
	BvpsSource              string      `json:"BVPS_SOURCE"`
	ToiSource               string      `json:"TOI_SOURCE"`
	OirSource               string      `json:"OIR_SOURCE"`
	NetprofitSource         string      `json:"NETPROFIT_SOURCE"`
	NetprofitRatioSource    string      `json:"NETPROFIT_RATIO_SOURCE"`
	GprSource               string      `json:"GPR_SOURCE"`
	NprSource               string      `json:"NPR_SOURCE"`
	RoeSource               string      `json:"ROE_SOURCE"`
	DebtSource              string      `json:"DEBT_SOURCE"`
	NPLSource               string      `json:"NPL_SOURCE"`
	AllowanceNPLSource      string      `json:"ALLOWANCE_NPL_SOURCE"`
	CarSource               string      `json:"CAR_SOURCE"`
	CommnreveSource         string      `json:"COMMNREVE_SOURCE"`
	CommnreveYoySource      string      `json:"COMMNREVE_YOY_SOURCE"`
	EarnedPremiumSource     string      `json:"EARNED_PREMIUM_SOURCE"`
	CompensateExpenseSource string      `json:"COMPENSATE_EXPENSE_SOURCE"`
	SrlSource               string      `json:"SRL_SOURCE"`
	SolvencyArSource        string      `json:"SOLVENCY_AR_SOURCE"`
	ResearchExpenseSource   string      `json:"RESEARCH_EXPENSE_SOURCE"`
	RsexpenseRatioSource    string      `json:"RSEXPENSE_RATIO_SOURCE"`
	ResearchNumSource       string      `json:"RESEARCH_NUM_SOURCE"`
	RnrSource               string      `json:"RNR_SOURCE"`
	TotalSharesSource       string      `json:"TOTAL_SHARES_SOURCE"`
	TmcSource               string      `json:"TMC_SOURCE"`
	CdrShareSource          string      `json:"CDR_SHARE_SOURCE"`
	CcrSource               string      `json:"CCR_SOURCE"`
	AseSource               string      `json:"ASE_SOURCE"`
	FasSource               string      `json:"FAS_SOURCE"`
	McfaSource              string      `json:"MCFA_SOURCE"`
	PledgeRatioSource       string      `json:"PLEDGE_RATIO_SOURCE"`
	MCASource               string      `json:"MCA_SOURCE"`
	GoodwillSource          string      `json:"GOODWILL_SOURCE"`
	BseSource               string      `json:"BSE_SOURCE"`
	McbSource               string      `json:"MCB_SOURCE"`
	FbsSource               string      `json:"FBS_SOURCE"`
	McfbSource              string      `json:"MCFB_SOURCE"`
	EquityNewReport         float64     `json:"EQUITY_NEW_REPORT"`
	PEDynamic               float64     `json:"PE_DYNAMIC"`
	PbNewNotice             float64     `json:"PB_NEW_NOTICE"`
	TotalMarketCap          int64       `json:"TOTAL_MARKET_CAP"`
	MarketcapFreeB          int64       `json:"MARKETCAP_FREE_B"`
	PEStatic                float64     `json:"PE_STATIC"`
	PETtm                   float64     `json:"PE_TTM"`
	MarketcapFreeA          int64       `json:"MARKETCAP_FREE_A"`
	F2                      float64     `json:"f2"`
	F18                     float64     `json:"f18"`
	PbMrqRealtime           float64     `json:"PB_MRQ_REALTIME"`
}

func (e EastMoney) QueryIndicator(ctx context.Context, secuCode string) ([]Datum, error) {
	apiurl := "https://datacenter.eastmoney.com/securities/api/data/v1/get"
	params := map[string]string{
		"client":       "APP",
		"source":       "HSF10",
		"reportName":   "RPT_DMSK_NEWINDICATOR",
		"columns":      "SECUCODE,SECURITY_CODE,ORG_CODE,SECURITY_NAME_ABBR,EPS,BVPS,TOTAL_OPERATE_INCOME,OPERATE_INCOME_RATIO,NETPROFIT,NETPROFIT_RATIO,GROSS_PROFIT_RATIO,NPR,ROE,DEBT,CAPITAL_ADEQUACY_RATIO,NPL,ALLOWANCE_NPL,COMMNREVE,COMMNREVE_YOY,EARNED_PREMIUM,COMPENSATE_EXPENSE,SURRENDER_RATE_LIFE,SOLVENCY_AR,RESEARCH_EXPENSE,RSEXPENSE_RATIO,RESEARCH_NUM,RESEARCH_NUM_RATIO,TOTAL_SHARES,A_SHARES_EQUITY,FREE_A_SHARES,PLEDGE_RATIO,GOODWILL,CDR_SHARE,CDR_CONVERT_RATIO,MARKETCAP_A,B_SHARES_EQUITY,MARKETCAP_B,FREE_B_SHARES,B_UNIT,SECURITYTYPE,TRADEMARKET,DATE_TYPE,IS_PROFIT,ORG_TYPE,IS_VOTE_DIFF,LISTING_STATE,PE_DYNAMIC_SOURCE,PB_NOTICE_SOURCE,EPS_SOURCE,BVPS_SOURCE,TOI_SOURCE,OIR_SOURCE,NETPROFIT_SOURCE,NETPROFIT_RATIO_SOURCE,GPR_SOURCE,NPR_SOURCE,ROE_SOURCE,DEBT_SOURCE,NPL_SOURCE,ALLOWANCE_NPL_SOURCE,CAR_SOURCE,COMMNREVE_SOURCE,COMMNREVE_YOY_SOURCE,EARNED_PREMIUM_SOURCE,COMPENSATE_EXPENSE_SOURCE,SRL_SOURCE,SOLVENCY_AR_SOURCE,RESEARCH_EXPENSE_SOURCE,RSEXPENSE_RATIO_SOURCE,RESEARCH_NUM_SOURCE,RNR_SOURCE,TOTAL_SHARES_SOURCE,TMC_SOURCE,CDR_SHARE_SOURCE,CCR_SOURCE,ASE_SOURCE,FAS_SOURCE,MCFA_SOURCE,PLEDGE_RATIO_SOURCE,MCA_SOURCE,GOODWILL_SOURCE,BSE_SOURCE,MCB_SOURCE,FBS_SOURCE,MCFB_SOURCE,EQUITY_NEW_REPORT",
		"quoteColumns": "f9~01~SECURITY_CODE~PE_DYNAMIC,f23~01~SECURITY_CODE~PB_NEW_NOTICE,f20~01~SECURITY_CODE~TOTAL_MARKET_CAP,f21~01~SECURITY_CODE~MARKETCAP_FREE_B,f114~01~SECURITY_CODE~PE_STATIC,f115~01~SECURITY_CODE~PE_TTM,f21~01~SECURITY_CODE~MARKETCAP_FREE_A,f2~01~SECURITY_CODE~f2,f18~01~SECURITY_CODE~f18",
		"filter":       fmt.Sprintf(`(SECUCODE="%s")`, strings.ToUpper(secuCode)),
		"pageSize":     "200",
		"pageNumber":   "1",
	}
	apiurl, err := http.NewHTTPGetURLWithQueryString(ctx, apiurl, params)
	if err != nil {
		return nil, err
	}
	resp := RespIndicator{}

	err = http.HTTPGET(ctx, e.HTTPClient, apiurl, nil, &resp)
	if err != nil {
		return nil, err
	}
	// if resp.Code != 0 {
	// 	return nil, fmt.Errorf("%s %#v", secuCode, resp)
	// }
	return resp.Result.Data, nil
}

// // 带市场股票编码 600585.SH
// Secucode string `json:"SECUCODE"`
// // 股票编码 600585
// SecurityCode string `json:"SECURITY_CODE"`
// // 10002644
// OrgCode string `json:"ORG_CODE"`
// // 上市名称 10002644
// SecurityNameAbbr string `json:"SECURITY_NAME_ABBR"`
// // 每股收益（元）
// Eps float64 `json:"EPS"`
// // 每股净资产（元）
// Bvps float64 `json:"BVPS"`
// // 营业总收入（元）
// TotalOperateIncome int64 `json:"TOTAL_OPERATE_INCOME"`
// // 营收同比
// OperateIncomeRatio float64 `json:"OPERATE_INCOME_RATIO"`
// // 净利润（元）
// Netprofit int64 `json:"NETPROFIT"`
// // 净利润同比
// NetprofitRatio float64 `json:"NETPROFIT_RATIO"`
// // 毛利率
// GrossProfitRatio float64 `json:"GROSS_PROFIT_RATIO"`
// // 净利率
// Npr float64 `json:"NPR"`
// // 净资产收益率
// Roe float64 `json:"ROE"`
// // 负债率
// Debt                 float64     `json:"DEBT"`
// CapitalAdequacyRatio interface{} `json:"CAPITAL_ADEQUACY_RATIO"`
// NPL                  interface{} `json:"NPL"`
// AllowanceNPL         interface{} `json:"ALLOWANCE_NPL"`
// Commnreve            interface{} `json:"COMMNREVE"`
// CommnreveYoy         interface{} `json:"COMMNREVE_YOY"`
// EarnedPremium        interface{} `json:"EARNED_PREMIUM"`
// CompensateExpense    interface{} `json:"COMPENSATE_EXPENSE"`
// SurrenderRateLife    interface{} `json:"SURRENDER_RATE_LIFE"`
// SolvencyAr           interface{} `json:"SOLVENCY_AR"`
// ResearchExpense      interface{} `json:"RESEARCH_EXPENSE"`
// RsexpenseRatio       interface{} `json:"RSEXPENSE_RATIO"`
// ResearchNum          interface{} `json:"RESEARCH_NUM"`
// ResearchNumRatio     interface{} `json:"RESEARCH_NUM_RATIO"`
// // 总股本
// TotalShares int64 `json:"TOTAL_SHARES"`
// // A股股本
// ASharesEquity int64 `json:"A_SHARES_EQUITY"`
// // 流通A股
// FreeAShares int64 `json:"FREE_A_SHARES"`

// PledgeRatio interface{} `json:"PLEDGE_RATIO"`
// // 商誉规模
// Goodwill        int64       `json:"GOODWILL"`
// CdrShare        interface{} `json:"CDR_SHARE"`
// CdrConvertRatio interface{} `json:"CDR_CONVERT_RATIO"`
// // 流A市值（元）
// MarketcapA    float64     `json:"MARKETCAP_A"`
// BSharesEquity interface{} `json:"B_SHARES_EQUITY"`
// MarketcapB    interface{} `json:"MARKETCAP_B"`
// FreeBShares   int64       `json:"FREE_B_SHARES"`
// BUnit         interface{} `json:"B_UNIT"`
// Securitytype  string      `json:"SECURITYTYPE"`
// Trademarket   string      `json:"TRADEMARKET"`
// // 数据来源类型 2021三季报
// DateType     string `json:"DATE_TYPE"`
// IsProfit     string `json:"IS_PROFIT"`
// OrgType      string `json:"ORG_TYPE"`
// IsVoteDiff   string `json:"IS_VOTE_DIFF"`
// ListingState string `json:"LISTING_STATE"`
// // 基于2021三季报净利润
// PEDynamicSource string `json:"PE_DYNAMIC_SOURCE"`
// // 基于2021-10-29公告股东权益
// PbNoticeSource string `json:"PB_NOTICE_SOURCE"`
// // 基于2021三季报净利润
// EpsSource string `json:"EPS_SOURCE"`
// // 基于2021-10-29公告股东权益
// BvpsSource              string `json:"BVPS_SOURCE"`
// ToiSource               string `json:"TOI_SOURCE"`
// OirSource               string `json:"OIR_SOURCE"`
// NetprofitSource         string `json:"NETPROFIT_SOURCE"`
// NetprofitRatioSource    string `json:"NETPROFIT_RATIO_SOURCE"`
// GprSource               string `json:"GPR_SOURCE"`
// NprSource               string `json:"NPR_SOURCE"`
// RoeSource               string `json:"ROE_SOURCE"`
// DebtSource              string `json:"DEBT_SOURCE"`
// NPLSource               string `json:"NPL_SOURCE"`
// AllowanceNPLSource      string `json:"ALLOWANCE_NPL_SOURCE"`
// CarSource               string `json:"CAR_SOURCE"`
// CommnreveSource         string `json:"COMMNREVE_SOURCE"`
// CommnreveYoySource      string `json:"COMMNREVE_YOY_SOURCE"`
// EarnedPremiumSource     string `json:"EARNED_PREMIUM_SOURCE"`
// CompensateExpenseSource string `json:"COMPENSATE_EXPENSE_SOURCE"`
// SrlSource               string `json:"SRL_SOURCE"`
// SolvencyArSource        string `json:"SOLVENCY_AR_SOURCE"`
// ResearchExpenseSource   string `json:"RESEARCH_EXPENSE_SOURCE"`
// RsexpenseRatioSource    string `json:"RSEXPENSE_RATIO_SOURCE"`
// ResearchNumSource       string `json:"RESEARCH_NUM_SOURCE"`
// RnrSource               string `json:"RNR_SOURCE"`
// TotalSharesSource       string `json:"TOTAL_SHARES_SOURCE"`
// TmcSource               string `json:"TMC_SOURCE"`
// CdrShareSource          string `json:"CDR_SHARE_SOURCE"`
// CcrSource               string `json:"CCR_SOURCE"`
// AseSource               string `json:"ASE_SOURCE"`
// FasSource               string `json:"FAS_SOURCE"`
// McfaSource              string `json:"MCFA_SOURCE"`
// PledgeRatioSource       string `json:"PLEDGE_RATIO_SOURCE"`
// MCASource               string `json:"MCA_SOURCE"`
// GoodwillSource          string `json:"GOODWILL_SOURCE"`
// BseSource               string `json:"BSE_SOURCE"`
// McbSource               string `json:"MCB_SOURCE"`
// FbsSource               string `json:"FBS_SOURCE"`
// McfbSource              string `json:"MCFB_SOURCE"`
// EquityNewReport         int64  `json:"EQUITY_NEW_REPORT"`
// // 市盈率（动）
// PEDynamic float64 `json:"PE_DYNAMIC"`
// // 市净率（最新公告）
// PbNewNotice    float64 `json:"PB_NEW_NOTICE"`
// TotalMarketCap int64   `json:"TOTAL_MARKET_CAP"`
// MarketcapFreeB int64   `json:"MARKETCAP_FREE_B"`
// // 市盈率（静）
// PEStatic float64 `json:"PE_STATIC"`
// // 市盈率（TTM）
// PETtm          float64 `json:"PE_TTM"`
// MarketcapFreeA int64   `json:"MARKETCAP_FREE_A"`
// // 收盘价
// F2  float64 `json:"f2"`
// F18 float64 `json:"f18"`
// // 市净率（MRQ）
// PbMrqRealtime float64 `json:"PB_MRQ_REALTIME"`
