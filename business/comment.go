package business

import (
	"github.com/binqibang/mini-douyin/model"
	"strconv"
)

var commentDao model.VideoCommentDao = model.VideoCommentDao{}

func GetCommentsByVideoId(videoId string) ([]model.VideoComment, error) {
	videoIdInt64, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		return nil, err
	}
	return commentDao.QueryByVideo(videoIdInt64)
}
