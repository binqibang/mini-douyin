package business

import (
	"github.com/binqibang/mini-douyin/config"
	"github.com/binqibang/mini-douyin/model"
	"sort"
)

var videoDao model.VideoDao

// 这个函数用来获得十个最近更新的视频
// 首先，获取数据库中的所有视频，然后按照时间排序，取前10个视频
func GetTenVideos() ([10]model.Video, error) {
	conf, err := config.LoadConfig("C:/Users/13099/GolandProjects/mini-douyin/config/settings_dev.yml")
	ip := conf.App.Address
	videos, err := videoDao.QueryVideoList()
	sort.SliceStable(videos, func(i, j int) bool {
		return videos[i].PublishTime.Unix() < videos[j].PublishTime.Unix()
	})
	//把前10条复制到videolist
	vl := [10]model.Video{}
	var v []model.Video
	if len(videos) > 10 {
		v = videos[:9]
	} else {
		v = videos
	}
	for i := 0; i < len(v); i++ {

		vl[i].VideoId = v[i].VideoId
		vl[i].Author = v[i].Author
		vl[i].Title = v[i].Title
		vl[i].CoverUrl = ip + "/douyin/feed_photo/?path=" + v[i].CoverUrl
		vl[i].PlayUrl = ip + "/douyin/feed_video/?path=" + v[i].PlayUrl
		vl[i].FavoriteCount = v[i].FavoriteCount
		vl[i].CommentCount = v[i].CommentCount
		vl[i].IsFavorite = v[i].IsFavorite
		vl[i].PublishTime = v[i].PublishTime
	}

	return vl, err
}
