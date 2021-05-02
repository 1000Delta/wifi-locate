package services

import (
	"net/http"

	"github.com/1000Delta/wifi-locate/pkg/locate"
	"github.com/1000Delta/wifi-locate/pkg/locate/service"
	"github.com/gin-gonic/gin"
)

var client *service.Client

type LocateReq struct {
	ScanList locate.APInfoList `json:"scanList"`
}

func NewLocateReq() *LocateReq {
	return &LocateReq{}
}

type LocateResp struct {
	baseResp
	Data locate.LocationInfo `json:"data,omitempty"`
}

func NewLocateResp(x, y int) *LocateResp {
	return &LocateResp{
		Resp(0, ""),
		locate.LocationInfo{
			X: x,
			Y: y,
		},
	}
}

// Locate provide frontend to compute location info
func Locate(c *gin.Context) {

	req := NewLocateReq()
	c.BindJSON(req)

	// rpc call
	location := &locate.LocationInfo{}
	err := client.Locate(req.ScanList, location)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	locationData := NewLocateResp(location.X, location.Y)
	c.JSON(http.StatusOK, locationData)
}

func init() {
	client = service.NewClient()
}
