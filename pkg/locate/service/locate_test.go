package service

import (
	"testing"

	"github.com/1000Delta/wifi-locate/pkg/locate"
)

func TestLocate(t *testing.T) {
	testCases := []struct {
		desc   string
		input  locate.APInfoList
		output *locate.LocationInfo
	}{
		{
			desc: "default",
			input: locate.APInfoList{
				{
					BSSID: "Test",
					RSSI:  "-95dbm",
				},
			},
			output: &locate.LocationInfo{1, 1},
		},
	}

	locate := new(Locate)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			locate.Locate(tC.input, tC.output)
		})
	}
}
