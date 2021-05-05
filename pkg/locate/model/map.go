package model

import "gorm.io/gorm"

type LocationMap struct {
	gorm.Model
	Name   string
	Path   string `gorm:"comment:地图图像存储的路径"`
	Width  int64  `gorm:"comment:地图宽度"`
	Height int64  `gorm:"comment:地图高度"`
}

// Add the map into DB
func (m LocationMap) Add() error {
	if err := db.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

// MapExist use count() to check if the map exist.
// 
// if count > 0, return true,
// if error, return false and error.
func MapExist(mapID uint) (bool, error) {
	var mapCount int64
	if err := db.Where("id = ?", mapID).Count(&mapCount).Error; err != nil {
		return false, err
	}
	return mapCount > 0, nil
}

// GetMap return the map info of mapID in DB,
// 
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

// GetAllMap return all location map in DB
func GetAllMap() (maps []*LocationMap, err error) {
	if err = db.Find(&maps).Error; err != nil {
		return
	}
	err = nil
	return
}
