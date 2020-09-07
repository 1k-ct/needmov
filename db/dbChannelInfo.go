package db

import (
	"needmov/entity"
)

// InsertChannelInfo channelInfoに追加
func InsertChannelInfo(
	//id uint64,
	channelID string,
	channelName string,
	viewCount uint64,
	subscriberCount uint64,
	videoCount uint64,
) {
	db := ConnectGorm()
	db.Create(&entity.ChannelInfos{
		//ID:              id,
		ChannelID:       channelID,
		ChannelName:     channelName,
		ViewCount:       viewCount,
		SubscriberCount: subscriberCount,
		VideoCount:      videoCount,
	})
	defer db.Close()
}

// GetDBChannelInfo channelInfo の DB 全て取得
func GetDBChannelInfo() []entity.ChannelInfos {
	db := ConnectGorm()
	var channelInfo []entity.ChannelInfos
	db.Find(&channelInfo)
	db.Close()
	return channelInfo
}

// DeleteDBChannelInfo 選択したidをchannelInfo DB から削除
func DeleteDBChannelInfo(id int) {
	db := ConnectGorm()
	var channelInfo entity.ChannelInfos
	db.Find(&channelInfo, id)
	db.Delete(&channelInfo)
	db.Close()
}
