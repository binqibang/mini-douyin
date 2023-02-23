package controller

import (
	"github.com/binqibang/mini-douyin/business"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList [10]Video `json:"video_list"`
	NextTime  int64     `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
// 推流最近的10条视频功能

func Feed(c *gin.Context) {

	//返回从数据库中查找最近的一条数据
	vl := [10]Video{}
	v, _ := business.GetTenVideos()
	for i := 0; i < 10; i++ {

		vl[i].Id = v[i].VideoId
		vl[i].Author.Name = v[i].Author
		vl[i].Title = v[i].Title
		vl[i].PlayUrl = v[i].PlayUrl
		vl[i].FavoriteCount = v[i].FavoriteCount
		vl[i].CommentCount = v[i].CommentCount
		vl[i].IsFavorite = v[i].IsFavorite
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0},
		//VideoList: DemoVideos,
		VideoList: vl,
		NextTime:  time.Now().Unix(),
	})
}
