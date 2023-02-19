package main

//存放所有路由，通过路由跳转到controller层
import (
	"github.com/binqibang/mini-douyin/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	//视频Feed流，获取视频列表
	apiRouter.GET("/feed/", controller.Feed)
	//feed_vedio这条路由多余？？？
	apiRouter.GET("/feed_vedio/", controller.FeedVedio)
	//用户登录、注册、个人基本信息
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	//用户发布视频投稿
	apiRouter.POST("/publish/action/", controller.Publish)
	// 作品列表
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	//喜欢视频，点赞按钮
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//喜欢列表
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//用户评论
	apiRouter.POST("/comment/action/", controller.CommentAction)
	//评论列表
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	//用户关系
	apiRouter.POST("/relation/action/", controller.RelationAction)
	//用户关注列表
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	//用户好友列表
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	//用户聊天、发送消息
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
