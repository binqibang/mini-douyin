package model

import (
	"gorm.io/gorm"
	"time"
)

// Video mapped from table <video>
type Video struct {
	VideoID        int64          `gorm:"column:video_id;type:int;primaryKey;autoIncrement:true" json:"video_id"`
	Title          string         `gorm:"column:title;type:varchar(255)" json:"title"`
	UserID         int64          `gorm:"column:user_id;type:int" json:"user_id"`
	PlayURL        string         `gorm:"column:play_url;type:varchar(255)" json:"play_url"`
	CoverURL       string         `gorm:"column:cover_url;type:varchar(255)" json:"cover_url"`
	IsFavourite    int64          `gorm:"column:is_favourite;type:int" json:"is_favourite"`
	FavouriteCount int64          `gorm:"column:favourite_count;type:int" json:"favourite_count"`
	CommentCount   int64          `gorm:"column:comment_count;type:int" json:"comment_count"`
	CreatedAt      time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
}

// TableName Video's table name
func (*Video) TableName() string {
	return "video"
}
