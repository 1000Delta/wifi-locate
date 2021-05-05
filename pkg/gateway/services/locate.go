package services

import (
	"net/http"

	"github.com/1000Delta/wifi-locate/pkg/locate"
	"github.com/1000Delta/wifi-locate/pkg/locate/service"
	"github.com/gin-gonic/gin"
)

var client *service.Client

type LocateReq struct {
	MapID    uint             `json:"mapID"`
	ScanList []*locate.APInfo `json:"scanList"`
}

func NewLocateReq() *LocateReq {
	return &LocateReq{}
}

type LocateResp struct {
	baseResp
	Data locate.LocationInfo `json:"data,omitempty"`
}

func NewLocateResp(locInfo *locate.LocationInfo) *LocateResp {
	return &LocateResp{
		Resp(0, ""),
		*locInfo,
	}
}

// Locate provide frontend to compute location info
func Locate(c *gin.Context) {

	req := NewLocateReq()
	c.BindJSON(req)

	// rpc call
	locateReq := service.LocateReq{
		MapID:  req.MapID,
		APList: req.ScanList,
	}
	location := &locate.LocationInfo{}
	err := client.Locate(locateReq, location)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	locationData := NewLocateResp(location)
	c.JSON(http.StatusOK, locationData)
}

func init() {
	client = service.NewClient()
}
