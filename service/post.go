package service

import (
	"github.com/Vzttfromxduszu/golang-1.git/common/global"
	models "github.com/Vzttfromxduszu/golang-1.git/model"
)

type PostService struct {
}

// 添加
func (ps *PostService) AddPost(post models.Post) int64 {
	return global.Db.Table("post").Create(&post).RowsAffected // RowsAffected：返回插入操作影响的行数，用于确认插入是否成功。
}

// 更新
func (ps *PostService) UpdatePost(post models.Post) int64 {
	return global.Db.Updates(&post).RowsAffected
}

// 删除
func (ps *PostService) DeletePost(id int) int64 {
	return global.Db.Delete(&models.Post{}, id).RowsAffected
}

// 查询
func (ps *PostService) GetPost(id int) models.Post {
	var post models.Post
	global.Db.First(&post, id)
	return post
}

// 查询所有
func (ps *PostService) GetPostList() []models.Post {
	var posts []models.Post
	global.Db.Find(&posts)
	return posts
}
