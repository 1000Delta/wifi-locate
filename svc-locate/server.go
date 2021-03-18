package main

// WifiInfo is useful info of WIFI AP for location
type WifiInfo struct {
	BSSID string `json:"bssid"`
	RSSI  string `json:"rssi"` // Recieved Signal Strength Indicator
}

type ScanList struct {
	Data []*WifiInfo `json:"data"`
}

type LocationInfo struct {
	
}

func Locate(scanList ScanList, location *LocationInfo) error {

	return nil
}
