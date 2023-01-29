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

var userIdSequence = int64(1)

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
	username := c.Query("username")
	password := business.Encrypt(c.Query("password"))

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
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	//获取鉴权信息
	token := c.Query("token")

	//获取hostId, 如果不成功，返回
	var hostId int64
	if user, exist := usersLoginInfo[token]; exist {
		hostId = user.Id
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User token is not valid (not int)"},
		})
		return
	}

	//获取用户id
	userIdString := c.Query("user_id")

	if userIdString == "" {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User id is empty"},
		})
		return
	}

	//将用户id字符串转换为int64类型
	userId, err := strconv.ParseInt(userIdString, 10, 64)
	//检查是否转换成功
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User id is not valid (not int)" + err.Error()},
		})
		return
	}
	//从数据库中读取用户信息
	userInfo, err := business.GetUserInfo(userId)

	if err == nil && userInfo != nil {
		//将结构体转换为User结构体
		userInfo, convErr := convUserModel2UserInfo(userInfo, userId, hostId)
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

func convUserModel2UserInfo(userModel *model.User, userId int64, hostId int64) (User, error) {
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
