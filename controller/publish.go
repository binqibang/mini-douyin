package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/binqibang/mini-douyin/business"
	"os"
	"path/filepath"
	"time"
)
type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	//连接数据库
	//存入日期，author

	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "title不能为空",
		})
		return
	}

	token := c.PostForm("token")
	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "token不能为空",
		})
		return
	}

	data, err := c.FormFile("data")
	if data == nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "data不能为空",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	now := time.Now()
	finalName := fmt.Sprintf("%d_%d_%d_%s", now.Year(), now.Month(), now.Day(), filename)
	if !IsExistPath("./public/") {
		err = os.MkdirAll("./public/", os.ModePerm)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
	}
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " upload success",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	user_id := c.Query("user_id")
	ok, err := business.Authentication(token, user_id)
	if !ok {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	videos, _ := business.QueryVideosByUserId(user_id)
	user, _ := business.GetUserInfo(user_id)
	vl := []Video{}
	for i := 0; i < len(videos); i++ {
		vl = append(vl, Video{})
		vl[i].Id = videos[i].Id
		vl[i].Author = User{Id: user.UserID, Name: user.Username, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount, IsFollow: user.IsFollow}
		vl[i].CoverUrl = "http://10.128.61.15:8080/douyin/feed_vedio/?path=" + videos[i].CoverUrl
		vl[i].PlayUrl = "http://10.128.61.15:8080/douyin/feed_vedio/?path=" + videos[i].PlayUrl
		vl[i].FavoriteCount = videos[i].FavoriteCount
		vl[i].CommentCount = videos[i].CommentCount
		vl[i].IsFavorite = videos[i].IsFavorite == 1
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: vl,
	})
}
