package locate

// APVector define a general mode to get vector value
type APVector interface {
	GetVecVal(dimension int) int64
	GetLocation() (float64, float64)
}

// Locator is common interface of all locate algorithms
type Locator interface {
	// Locate device by compute current apList1 with apList2 in database.
	Locate(APVector, []APVector) LocationInfo
}
