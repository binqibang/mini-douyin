package controller

import (
	"github.com/gin-gonic/gin"
)

func FeedVideo(c *gin.Context) {
	path := c.Query("path")
	path = "public/" + path
	c.File(path)
}

func FeedPhoto(c *gin.Context) {
	path := c.Query("path")
	path = "public/" + path
	c.File(path)
}
