package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type signal struct {
	Device string `json:"device"`
	Strength string `json:"strength"`
}

type LocateReq struct {
	Signals []signal `json:"signals"`
}

func NewLocateReq() *LocateReq {
	return &LocateReq{}
}

func Locate(c *gin.Context) {
	
	req := NewLocateReq()
	c.BindJSON(req)

	c.JSON(http.StatusOK, req)
}