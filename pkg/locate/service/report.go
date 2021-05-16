package service

import (
	"log"

	"github.com/1000Delta/wifi-locate/pkg/locate"
	"github.com/1000Delta/wifi-locate/pkg/locate/model"
)

type CollectReq struct {
	MapID    uint
	APList   []*locate.APInfo
	Location locate.LocationInfo
}

type CollectResp struct{}

// Collect APInfo prepare to locate device online
func (LocateService) Collect(req CollectReq, resp *CollectResp) error {
	// 转换向量
	vec := GetMapVector(req.MapID, req.APList)
	// 记录坐标
	vec.LocX, vec.LocY = req.Location.X, req.Location.Y

	if err := vec.Add(); err != nil {
		log.Printf("add vector error, msg = %v", err)
		return err
	}

	return nil
}

type CreateMapReq struct {
	Name   string
	Path   string
	Width  int64
	Height int64
}

type CreateMapResp struct {
	MapID uint
}

// CreateMap
func (LocateService) CreateMap(req CreateMapReq, resp *CreateMapResp) error {
	m := &model.LocationMap{
		Name:   req.Name,
		Path:   req.Path,
		Width:  req.Width,
		Height: req.Height,
	}
	if err := m.Add(); err != nil {
		log.Printf("add map error, msg = %v", err)
		return err
	}

	*resp = CreateMapResp{MapID: m.ID}

	return nil
}
