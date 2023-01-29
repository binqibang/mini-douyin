package main

import (
	"github.com/binqibang/mini-douyin/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	err := r.Run()
	if err != nil {
		log.Fatalf("can not start webserver, %s", err)
	}
}
