package services

import (
	"log"
	"net/http"
	"net/rpc"

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

// LocateHandler provide frontend to compute location info
func LocateHandler(c *gin.Context) {

	req := NewLocateReq()
	c.BindJSON(req)

	// rpc call
	locateReq := service.LocateReq{
		MapID:  req.MapID,
		APList: req.ScanList,
	}
	var location *locate.LocationInfo
	cl, err := getClient()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := cl.Locate(locateReq, location); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	locationData := NewLocateResp(location)
	c.JSON(http.StatusOK, locationData)
}

func getClient() (*service.Client, error) {
	// 添加重连，异步执行
	if err := client.Beat(); err == rpc.ErrShutdown {
		go func() { client = service.DefaultClient() }()
		return nil, err
	}

	return client, nil
}

func init() {
	client = service.DefaultClient()
	if client != nil {
		log.Println("connect locate service successed.")
	}
}
