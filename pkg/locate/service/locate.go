package service

import (
	"errors"
	"log"

	"github.com/1000Delta/wifi-locate/pkg/locate"
	"github.com/1000Delta/wifi-locate/pkg/locate/model"
)

var (
	ErrNoRecord = errors.New("query a not existed record")
)

type LocateService struct {
	locator locate.Locator
}

func NewLocateService(locator locate.Locator) *LocateService {
	return &LocateService{
		locator: locator,
	}
}

type LocateReq struct {
	MapID  uint
	APList []*locate.APInfo
}

// Locate compute the location of report WLAN scanList.
func (l *LocateService) Locate(req LocateReq, location *locate.LocationInfo) error {

	// 转换请求参数到向量
	vec := GetMapVector(req.MapID, req.APList)
	// TODO 实现定位算法调用计算结果流程

	// 查找定位 map
	tgtMap, err := model.GetMap(vec.MapID)
	if err != nil {
		return err
	}
	if tgtMap == nil {
		log.Printf("mapID = %d", vec.MapID)
		return ErrNoRecord
	}

	// 获取 map 离线采集的向量
	dbVecs, err := model.GetVecByMap(tgtMap)
	if err != nil {
		return err
	}
	if len(dbVecs) == 0 {
		log.Printf("mapID = %d", tgtMap.ID)
		return ErrNoRecord
	}

	// 调用定位算法，默认使用 KNN, k=4
	*location = l.locator.Locate(vec, dbVecs)

	log.Printf("req: %v, output: %v\n", req, location)

	return nil
}

// APConvertor convert predefined AP BSSID to APVector fields
var apConvertor map[string]func(*model.APVector, int64)

func InitAPConvertor(markAP []string) {
	if len(markAP) < 2 {
		log.Fatalf("invalid markAP numbers, markAP = %v", markAP)
	}
	// init first 2 convertor
	apConvertor[markAP[0]] = func(tgtVec *model.APVector, rssi int64) {
		tgtVec.I0 = rssi
	}
	apConvertor[markAP[1]] = func(tgtVec *model.APVector, rssi int64) {
		tgtVec.I1 = rssi
	}
	// condition judge and init convertor 3, 4
	if len(markAP) < 3 {
		return
	}
	apConvertor[markAP[2]] = func(tgtVec *model.APVector, rssi int64) {
		tgtVec.I2 = rssi
	}
	if len(markAP) < 4 {
		return
	}
	apConvertor[markAP[3]] = func(tgtVec *model.APVector, rssi int64) {
		tgtVec.I3 = rssi
	}
}

func GetMapVector(mapID uint, apList []*locate.APInfo) model.APVector {
	vec := model.APVector{MapID: mapID}
	for _, ap := range apList {
		if fn, ok := apConvertor[ap.BSSID]; ok {
			fn(&vec, ap.RSSI)
		}
	}

	return vec
}
