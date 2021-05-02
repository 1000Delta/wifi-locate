package algo

import "github.com/1000Delta/wifi-locate/pkg/locate"

type KNNLocator struct {}

func (KNNLocator) Locate(apList []*locate.APInfo, adListDB []*locate.APInfo) locate.LocationInfo {

	// TODO 实现 KNN 算法
	

	return locate.LocationInfo{X: 0, Y: 0}
}


