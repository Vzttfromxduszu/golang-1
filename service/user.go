package service

import (
	global "github.com/Vzttfromxduszu/golang-1.git/common/global"
	models "github.com/Vzttfromxduszu/golang-1.git/model"
)

type UserService struct {
}

// 新增用户
func (u *UserService) Register(user models.User) int64 {
	return global.Db.Table("users").Create(&user).RowsAffected
}

func (u *UserService) GetUserList() []models.User {
	var users []models.User
	global.Db.Find(&users)
	return users
}
