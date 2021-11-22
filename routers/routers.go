package routers

import (
	"smzdtz-server/controllers"
	"smzdtz-server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter() (r *gin.Engine) {
	router := gin.New()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())
	// 安全，限制接口访问白名单
	// TODO：未来要实现限流
	router.Use(middlewares.IPWhiteList())

	api := router.Group("api")
	{
		api.GET("/ping", controllers.Test)
	}
	install := api.Group("install")
	{
		install.GET("/status", controllers.Status)
		install.POST("/store", controllers.Store)
	}
	// 股票
	stock := api.Group("stock")
	{
		stock.GET("/getEMProfile", controllers.GetProfile)
		stock.GET("/getEMStockNews", controllers.GetStockNews)
		stock.GET("/getEMZongHePingJia", controllers.GetZongHePingJia)
		stock.GET("getEMFreeHolderse", controllers.GetFreeHolderse)
		stock.GET("/getEMIndicator", controllers.GetIndicator)
		stock.GET("/getEMJiaZhiPingGu", controllers.GetJiaZhiPingGu)
		stock.GET("/getEMStockTrends", controllers.GetStockTrends)

		stock.GET("/getZSXGCommentNew", controllers.GetCommentNew)
		stock.GET("/getCNINFOStockList", controllers.GetStockList)

		stock.GET("search", controllers.SearchStock)
	}
	// 基金
	fund := api.Group("fund")
	{
		fund.GET("/getEMInfo", controllers.GetFundInfo)
	}

	return router
}
