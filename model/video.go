package model

import (
	"fmt"
	"time"
)

// Video mapped from table <video>
type Video struct {
	//VideoID        int64          `gorm:"column:video_id;type:int;primaryKey;autoIncrement:true" json:"video_id"`
	//Title          string         `gorm:"column:title;type:varchar(255)" json:"title"`
	//UserID         int64          `gorm:"column:user_id;type:int" json:"user_id"`
	//PlayURL        string         `gorm:"column:play_url;type:varchar(255)" json:"play_url"`
	//CoverURL       string         `gorm:"column:cover_url;type:varchar(255)" json:"cover_url"`
	//IsFavourite    int64          `gorm:"column:is_favourite;type:int" json:"is_favourite"`
	//FavouriteCount int64          `gorm:"column:favourite_count;type:int" json:"favourite_count"`
	//CommentCount   int64          `gorm:"column:comment_count;type:int" json:"comment_count"`
	//CreatedAt      time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	//UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	//DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	Id            int64     `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,omitempty"`
	Author        string    `gorm:"column:author;type:varchar(255);" json:"author"`
	Title         string    `gorm:"column:title;type:varchar(255);" json:"title"`
	PlayUrl       string    `gorm:"column:play_url;type:varchar(255);" json:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url;type:varchar(255);" json:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count;type:int;" json:"favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count;type:int;" json:"comment_count"`
	IsFavorite    bool      `gorm:"column:is_favorite;type:int;" json:"is_favorite"`
	PublishTime   time.Time `gorm:"column:publish_time;type:datetime(3);" json:"publish_time,omitempty"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return "video"
}

type Videos []Video

func (*Videos) TableName() string {
	return "video"
}

//func (s Videos) Len() int {
//	return len(s)
//}
//func (s Videos) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//func (s Videos) Less(i, j int) bool {
//	return s[i].PublishTime.Unix() < s[j].PublishTime.Unix()
//}

type VideoDao struct {
}

func (*VideoDao) QueryVideoList() ([]Video, error) {

	videos := []Video{}
	err := db.Find(&videos).Error
	fmt.Println(videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
