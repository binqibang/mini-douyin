package handler

import (
	"github.com/binqibang/mini-douyin/model"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB
var initOnce sync.Once
var userDao model.UserDao

const ConfigPath = "config/settings_dev.yml"

// 初始化数据库连接和用户操作对象
func initUserInfoService() {
	db = model.InitDB(ConfigPath)
	userDao = model.UserDao{}
}

// GetUserInfo 获取用户信息
func GetUserInfo(userid int64) (*model.User, error) {
	initOnce.Do(initUserInfoService)
	user, err := userDao.QueryByUserById(db, userid)
	if err != nil {
		return nil, err
	}
	return user, err
}

// TODO: 完善设计

// CheckIsFollow 检查一个用户是否关注另一用户
func CheckIsFollow() bool {
	return true
}
