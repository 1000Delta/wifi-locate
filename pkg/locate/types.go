package locate

// APInfo is useful info of WIFI AP for location
type APInfo struct {
	BSSID string `json:"bssid"`
	RSSI  int64  `json:"rssi"` // Recieved Signal Strength Indicator
}

type LocationInfo struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
