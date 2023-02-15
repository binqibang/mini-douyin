package model

import (
	"github.com/binqibang/mini-douyin/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db       *gorm.DB
	initOnce sync.Once
)

// 本地开发时改为自己本机上的绝对路径，便于单元测试
const configPathString = "C:\\Users\\13099\\GolandProjects\\mini-douyin\\config\\settings_dev.yml"

func InitDB(configPath string) *gorm.DB {
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("can not read conf file: %s", err)
	}
	// connect database
	mysqlConf := conf.Database.Mysql
	dsn := mysqlConf.UserName + ":" + mysqlConf.Password
	dsn += "@(" + mysqlConf.Address + ")/" + mysqlConf.Database + "?charset=utf8&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("can not connect database, %s", err)
	}
	return db
}

// 将InitDB函数封装为无参数无返回值函数，方便供initOnce使用
func initDBProxy() {
	db = InitDB(configPathString)
}

func init() {
	initOnce.Do(initDBProxy)
}
