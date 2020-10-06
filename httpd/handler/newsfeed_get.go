package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github/bcarrol2/GOregon/platform/newsfeed"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
	return func (c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)
		// 	c.JSON(http.StatusOK, map[string]string{
		// 	"hello": "working",
		// })
	}
}