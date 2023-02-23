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
	//测试
	user_like := model.UserLike{}
	user_id, _ := strconv.Atoi(c.Query("user_id"))
	temp, _ := user_like.QueryUserLikeById(int64(user_id))

	conf, _ := config.LoadConfig("./config/settings_dev.yml")
	ip := conf.App.Address
	vl := []Video{}
	videoDao := model.VideoDao{}
	for i := 0; i < len(temp); i++ {
		videoTemp, _ := videoDao.QueryVideoById(temp[i].VideoID)

		video := Video{}
		video.Id = videoTemp.VideoId
		video.Author.Name = videoTemp.Author
		video.Title = videoTemp.Title
		video.CoverUrl = ip + "/douyin/feed_photo/?path=" + videoTemp.CoverUrl
		video.PlayUrl = ip + "/douyin/feed_video/?path=" + videoTemp.PlayUrl
		video.FavoriteCount = videoTemp.FavoriteCount
		video.CommentCount = videoTemp.CommentCount
		video.IsFavorite = videoTemp.IsFavorite

		vl = append(vl, video)
		log.Printf(strconv.Itoa(int(video.FavoriteCount)))
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: vl,
	})
}
