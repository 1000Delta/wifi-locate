package services

import (
	"net/http"

	"github.com/1000Delta/wifi-locate/pkg/locate"
	"github.com/gin-gonic/gin"
)

// AP4LocationReq define that you should
// report location info and AP status info
// of the location.
type ap4LocationReq struct {
	Location   locate.LocationInfo `json:"location"`
	APInfoList locate.APInfoList
}

type ap4LocationResp struct {
	baseResp
}

// ReportAP4Location to do the offline report,
// before online location, collect location data.
func ReportAP4Location(ctx *gin.Context) {
	var req = new(ap4LocationReq)
	ctx.BindJSON(req)

	// TODO 调用信息上报 RPC 服务

	resp := &ap4LocationResp{
		baseResp{
			Code: 0,
		},
	}
	ctx.JSON(http.StatusOK, resp)
}
