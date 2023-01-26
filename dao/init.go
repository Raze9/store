package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB

func DateBase(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,
		DontSupportRenameColumn:   true,
		DontSupportRenameIndex:    true,
		DisableDatetimePrecision: true,
		SkipInitializeWithVersion: false,
	}),&gorm.Config{
        
	}
}