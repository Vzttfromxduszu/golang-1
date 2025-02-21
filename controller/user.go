package controller

import (
	"fmt"

	models "github.com/Vzttfromxduszu/golang-1.git/model"
	service "github.com/Vzttfromxduszu/golang-1.git/service"
	"github.com/gin-gonic/gin"
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
