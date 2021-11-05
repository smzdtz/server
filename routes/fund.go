package routes

import (
	"net/http"
	"smzdtz-server/datacenter"
	"smzdtz-server/datacenter/eastmoney"

	"github.com/gin-gonic/gin"
)

func addFundRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/fund")

	ping.GET("/info", func(c *gin.Context) {
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
