package controller

import (
	"github.com/gin-gonic/gin"
)

func FeedVedio(c *gin.Context) {
	path := c.Query("path")
	path = "public/" + path
	c.File(path)
}
