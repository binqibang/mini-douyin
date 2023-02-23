package controller

import (
	"github.com/binqibang/mini-douyin/business"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
// 获取视频id，把视频属性改为favorite
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	uid := c.Query("user_id")
	vid := c.Query("video_id")
	if exist, _ := business.Authentication(token, uid); exist {
		user, _ := business.GetUserInfo(uid)
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	uid := c.Query("user_id")
	user, _ := business.GetUserInfo(uid)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
