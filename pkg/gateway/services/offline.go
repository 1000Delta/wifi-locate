package services

import (
	"net/http"

	"github.com/1000Delta/wifi-locate/pkg/locate"
	"github.com/1000Delta/wifi-locate/pkg/locate/service"
	"github.com/gin-gonic/gin"
)

// AP4LocationReq define that you should
// report location info and AP status info
// of the location.
type ap4LocationReq struct {
	Location   locate.LocationInfo `json:"location"`
	APInfoList []*locate.APInfo
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

type addMapReq struct {
	Name string `json:"name"`
}

type addMapResp struct {
	baseResp
	ID uint `json:"id"`
}

func AddMapHandler(ctx *gin.Context) {
	req := new(addMapReq)
	ctx.BindJSON(req)

	rpcReq := service.CreateMapReq{
		Name: req.Name,
	}

	rpcResp := &service.CreateMapResp{}

	cl, err := getClient()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if err := cl.CreateMap(rpcReq, rpcResp); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp := &addMapResp{ID: rpcResp.MapID}
	resp.Code = 0

	ctx.JSON(http.StatusOK, resp)
}
