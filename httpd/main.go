package main

import (
	// "fmt"
	"github/bcarrol2/GOregon/httpd/handler"
	"github.com/gin-gonic/gin"
	"github/bcarrol2/GOregon/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))
	r.Run()
	// feed := newsfeed.New()
	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"Hello", "How are things?"})
	// fmt.Println(feed)
}