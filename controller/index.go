package controller

import (
	"net/http"

	"github.com/Vzttfromxduszu/golang-1.git/service"
	gintemplate "github.com/foolin/gin-template"
	gin "github.com/gin-gonic/gin"
)

var indexChannelService service.ChannelService
var indexPostService service.PostService

// 后台首页
func AdminIndex(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "admin/index", gin.H{})
}

// 前台首页
func Index(c *gin.Context) {
	channels := indexChannelService.GetChannelList()
	posts := indexPostService.GetPostList()
	gintemplate.HTML(c, http.StatusOK, "index", gin.H{"clist": channels, "post": posts})
}
