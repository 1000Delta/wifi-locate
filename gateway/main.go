package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// init
	e := gin.Default()

	registerRouter(e)

	e.Run(":8080")
}
