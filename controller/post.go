package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	models "github.com/Vzttfromxduszu/golang-1.git/model"
	"github.com/Vzttfromxduszu/golang-1.git/service"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

var postService service.PostService
var channelService service.ChannelService

// post list
func ListPost(c *gin.Context) {
	plist := postService.GetPostList()
	gintemplate.HTML(c, http.StatusOK, "post/list", gin.H{"plist": plist})
}

// view post
func ViewPost(c *gin.Context) {
	sid, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(sid)
	post := postService.GetPost(id)
	channels := channelService.GetChannelList()
	gintemplate.HTML(c, http.StatusOK, "post/view", gin.H{"post": post, "channels": channels})
}

// delete post
func DeletePost(c *gin.Context) {
	sid, _ := c.GetQuery("id")
	id, _ := strconv.Atoi(sid)
	postService.DeletePost(id)
	c.Redirect(http.StatusMovedPermanently, "/admin/post/list")
}

// post detail 用来处理前端请求
func PostDetail(c *gin.Context) {
	sid := c.Param("id")
	id, _ := strconv.Atoi(sid)
	post := postService.GetPost(id)

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs // 为了生成标题的锚点
	parser := parser.NewWithExtensions(extensions)                // 创建一个新的 Markdown 解析器

	md := []byte(post.Content)
	md = markdown.NormalizeNewlines(md)      // 规范化换行符
	html := markdown.ToHTML(md, parser, nil) // 将 Markdown 转换为 HTML
	gintemplate.HTML(c, http.StatusOK, "post/detail", gin.H{"post": post, "content": template.HTML(html)})

}

// 上传封面
func UploadThumbnails(c *gin.Context) {
	file, _ := c.FormFile("file")
	extension := filepath.Ext(file.Filename)       // 获取文件的扩展名
	newFileName := uuid.New().String() + extension // 生成一个新的文件名 uuid.New().String() 生成一个新的 UUID
	// 文件已被接收
	pwd, _ := os.Getwd()                                            // 获取当前工作目录
	filepath := fmt.Sprint(pwd, "/assets/thumbnails/", newFileName) // 拼接文件路径

	relativePath := fmt.Sprint("/assets/thumbnails/", newFileName) // 相对路径
	err := c.SaveUploadedFile(file, filepath)                      // 保存文件
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "上传失败"})
		print(err)
		return
	}
	c.String(http.StatusOK, relativePath)

}

// add post or update post
func SavePost(c *gin.Context) {
	var post models.Post
	title := c.PostForm("title")
	content := c.PostForm("content")
	thumbnail := c.PostForm("thumbnail")
	channelId, _ := strconv.Atoi(c.PostForm("channel_id"))

	summary, _ := stripmd.Strip(content) // 从 Markdown 中提取摘要
	l := len(summary)
	if l >= 200 {
		summary = summary[:200]
	} else {
		summary = summary[:l]
	}
	post.Title = title
	post.Content = content
	post.Thumbnail = thumbnail
	post.Summary = summary

	post.AuthorId = 1 //实现注册登录后动态获得
	post.ChannelId = cast.ToInt(channelId)

	if err := c.ShouldBind(&post); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	id := cast.ToInt(c.PostForm("id"))

	if id != 0 {
		postService.UpdatePost(post)
	} else {
		postService.AddPost(post)
	}
	c.Redirect(http.StatusFound, "/admin/post/list")
}
