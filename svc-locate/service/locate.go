package service

import "log"

// WifiInfo is useful info of WIFI AP for location
type WifiInfo struct {
	BSSID string `json:"bssid"`
	RSSI  string `json:"rssi"` // Recieved Signal Strength Indicator
}

type ScanList []WifiInfo

type LocationInfo struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Locate struct{}

func (l *Locate)Locate(scanList ScanList, location *LocationInfo) error {

	location.X = 1
	location.Y = 1

	log.Printf("fmtinput: %v, output: %v\n", scanList, location)

	return nil
}
