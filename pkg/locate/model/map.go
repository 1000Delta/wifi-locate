package model

import "gorm.io/gorm"

type LocationMap struct {
	gorm.Model
	Name string
}

// GetMap return the map info of mapID in DB,
// if map not existed, return map will be nil.
func GetMap(mapID uint) (*LocationMap, error) {
	var m *LocationMap
	if err := db.Where("id = ?", mapID).First(m).Error; err != nil {
		// no record not means error
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return m, nil
}
