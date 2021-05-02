package main

import (
	"github.com/1000Delta/wifi-locate/pkg/gateway"
	"github.com/gin-gonic/gin"
)

func main() {
	// init
	e := gin.Default()

	gateway.RegisterRouter(e)

	e.Run(":8080")
}
