package test

import (
	"github.com/binqibang/mini-douyin/controller"
	"github.com/binqibang/mini-douyin/model"
	"github.com/stretchr/testify/require"
	"testing"
)

var userDao = model.UserDao{}

func TestCreateUser(t *testing.T) {
	user := controller.User{Username: "tom"}
	require.NoError(t, userDao.CreateUser(&user))
}
