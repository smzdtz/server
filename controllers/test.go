package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
	})

}
