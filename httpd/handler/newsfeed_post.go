package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github/bcarrol2/GOregon/platform/newsfeed"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post string `json:"post"`
}

func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func (c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)
		
		item := newsfeed.Item {
			Title: requestBody.Title,
			Post: requestBody.Post,
		}
		feed.Add(item)
		
		c.Status(http.StatusNoContent)
		// 	c.JSON(http.StatusOK, map[string]string{
		// 	"hello": "working",
		// })
	}
}