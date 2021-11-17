// 检索股票

package routes

import (
	"net/http"

	"smzdtz-server/datacenter"
	"smzdtz-server/datacenter/eastmoney"
	"smzdtz-server/datacenter/sina"
	"smzdtz-server/datacenter/zsxg"

	"github.com/gin-gonic/gin"
)

func AddStockRoutes(rg *gin.RouterGroup) {
	stock := rg.Group("/stock")
	// 东方财富 - 最新指标
	stock.GET("/eastmoney/getIndicator", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := datacenter.EastMoney.QueryIndicator(c, params.Code)
		if err != nil {
			c.JSON(http.StatusOK, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": data})
	})
	// 东方财富 - 价值评估
	stock.GET("/eastmoney/getJiaZhiPingGu", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := datacenter.EastMoney.QueryJiaZhiPingGu(c, params.Code)
		if err != nil {
			c.JSON(http.StatusOK, data)
			return
		}
		c.JSON(http.StatusOK, data)
	})
	// 东方财富 - 综合评价
	stock.GET("/eastmoney/getZongHePingJia", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := datacenter.EastMoney.QueryZongHePingJia(c, params.Code)
		if err != nil {
			c.JSON(http.StatusOK, data)
			return
		}
		c.JSON(http.StatusOK, data)
	})
	// 东方财富 - 股票相关资讯
	stock.GET("/eastmoney/getStockNews", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		var data = gin.H{
			"data": []eastmoney.Article{},
		}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := datacenter.EastMoney.GetStockNews(c, params.Code)
		if err != nil {
			c.JSON(http.StatusOK, data)
		}
		data["data"] = resp
		c.JSON(http.StatusOK, data)
	})
	// Search 检索股票
	stock.GET("/search", func(c *gin.Context) {
		var params struct {
			Keyword string `form:"keyword"`
		}
		var data = gin.H{
			"data": []sina.SearchResult{},
		}

		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//s := core.NewSearcher(c)
		//k := []string{"招商银行", "贵州茅台", "600038"}
		//results, err := s.SearchStocks(c, k)
		results, err := datacenter.Sina.KeywordSearch(c, params.Keyword)
		if err != nil {
			data["error"] = err.Error()
			c.JSON(http.StatusOK, data)
			return
		}
		// StockInfoList 股票列表
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
		data["data"] = result
		c.JSON(http.StatusOK, data)
	})
	// 东方财富 - 股票概况
	stock.GET("/eastmoney/getProfile", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		var data = eastmoney.CompanyProfile{}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, err := datacenter.EastMoney.QueryCompanyProfile(c, params.Code)
		if err != nil {
			c.JSON(http.StatusOK, data)
			return
		}
		data = resp
		c.JSON(http.StatusOK, data)
	})
	// 东方财富 - 估值状态
	stock.GET("/eastmoney/getValuationStatus", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		// var data := map[string]string{}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := datacenter.EastMoney.QueryValuationStatus(c, params.Code)
		if err != nil {
			c.JSON(http.StatusOK, data)
			return
		}
		c.JSON(http.StatusOK, data)
	})
	// 芝士财富 - 股票概况
	stock.GET("/zsxg/info", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		var data = zsxg.CapitalInfo{}
		// var data = gin.H{
		// 	"Code":    200,
		// 	"Message": "success",
		// 	"Data":    zsxg.CapitalInfo{},
		// }
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		resp, _ := datacenter.Zsxg.QueryCommentNew(c, params.Code)
		// if err != nil {
		// 	data["Message"] = err.Error()
		// 	c.JSON(http.StatusOK, data)
		// 	return
		// }
		data = resp
		// data["Data"] = resp
		c.JSON(http.StatusOK, data)
	})
	stock.GET("/cninfo/stock", func(c *gin.Context) {
		resp, _ := datacenter.CnInfo.StockList(c)
		c.JSON(http.StatusOK, resp)
	})
}
