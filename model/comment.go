package model

import (
	"gorm.io/gorm"
	"time"
)

// VideoComment mapped from table <video_comment>
type VideoComment struct {
	ID        int64          `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	VideoID   int64          `gorm:"column:video_id;type:int;not null" json:"video_id"`
	UserID    int64          `gorm:"column:user_id;type:int;not null" json:"user_id"`
	Comment   string         `gorm:"column:comment;type:text" json:"comment"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
}

// TableName VideoComment's table name
func (*VideoComment) TableName() string {
	return "video_comment"
}
