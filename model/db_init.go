package model

import (
	"github.com/binqibang/mini-douyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB(configPath string) *gorm.DB {
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("cannot read conf file: %s", err)
	}
	// connect database
	mysqlConf := conf.Database.Mysql
	dsn := mysqlConf.UserName + ":" + mysqlConf.Password
	dsn += "@(" + mysqlConf.Address + ")/" + mysqlConf.Database + "?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("can not connect database, %s", err)
	}
	return db
}
