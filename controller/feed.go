package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"sort"
	"time"
)

type FeedResponse struct {
	Response
	VideoList [10]Video `json:"video_list,omitempty"`
	NextTime  int64     `json:"next_time,omitempty"`
}

type Video_x struct {
	Id            int64     `json:"id,omitempty"`
	Author        string    `json:"author"`
	PlayUrl       string    `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string    `json:"cover_url,omitempty"`
	FavoriteCount int64     `json:"favorite_count,omitempty"`
	CommentCount  int64     `json:"comment_count,omitempty"`
	IsFavorite    bool      `json:"is_favorite,omitempty"`
	PublishTime   time.Time `json:"publish_time,omitempty"`
}
type Video_xes []Video_x

func (s Video_xes) Len() int {
	return len(s)
}
func (s Video_xes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Video_xes) Less(i, j int) bool {
	return s[i].PublishTime.Unix() < s[j].PublishTime.Unix()
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//返回从数据库中查找最近的一条数据
	fmt.Println("test")
	dsn := "root:123123@tcp(localhost:3306)/mini-douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败")
		return
	}

	videos := []Video_x{}
	db.Find(&videos)
	//日期从大到小排列
	sort.Sort(sort.Reverse(Video_xes(videos)))
	//把前10条复制到videolist
	vl := [10]Video{}
	var v []Video_x
	if len(videos) > 10 {
		v = videos[:9]
	} else {
		v = videos
	}
	for i := 0; i < len(v); i++ {
		vl[i].Id = v[i].Id
		vl[i].Author.Name = v[i].Author
		vl[i].CoverUrl = v[i].CoverUrl
		vl[i].PlayUrl = "http://192.168.113.97:8080/douyin/feed_vedio/?path=" + v[i].PlayUrl
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
