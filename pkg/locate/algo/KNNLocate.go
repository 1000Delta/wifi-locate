package algo

import (
	"math"
	"sort"

	"github.com/1000Delta/wifi-locate/pkg/locate"
)

// APVector define a general mode to get vector value
type APVector interface {
	GetVecVal(dimension int) int64
	GetLocation() (float64, float64)
}

type KNNLocator struct {
	k int // k nums target to locate
}

func NewKNNLocator(k int) *KNNLocator {
	return &KNNLocator{k: k}
}

func (cfg KNNLocator) Locate(scanVec APVector, vecListDB []APVector) locate.LocationInfo {
	diffList := make(apDiffList, len(vecListDB))
	
	// 计算欧式距离
	for i, baseVec := range vecListDB {
		diffList[i] = getDiff(baseVec, scanVec)
	}
	sort.Sort(diffList)

	// 提取排序后最近 k 个目标并计算平均坐标，作为最后的结果坐标
	x, y := getAvgLocation(diffList[:cfg.k])

	return locate.LocationInfo{X: x, Y: y}
}

type apDiff struct {
	// Euclidean Distance
	EuDist  float64
	BaseVec APVector
}

func getDiff(baseVec, dbVec APVector) *apDiff {
	euDist := math.Sqrt(
		math.Pow(float64(baseVec.GetVecVal(0)-dbVec.GetVecVal(0)), 2) +
			math.Pow(float64(baseVec.GetVecVal(1)-dbVec.GetVecVal(1)), 2) +
			math.Pow(float64(baseVec.GetVecVal(2)-dbVec.GetVecVal(2)), 2) +
			math.Pow(float64(baseVec.GetVecVal(3)-dbVec.GetVecVal(3)), 2),
	)
	return &apDiff{
		EuDist: euDist,
	}
}

type apDiffList []*apDiff

// apDiffList sort required methods
func (l apDiffList) Len() int           { return len(l) }
func (l apDiffList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l apDiffList) Less(i, j int) bool { return l[i].EuDist < l[j].EuDist }

// compute average coordinate of target location
func getAvgLocation(diffList apDiffList) (float64, float64) {
	var xall, yall float64
	countF := float64(len(diffList))

	for _, diff := range diffList {
		x, y := diff.BaseVec.GetLocation()
		xall, yall = xall+x, yall+y
	}

	// TODO 确认结果正确性，论文中多计算了一个欧氏距离

	return xall/countF , yall/countF
}
