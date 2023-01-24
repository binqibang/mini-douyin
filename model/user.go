package model

import (
	"gorm.io/gorm"
	"time"
)

// User mapped from table <user>
type User struct {
	UserID        int64          `gorm:"column:user_id;type:int;primaryKey;autoIncrement:true" json:"user_id"`
	Username      string         `gorm:"column:username;type:varchar(64)" json:"username"`
	Email         string         `gorm:"column:email;type:varchar(128)" json:"email"`
	Password      string         `gorm:"column:password;type:varchar(128)" json:"password"`
	FollowCount   int64          `gorm:"column:follow_count;type:int" json:"follow_count"`
	FollowerCount int64          `gorm:"column:follower_count;type:int" json:"follower_count"`
	IsFollow      bool           `gorm:"column:is_follow;type:tinyint(1)" json:"is_follow"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return "user"
}

// UserDao define the database crud operations to User
type UserDao struct {
}

func (*UserDao) QueryByUserById(userid int64) (*User, error) {
	var user User
	err := db.Where("user_id = ?", userid).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*UserDao) CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
