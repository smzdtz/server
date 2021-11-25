package stock

import (
	"net/http"
	"smzdtz-server/internal/datacenter"
	"smzdtz-server/internal/datacenter/sina"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Code string `form:"code" binding:"required"`
}

// 东方财富 - 最新指标
func GetIndicator(c *gin.Context) {
	params := Params{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.EastMoney.QueryIndicator(c, params.Code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

// 东方财富 - 价值评估
func GetJiaZhiPingGu(c *gin.Context) {
	params := Params{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.EastMoney.QueryJiaZhiPingGu(c, params.Code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

// 东方财富 - 综合评价
func GetZongHePingJia(c *gin.Context) {
	params := Params{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.EastMoney.QueryZongHePingJia(c, params.Code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

//  东方财富 - 股票相关资讯
func GetStockNews(c *gin.Context) {
	params := Params{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.EastMoney.GetStockNews(c, params.Code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

// 检索股票
func SearchStock(c *gin.Context) {
	var params struct {
		Keyword string `form:"keyword"`
	}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	results, err := datacenter.Sina.KeywordSearch(c, params.Keyword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	type StockInfoList []sina.SearchResult
	result := StockInfoList{}
	for i, v := range results {
		resp, _ := datacenter.EastMoney.QueryCompanyProfile(c, v.Secucode)
		if resp.Name != "" {
			result = append(result, v)
		}
		// 性能问题临时解决方案
		if i == 4 {
			break
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    result,
	})
}

//  东方财富 - 股票概况
func GetProfile(c *gin.Context) {
	params := Params{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.EastMoney.QueryCompanyProfile(c, params.Code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

//  芝士财富 - 股票概况
func GetCommentNew(c *gin.Context) {
	params := Params{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.Zsxg.QueryCommentNew(c, params.Code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

// 浪潮资讯 - 股票列表
func GetStockList(c *gin.Context) {
	data, err := datacenter.CnInfo.StockList(c)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

//  东方财富 - 股票概况
func GetFreeHolderse(c *gin.Context) {
	params := Params{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.EastMoney.QueryFreeHolders(c, params.Code)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

//  东方财富-实时行情
// 0.399006 创业板指
// 0.399001 深证成指
// 1.000001 上证指数
func GetStockTrends(c *gin.Context) {
	var params struct {
		SecId string `form:"secid"`
	}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.EastMoney.QueryStockTrends(c, params.SecId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}
