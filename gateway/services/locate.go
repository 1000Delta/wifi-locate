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

type location struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type LocateResp struct {
	baseResp
	Data location `json:"data,omitempty"`
}

func NewLocateResp(x, y int) *LocateResp {
	return &LocateResp{
		Resp(0, ""),
		location{x, y},
	}
}
func Locate(c *gin.Context) {
	
	req := NewLocateReq()
	c.BindJSON(req)

	//TODO rpc

	location := NewLocateResp(0, 0)
	c.JSON(http.StatusOK, location)
}