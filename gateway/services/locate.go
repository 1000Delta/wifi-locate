package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	locateSvc "github.com/1000Delta/wifi-locate/svc-locate/service"
)

var client *locateSvc.Client

type LocateReq struct {
	ScanList locateSvc.ScanList `json:"scanList"`
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

// Locate provide frontend to compute location info
func Locate(c *gin.Context) {
	
	req := NewLocateReq()
	c.BindJSON(req)

	// rpc call
	location := &locateSvc.LocationInfo{}
	err := client.Locate(req.ScanList, location)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	locationData := NewLocateResp(location.X, location.Y)
	c.JSON(http.StatusOK, locationData)
}

func init() {
	client = locateSvc.NewClient()
}