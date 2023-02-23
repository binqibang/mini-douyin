package controller

import (
	"github.com/binqibang/mini-douyin/config"
	"github.com/binqibang/mini-douyin/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
// 获取视频id，把视频属性改为favorite
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

type VideoList struct{}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: vl,
	})
}
