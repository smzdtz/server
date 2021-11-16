package routes

import (
	"net/http"
	"smzdtz-server/cron"
	"smzdtz-server/datacenter"
	"smzdtz-server/datacenter/eastmoney"

	"github.com/gin-gonic/gin"
)

func AddFundRoutes(rg *gin.RouterGroup) {
	fund := rg.Group("/fund")

	fund.GET("/sync", func(c *gin.Context) {
		cron.SyncFund()
		c.JSON(http.StatusOK, gin.H{
			"Code":    200,
			"Message": "success",
		})
	})

	fund.GET("/info", func(c *gin.Context) {
		var params struct {
			Code string `form:"code"`
		}
		var data = gin.H{
			"Code":    200,
			"Message": "success",
			"Data":    eastmoney.RespFundInfo{},
		}
		if err := c.ShouldBindQuery(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fundresp, _ := datacenter.EastMoney.QueryFundInfo(c, params.Code)
		data["Data"] = fundresp
		c.JSON(http.StatusOK, data)
	})
}
