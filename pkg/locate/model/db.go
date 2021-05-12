package model

import (
	"log"
	"time"

	"github.com/1000Delta/wifi-locate/pkg/common/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	initReconnectDelay time.Duration = 1
	// 数据库重新连接最大间隔
	maxReconnectDelay time.Duration = 16
)

var db *gorm.DB

// InitDB before you call any method about data model.
func InitDB(logw logger.Writer) {

	dsn := "root:root@tcp(wifi-locate-db:3306)/location_db?charset=utf8mb4&parseTime=True&loc=Local"

	logr := logger.New(logw, logger.Config{})

	var err error

	// 重试连接 DB
	utils.CallUntilNoError(func(delay time.Duration) error {
		if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logr}); err == nil {
			log.Printf("connect db success.")
			return nil
		}
		log.Printf("connect db error: %v, will reconnect after %ds.", err, delay)
		return err
	},
		initReconnectDelay,
		maxReconnectDelay,
	)

	db.AutoMigrate(&LocationMap{}, &APVector{})
}
