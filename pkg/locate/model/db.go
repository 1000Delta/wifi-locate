package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// InitDB before you call any method about data model.
func InitDB(log logger.Writer) {
	
	dsn := "root:root@tcp(wifi-locate-db:3306)/location_db?charset=utf8mb4&parseTime=True&loc=Local"

	logr := logger.New(log, logger.Config{})
	
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logr})
	if err != nil {
		panic(err)
	}
}
