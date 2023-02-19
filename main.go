package main

//整个项目的入口，gin框架的初始化和启动
import (
	"github.com/binqibang/mini-douyin/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//开启服务，go关键字实现并发
	go service.RunMessageServer()

	//创建默认的路由引擎对象
	r := gin.Default()

	//初始化路由
	initRouter(r)

	err := r.Run()
	if err != nil {
		log.Fatalf("can not start webserver, %s", err)
	}
}
