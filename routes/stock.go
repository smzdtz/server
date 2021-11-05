// 检索股票

package routes

import (
	"net/http"
	"strings"

	"smzdtz-server/datacenter"
	"smzdtz-server/datacenter/eastmoney"
	"smzdtz-server/datacenter/sina"

	"github.com/gin-gonic/gin"
)

func addStockRoutes(rg *gin.RouterGroup) {
	stock := rg.Group("/stock")

	// Search 检索股票
	stock.GET("/search", func(c *gin.Context) {
		var params struct {
			Keyword string `form:"keyword"`
		}
		var data = gin.H{
			"Results": []sina.SearchResult{},
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
			data["Error"] = err.Error()
			c.JSON(http.StatusOK, data)
			return
		}

		// StockInfoList 股票列表
		type StockInfoList []sina.SearchResult
		result := StockInfoList{}
		for _, i := range results {
			// 沪市主板：600、601、603、605
			if strings.HasPrefix(i.Secucode, "600") || strings.HasPrefix(i.Secucode, "601") || strings.HasPrefix(i.Secucode, "603") || strings.HasPrefix(i.Secucode, "605") {
				result = append(result, i)
			}
			// 深市000开头，深市中小板002
			if strings.HasPrefix(i.Secucode, "000") || strings.HasPrefix(i.Secucode, "002") {
				result = append(result, i)
			}
			// 创业板
			if strings.HasPrefix(i.Secucode, "300") {
				result = append(result, i)
			}
			// 科创板
			if strings.HasPrefix(i.Secucode, "688") {
				result = append(result, i)
			}
		}

		data["Results"] = result
		c.JSON(http.StatusOK, data)
	})
	// Profile 股票概况
	stock.GET("/profile", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		var data = gin.H{
			"Results": eastmoney.CompanyProfile{},
		}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resp, err := datacenter.EastMoney.QueryCompanyProfile(c, params.Code)
		if err != nil {
			data["Error"] = err.Error()
			c.JSON(http.StatusOK, data)
			return
		}
		c.JSON(http.StatusOK, resp)
	})
}
