package model

import (
	"gorm.io/gorm"
	"time"
)

// UserLike mapped from table <user_like>
type UserLike struct {
	ID        int64          `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	UserID    int64          `gorm:"column:user_id;type:int;not null" json:"user_id"`
	VideoID   int64          `gorm:"column:video_id;type:int;not null" json:"video_id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
}

// TableName UserLike's table name
func (*UserLike) TableName() string {
	return "user_like"
}

// 获取ID对应用户的喜欢列表
func (*UserLike) QueryUserLikeById(id int64) ([]UserLike, error) {
	var user_like []UserLike
	err := db.Where("user_id = ?", id).Find(&user_like).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user_like, nil
}
