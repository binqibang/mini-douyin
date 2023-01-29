package business

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/binqibang/mini-douyin/model"
	"sync"
)

var uerInitOnce sync.Once
var userDao model.UserDao

// 初始化数据库连接和用户操作对象
func initUserService() {
	userDao = model.UserDao{}
}

// GetUserInfo 获取用户信息
func GetUserInfo(userid int64) (*model.User, error) {
	uerInitOnce.Do(initUserService)
	user, err := userDao.QueryByUserById(userid)
	if err != nil {
		return nil, err
	}
	return user, err
}

// TODO: 完善设计

// CheckIsFollow 检查host用户是否关注user用户
func CheckIsFollow(userId int64, hostId int64) (bool, error) {
	uerInitOnce.Do(initUserService)
	//检查输入id是否合法
	if userId == 0 || hostId == 0 {
		err := errors.New("UserId or hostId is wrong")
		return false, err
	}
	//TODO: 查询是否关注
	return true, nil
}

const salt = "mini-douyin"

// Encrypt 将用户输入密码使用`MD5`加盐加密后再存储
func Encrypt(pwd string) string {
	return md5Hex(pwd + salt)
}

func md5Hex(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)
	return md5str
}

func CreateUser(user *model.User) error {
	uerInitOnce.Do(initUserService)
	return userDao.CreateUser(user)
}
