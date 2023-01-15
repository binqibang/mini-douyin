package test

import (
	"github.com/binqibang/mini-douyin/model"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	db      = model.InitDB("../config/settings_dev.yml")
	userDao = model.UserDao{}
)

func TestCreateUser(t *testing.T) {
	user := model.User{Username: "tom"}
	require.NoError(t, userDao.CreateUser(db, &user))
}
