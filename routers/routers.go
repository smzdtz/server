package routers

import (
	"smzdtz-server/controllers"
	"smzdtz-server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() (r *gin.Engine) {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())
	api := router.Group("api")
	{
		api.GET("/ping", controllers.Test)
	}
	stock := api.Group("stock")
	{
		stock.GET("/getEMProfile", controllers.GetProfile)
		stock.GET("/getEMStockNews", controllers.GetStockNews)
		stock.GET("/getEMZongHePingJia", controllers.GetZongHePingJia)
		stock.GET("getEMFreeHolderse", controllers.GetFreeHolderse)
		stock.GET("/getEMIndicator", controllers.GetIndicator)
		stock.GET("/getEMJiaZhiPingGu", controllers.GetJiaZhiPingGu)

		stock.GET("/getZSXGCommentNew", controllers.GetCommentNew)
		stock.GET("/getCNINFOStockList", controllers.GetStockList)

		stock.GET("search", controllers.SearchStock)
	}
	fund := api.Group("fund")
	{
		fund.GET("/getEMInfo", controllers.GetFundInfo)
	}
	return router
}
