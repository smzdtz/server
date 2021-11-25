package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func IPWhiteList() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := viper.GetString("app.whitelist")
		if !strings.Contains(ip, c.ClientIP()) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "Permission denied",
			})
			return
		}
	}
}
