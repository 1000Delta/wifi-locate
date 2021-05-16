package algo

import (
	"testing"

	"github.com/1000Delta/wifi-locate/pkg/locate"
	"github.com/1000Delta/wifi-locate/pkg/locate/model"
)

var (
	locator = NewKNNLocator(4)
)

func TestLocate(t *testing.T) {
	testCases := []struct {
		desc   string
		vec    locate.APVector
		dbVecs []locate.APVector
	}{
		{
			desc: "",
			vec: model.APVector{
				MapID: 0,
				I0:    -70,
				I1:    -70,
				I2:    -70,
				I3:    -70,
			},
			dbVecs: []locate.APVector{
				&model.APVector{
					I0:   -75,
					I1:   -75,
					I2:   -75,
					I3:   -75,
					LocX: 0,
					LocY: 0,
				},
				&model.APVector{
					I0:   -60,
					I1:   -60,
					I2:   -80,
					I3:   -80,
					LocX: 0,
					LocY: 1,
				},
				&model.APVector{
					I0:   -80,
					I1:   -80,
					I2:   -60,
					I3:   -60,
					LocX: 1,
					LocY: 0,
				},
				&model.APVector{
					I0:   -60,
					I1:   -60,
					I2:   -60,
					I3:   -60,
					LocX: 1,
					LocY: 1,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Log(locator.Locate(tC.vec, tC.dbVecs))
		})
	}
}
