package models

import "gorm.io/gorm"

// 数据库，博客数据映射模型
type User struct {
	gorm.Model
	Username string `gorm:"username" json:"username" html:"username" form:"username"`
	Password string `gorm:"password" json:"password" html:"password" form:"password"`
}

// 指定表名为user
func (User) TableName() string {
	return "users"
}
