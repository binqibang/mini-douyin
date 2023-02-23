package business

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/binqibang/mini-douyin/model"
	"github.com/dgrijalva/jwt-go/v4"
)

var uerInitOnce sync.Once
var userDao model.UserDao

// 初始化数据库连接和用户操作对象
func initUserService() {
	userDao = model.UserDao{}
}

// GetUserInfo 获取用户信息
func GetUserInfo(userid string) (*model.User, error) {
	uerInitOnce.Do(initUserService)
	//将用户id字符串转换为int64类型
	userId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		return nil, err
	}
	user, err := userDao.QueryByUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, err
}

// TODO: 完善设计

// CheckIsFollow 检查host用户是否关注user用户
func CheckIsFollow(userid string, hostId int64) (bool, error) {
	uerInitOnce.Do(initUserService)
	//检查输入id是否合法
	userId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		return false, err
	}
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

func GetUserById(userId int64) model.User {
	user, err := userDao.QueryByUserById(userId)
	if err != nil {
		return model.User{}
	}
	return *user
}

func CreateUser(user *model.User) error {
	uerInitOnce.Do(initUserService)
	return userDao.CreateUser(user)
}

// Check_login 登录验证
func Check_login(username string, password string) (*model.User, error) {
	uerInitOnce.Do(initUserService)
	user, err := userDao.QueryByUserByUsername(username, Encrypt(password))
	return user, err
}

// Authentication 验证token
func Authentication(token string, uidPost string) (bool, error) {
	uid, err := ParseToken(token, "bcdedit")
	if uid == uidPost {
		return true, nil
	}
	return false, err
}

func CreateToken(uid int64) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": strconv.FormatInt(uid, 10),
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	token, err := at.SignedString([]byte("bcdedit"))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if uid, ok := claim.Claims.(jwt.MapClaims)["uid"].(string); ok {
		return uid, nil
	}
	return "", fmt.Errorf("fail parse")
}
