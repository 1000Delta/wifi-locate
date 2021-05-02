package locate

// APInfo is useful info of WIFI AP for location
type APInfo struct {
	BSSID string `json:"bssid"`
	RSSI  string `json:"rssi"` // Recieved Signal Strength Indicator
}

// APInfoList is the slice type of APInfo
type APInfoList []APInfo

type LocationInfo struct {
	X int `json:"x"`
	Y int `json:"y"`
}