package services

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EchoHandler(ctx *gin.Context) {
	reqBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	log.Printf("[echo]:\n%s", reqBytes)
}
