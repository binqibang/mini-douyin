package controller

import (
	"github.com/binqibang/mini-douyin/business"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	uid := c.Query("user_id")

	if exist, err := business.Authentication(token, uid); exist {
		//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		//	return
		//}
		//if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			user, _ := business.GetUserInfo(uid)
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id: 1,
					User: User{
						Id:            user.UserID,
						Name:          user.Username,
						FollowerCount: user.FollowerCount,
						FollowCount:   user.FollowCount,
						IsFollow:      user.IsFollow},
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
