package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
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
		DisableDatetimePrecision:  true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDb.SetMaxOpenConns(100) //打开
	sqlDb.SetConnMaxLifetime(time.Second * 30)
	_db = db
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(connWrite)},
		Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)},
		Policy:   dbresolver.RandomPolicy{},
	}))
	Migration()
}
func NewDbClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
