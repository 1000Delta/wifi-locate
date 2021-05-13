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

	// offline
	e.POST("/map/add", services.AddMapHandler)

	// locate
	e.POST("/locate", services.LocateHandler)

	// test
	e.POST("/echo", services.EchoHandler)
}
