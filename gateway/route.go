package main

import (
	"net/http"

	"github.com/1000Delta/wifi-locate/gateway/services"
	"github.com/gin-gonic/gin"
)

func registerRouter(e *gin.Engine) {
	// 
	e.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello WIFI Location")
	})

	// locate
	e.POST("/locate", services.Locate)
}