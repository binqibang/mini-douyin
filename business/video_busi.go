package business

import (
	"github.com/binqibang/mini-douyin/model"
	"sort"
)

var videoDao model.VideoDao

func GetTenVideos() ([10]model.Video, error) {
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

		vl[i].Id = v[i].Id
		vl[i].Author = v[i].Author
		vl[i].Title = v[i].Title
		vl[i].CoverUrl = "http://192.168.178.97:8080/douyin/feed_vedio/?path=" + v[i].CoverUrl
		//PlayUrl改为本机IP
		vl[i].PlayUrl = "http://192.168.178.97:8080/douyin/feed_vedio/?path=" + v[i].PlayUrl
		vl[i].FavoriteCount = v[i].FavoriteCount
		vl[i].CommentCount = v[i].CommentCount
		vl[i].IsFavorite = v[i].IsFavorite
		vl[i].PublishTime = v[i].PublishTime
	}
	return vl, err
}
