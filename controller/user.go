package controller

import (

	"github.com/binqibang/mini-douyin/business"
	"github.com/binqibang/mini-douyin/model"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"

)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

//var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := business.Encrypt(c.PostForm("password"))

	user := model.User{Username: username, Password: password}
	err := business.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Register successfully!"},
		})
	}
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := business.Check_login(username, password)

	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		token, _ := business.CreateToken(user.UserID)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.UserID,
			Token:    token,
		})
	}
}

func UserInfo(c *gin.Context) {
	//获取token
	token := c.Query("token")
	userIdString := c.Query("user_id")

	if userIdString == "" || token == "" {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User id is empty"},
		})
		return
	}

	var hostId int64
	exist, err := business.Authentication(token, userIdString)
	if exist {
		hostIdTemp, err := business.ParseToken(token, "bcdedit")
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Parse User token error"},
			})
		}
		hostId, err = strconv.ParseInt(hostIdTemp, 10, 64)
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User token is not valid"},
		})
		return
	}

	//从数据库中读取用户信息
	userInfo, err := business.GetUserInfo(userIdString)

	if err == nil && userInfo != nil {
		//将结构体转换为User结构体
		userInfo, convErr := convUserModel2UserInfo(userInfo, userIdString, hostId)
		if convErr == nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 0},
				User:     userInfo,
			})
		} else {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: err.Error()},
			})
		}
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}

func convUserModel2UserInfo(userModel *model.User, userId string, hostId int64) (User, error) {
	isFollow, checkErr := business.CheckIsFollow(userId, hostId)
	if checkErr != nil {
		return User{}, checkErr
	}
	var userInfo = User{
		Id:            userModel.UserID,
		Name:          userModel.Username,
		FollowCount:   userModel.FollowCount,
		FollowerCount: userModel.FollowerCount,
		IsFollow:      isFollow,
	}
	return userInfo, nil
}

func convUserModel2UserInfo1(userModel *model.User, userId int64) (User, error) {
	isFollow, checkErr := business.CheckIsFollow(userId)
	if checkErr != nil {
		return User{}, checkErr
	}
	var userInfo = User{
		Id:            userModel.UserID,
		Name:          userModel.Username,
		FollowCount:   userModel.FollowCount,
		FollowerCount: userModel.FollowerCount,
		IsFollow:      isFollow,
	}
	return userInfo, nil
}
