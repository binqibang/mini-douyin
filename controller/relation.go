package controller

import (
	"fmt"
	"github.com/binqibang/mini-douyin/business"
	"github.com/gin-gonic/gin"
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
	userID, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	toUserID, err2 := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, err3 := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if _, exist := usersLoginInfo[token]; exist {
		if err1 != nil || err2 != nil || err3 != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Parsing failure"})
		}
		switch {
		// 关注
		case actionType == 1:
			go business.CreateFollowRelation(userID, toUserID)
			var event = FollowEvent{
				userID,
				toUserID,
				"success following",
				true,
			}
			fmt.Printf("Receive Message：%+v\n", event)
		// 取关
		case actionType == 2:
			go business.DeleteFollowRelation(userID, toUserID)
			var event = FollowEvent{
				userID,
				toUserID,
				"success cancling following",
				false,
			}
			fmt.Printf("Receive Message：%+v\n", event)
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userID, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if _, exist := usersLoginInfo[token]; exist {
		if err1 != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Parsing failure"})
		}
		userList := business.GetFollowingList(userID)
		users := make([]User, 0)
		for _, v := range userList {
			user, _ := convUserModel2UserInfo1(&v, userID)
			users = append(users, user)
		}
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: users,
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userID, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if _, exist := usersLoginInfo[token]; exist {
		if err1 != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Parsing failure"})
		}
		userList := business.GetFollowerList(userID)
		users := make([]User, 0)
		for _, v := range userList {
			user, _ := convUserModel2UserInfo1(&v, userID)
			users = append(users, user)
		}
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: users,
		})
	}
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
