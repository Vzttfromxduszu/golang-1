package controller

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/Vzttfromxduszu/golang-1.git/model"
	"github.com/Vzttfromxduszu/golang-1.git/service"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

var channel service.ChannelService

// channel list

func ListChannel(c *gin.Context) {

	channels := channel.GetChannelList()
	gintemplate.HTML(c, http.StatusOK, "channel/list", gin.H{"clist": channels})
}

// view channel
func ViewChannel(c *gin.Context) {
	sid, r := c.GetQuery("id")
	var chann models.Channel
	if r {
		id, _ := strconv.Atoi(sid)
		chann = channel.GetChannel(id)
	}
	gintemplate.HTML(c, http.StatusOK, "channel/view", gin.H{"clist": chann})
}

// delete channel
func DeleteChannel(c *gin.Context) {
	sid, r := c.GetQuery("id")
	if r {
		id, _ := strconv.Atoi(sid)
		channel.DeleteChannel(id)
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/channel/list") // 这是一个 HTTP 状态码，表示永久重定向（301 Moved Permanently）。这意味着请求的资源已被永久移动到新的 URL。
}

func GoSaveChannel(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "channel/view", gin.H{})
}

// add channel or update channel
func SaveChannel(c *gin.Context) {
	var chann models.Channel
	err := c.ShouldBind(&chann) // ShouldBind 是一个结合了 c.ShouldBindJSON 和 c.ShouldBindQuery 的方法。它根据请求的 Content-Type 自动选择绑定
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	chann.Status, _ = strconv.Atoi(c.PostForm("status"))
	id, _ := c.GetPostForm("id")

	if id != "0" {
		channel.UpdateChannel(chann)
	} else {
		channel.AddChannel(chann)
	}
	c.Redirect(http.StatusFound, "/admin/channel/list") // 这是一个 HTTP 状态码，表示资源被找到（302 Found）。临时重定向，这意味着请求的资源已被临时移动到新的 URL。

}
