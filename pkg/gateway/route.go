package gateway

import (
	"net/http"

	"github.com/1000Delta/wifi-locate/pkg/gateway/services"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	//
	e.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello WIFI Location")
	})

	// locate
	e.POST("/locate", services.Locate)

	// test
	e.POST("/echo", services.Echo)
}
