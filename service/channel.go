package service

import (
	"github.com/Vzttfromxduszu/golang-1.git/common/global"
	models "github.com/Vzttfromxduszu/golang-1.git/model"
)

type ChannelService struct {
}

// 添加
func (cs *ChannelService) AddChannel(channel models.Channel) int64 {
	return global.Db.Table("channel").Create(&channel).RowsAffected // RowsAffected：返回插入操作影响的行数，用于确认插入是否成功。
}

// 更新
func (cs *ChannelService) UpdateChannel(channel models.Channel) int64 {
	return global.Db.Where("id = ?", channel.Id).Updates(&channel).RowsAffected
}

// 删除
func (cs *ChannelService) DeleteChannel(id int) int64 {
	return global.Db.Delete(&models.Channel{}, id).RowsAffected
}

// 查询
func (cs *ChannelService) GetChannel(id int) models.Channel {
	var channel models.Channel
	global.Db.First(&channel, id)
	return channel
}

// 查询所有
func (cs *ChannelService) GetChannelList() []models.Channel {
	var channels []models.Channel
	global.Db.Find(&channels)
	return channels
}
