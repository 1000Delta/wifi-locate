package locate

type Locator interface {
	// Locate device by compute current apList1 with apList2 in database.
	Locate(APInfoList, APInfoList) LocationInfo
}
