package controller

import (
	"github.com/binqibang/mini-douyin/business"
	"github.com/binqibang/mini-douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList [10]model.Video `json:"video_list,omitempty"`
	NextTime  int64           `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//返回从数据库中查找最近的一条数据
	videos, _ := business.GetTenVideos()

	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{StatusCode: 0},
		//VideoList: DemoVideos,
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
