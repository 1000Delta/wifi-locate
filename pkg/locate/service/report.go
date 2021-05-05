package service

import (
	"github.com/1000Delta/wifi-locate/pkg/locate"
	"github.com/1000Delta/wifi-locate/pkg/locate/model"
)

type CollectReq struct {
	MapID  int
	APList []*locate.APInfo
}

type CollectResp struct{}

// Collect APInfo prepare to locate device online
func (LocateService) Collect(req CollectReq, resp *CollectResp) error {
	vec := GetMapVector(uint(req.MapID), req.APList)
	if err := vec.Add(); err != nil {
		return err
	}

	return nil
}

type CreateMapReq struct {
	Name string
	Path string
	Width int64
	Height int64
}

type CreateMapResp struct {}

// CreateMap
func (LocateService) CreateMap(req CreateMapReq, resp *CreateMapResp) error {
	m := &model.LocationMap{
		Name: req.Name,
		Path: req.Path,
		Width: req.Width,
		Height: req.Height,
	}
	if err := m.Add(); err != nil {
		return err
	}

	return nil
}
