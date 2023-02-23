package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserRelation struct {
	ID        int64          `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	UserID    int64          `gorm:"column:user_id;type:int;not null" json:"user_id"`
	ToUserID  int64          `gorm:"column:to_user_id;type:int;not null" json:"to_user_id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
}

// TableName UserFollow's table name
func (*UserRelation) TableName() string {
	return "user_relation"
}

func (*User) CreateRelation(userRelation *UserRelation) error {
	if err := db.Create(userRelation).Error; err != nil {
		return err
	}
	return nil
}

/*
	func (*User) UpdateRelation(userRelation *UserRelation) error {
		if err := db.Create(userRelation).Error; err != nil {
			return err
		}
		return nil
	}
*/

func (*User) QueryRelation(user_id int64, to_user_id int64) (*UserRelation, error) {
	//存储查询出的关系
	queryRelation := UserRelation{}
	if err := db.
		Where("user_id=?", user_id).
		Where("to_user_id", to_user_id).
		Take(&queryRelation).
		Error; err != nil {
		_ = fmt.Errorf("Relation does not exist")
		return nil, err
	}
	return &queryRelation, nil
}

func (*User) DeleteRelation(userRelation *UserRelation) error {
	if err := db.Delete(userRelation).Error; err != nil {
		return err
	}
	return nil
}

// &取地址
// *根据地址取值
func (*User) QueryFollowingList(userID int64) *[]User {
	userList := make([]User, 0)
	if err := db.
		Model(UserRelation{}).
		Where("user_id = ?", userID).
		Pluck("to_user_id", &userList).Error; err != nil {
		return nil
	}
	return &userList
}

func (*User) QueryFollowerList(userID int64) *[]User {
	userList := make([]User, 0)
	if err := db.
		Model(UserRelation{}).
		Where("to_user_id = ?", userID).
		Pluck("user_id", &userList).Error; err != nil {
		return nil
	}
	return &userList
}
