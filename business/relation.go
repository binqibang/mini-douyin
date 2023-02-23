package business

import (
	"fmt"
	"github.com/binqibang/mini-douyin/model"
)

func CreateFollowRelation(userID int64, toUserID int64) {
	userDao := model.UserDao{}
	user, err1 := userDao.QueryByUserById(userID)
	if err1 != nil {
		_ = fmt.Errorf("User does not exsit")
		return
	}
	userRelation := model.UserRelation{
		UserID:   userID,
		ToUserID: toUserID,
	}
	err2 := user.CreateRelation(&userRelation)
	if err2 != nil {
		_ = fmt.Errorf("Follow failure")
		return
	}
}

func DeleteFollowRelation(userID int64, toUserID int64) {
	userDao := model.UserDao{}
	user, err1 := userDao.QueryByUserById(userID)
	if err1 != nil {
		_ = fmt.Errorf("User does not exsit")
		return
	}
	isFollow, err2 := CheckIsFollow(userID)
	if err2 != nil {
		_ = fmt.Errorf("Fllow error")
		return
	}
	if isFollow != false {
		userRelation, err1 := user.QueryRelation(userID, toUserID)
		if userRelation == nil || err1 != nil {
			return
		}
		err := user.DeleteRelation(userRelation)
		if err != nil {
			_ = fmt.Errorf("Cancel follow failure")
			return
		}
	}
	_ = fmt.Errorf("User does not have follows")
	return
}

func GetFollowingList(userID int64) []model.User {
	userDao := model.UserDao{}
	user, err := userDao.QueryByUserById(userID)
	if err != nil {
		_ = fmt.Errorf("User does not exsit")
		return nil
	}
	isFollow, err2 := CheckIsFollow(userID)
	if err2 != nil {
		_ = fmt.Errorf("Follow error")
		return nil
	}
	if isFollow != true {
		return nil
	}
	userList := user.QueryFollowingList(userID)
	return *userList
}

func GetFollowerList(userID int64) []model.User {
	userDao := model.UserDao{}
	user, err := userDao.QueryByUserById(userID)
	if err != nil {
		_ = fmt.Errorf("User does not exsit")
		return nil
	}
	isFollow, err2 := CheckIsFollow(userID)
	if err2 != nil {
		_ = fmt.Errorf("Follow error")
		return nil
	}
	if isFollow != true {
		return nil
	}
	userList := user.QueryFollowerList(userID)
	return *userList
}
