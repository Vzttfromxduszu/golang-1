package controller

import (
	"fmt"
	"net/http"

	models "github.com/Vzttfromxduszu/golang-1.git/model"
	service "github.com/Vzttfromxduszu/golang-1.git/service"
	gintemplate "github.com/foolin/gin-template"

	// gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	// "github.com/golangci/golangci-lint/pkg/golinters/nilerr"
)

var user service.UserService

func Register(c *gin.Context) {
	// TODO
	var user1 models.User
	err := c.BindJSON(&user1)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	user.Register(user1)
	// c.Redirect(http.StatusFound, "/admin/user/list")

}

func GoRegister(c *gin.Context) {
	gintemplate.HTML(c, http.StatusOK, "user/register", gin.H{})
}

func ListUser(c *gin.Context) {
	users := user.GetUserList()
	gintemplate.HTML(c, http.StatusOK, "user/list", gin.H{"ulist": users})
}
