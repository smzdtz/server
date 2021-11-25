package fund

import (
	"net/http"
	"smzdtz-server/internal/datacenter"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Code string `form:"code" binding:"required"`
}

// 东方财富 - 基金信息
func GetFundInfo(c *gin.Context) {
	params := Params{}
	err := c.ShouldBindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := datacenter.EastMoney.QueryFundInfo(c, params.Code)

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
