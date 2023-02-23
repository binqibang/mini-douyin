package model

import (
	"gorm.io/gorm"
	"log"
	"time"
)

// Video mapped from table <video>
type Video struct {
	//原本代码
	VideoId       int64     `gorm:"column:video_id;type:int;primaryKey;autoIncrement:true" json:"id,omitempty"`
	UserId        int64     `gorm:"column:user_id;type:int;primaryKey;autoIncrement:true" json:"user_id,omitempty"`
	Author        string    `gorm:"column:author;type:varchar(255);" json:"author"`
	Title         string    `gorm:"column:title;type:varchar(255);" json:"title"`
	PlayUrl       string    `gorm:"column:play_url;type:varchar(255);" json:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url;type:varchar(255);" json:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count;type:int;" json:"favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count;type:int;" json:"comment_count"`
	IsFavorite    bool      `gorm:"column:is_favorite;type:int;" json:"is_favorite"`
	PublishTime   time.Time `gorm:"column:created_at;type:datetime(3);" json:"created_at,omitempty"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return "video"
}

type Videos []Video

func (*Videos) TableName() string {
	return "video"
}

type VideoDao struct {
}

// 这个函数用来获取所有的video
func (*VideoDao) QueryVideoList() ([]Video, error) {

	videos := []Video{}
	err := db.Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (*VideoDao) QueryVideoById(id int64) (Video, error) {
	var video Video
	err := db.Where("user_id = ?", id).Find(&video).Error
	log.Print(id)
	log.Print(video)
	if err == gorm.ErrRecordNotFound {
		return video, nil
	}
	if err != nil {
		return video, err
	}
	return video, nil
}
