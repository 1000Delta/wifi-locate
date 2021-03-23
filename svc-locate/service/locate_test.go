package service

import (
	"testing"
)

func TestLocate(t *testing.T) {
	testCases := []struct {
		desc   string
		input  ScanList
		output *LocationInfo
	}{
		{
			desc: "default",
			input: ScanList{
				{
					BSSID: "Test",
					RSSI:  "-95dbm",
				},
			},
			output: &LocationInfo{1, 1},
		},
	}

	locate := new(Locate)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			locate.Locate(tC.input, tC.output)
		})
	}
}
