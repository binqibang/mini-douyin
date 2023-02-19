package controller

import (
	"github.com/binqibang/mini-douyin/business"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	userID := c.Query("user_id")
	toUserID := c.Query("to_user_id")
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if _, exist := usersLoginInfo[token]; exist {
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "解析失败"})
		}
		switch {
		// 关注
		case actionType == 1:
			go business.AddFollowRelation(userID, toUserID)
			log.Println("关注成功")
		// 取关
		case actionType == 2:
			go business.DeleteFollowRelation(userID, toUserID)
			log.Println("取关成功")
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: []User{DemoUser},
	})
}
