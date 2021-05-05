package model

import (
	"log"

	"github.com/1000Delta/wifi-locate/pkg/locate"
	"gorm.io/gorm"
)

const (
	// limit vector dimension
	AP_LIMIT = 4
)

type APVector struct {
	gorm.Model
	MapID          uint `gorm:"map_id"`
	I0, I1, I2, I3 int64
	LocX, LocY     float64
}

func (APVector) TableName() string {
	return "ap_vector"
}

// Add the vector of map into DB
func (v APVector) Add() error {
	if ok, err := MapExist(v.MapID); err != nil {
		return err
	} else if !ok {
		// map not found
		log.Printf("%v, mapID = %d", gorm.ErrRecordNotFound, v.MapID)
		return gorm.ErrRecordNotFound
	}
	if err := db.Create(&v).Error; err != nil {
		return err
	}
	return nil
}

func GetVecByMap(m *LocationMap) (vecList []locate.APVector, err error) {
	vecQuery := &APVector{MapID: m.ID}
	err = db.Model(vecQuery).Where(vecQuery).Find(&vecList).Error 
	// TODO 检验跨函数传递 err 是否仍为 nil
	// TODO 检验类型是否可以互相转换
	// no record is not a error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

// implement algo.APVector methods
func (v APVector) GetVecVal(i int) int64 {
	switch i {
	case 0:
		return v.I0
	case 1:
		return v.I1
	case 2:
		return v.I2
	case 3:
		return v.I3
	default:
		return int64(0)
	}
}
func (v APVector) GetLocation() (float64, float64) {
	return v.LocX, v.LocY
}
