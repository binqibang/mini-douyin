package handler

import (
	"errors"
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

// CheckIsFollow 检查host用户是否关注user用户
func CheckIsFollow(userId int64, hostId int64) (bool, error) {
	initOnce.Do(initUserInfoService)
	//检查输入id是否合法
	if userId == 0 || hostId == 0 {
		err := errors.New("UserId or hostId is wrong")
		return false, err
	}
	//TODO: 查询是否关注
	return true, nil
}
