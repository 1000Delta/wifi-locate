package service

import (
	"log"

	"github.com/1000Delta/wifi-locate/pkg/locate"
)

type Locate struct{}

// Locate compute the location of report WLAN scanList.
func (l *Locate) Locate(scanList []*locate.APInfo, location *locate.LocationInfo) error {

	location.X = 1
	location.Y = 1

	log.Printf("fmtinput: %v, output: %v\n", scanList, location)

	return nil
}
