package db

import (
	"needmov/entity"
)

// InsertChannelInfo channelInfoに追加
func InsertChannelInfo(
	channelID string,
	channelName string,
	viewCount uint64,
	subscriberCount uint64,
	videoCount uint64,
) {
	db := ConnectGorm()
	db.Create(&entity.ChannelInfos{
		ChannelID:       channelID,
		ChannelName:     channelName,
		ViewCount:       viewCount,
		SubscriberCount: subscriberCount,
		VideoCount:      videoCount,
	})
	defer db.Close()
}

// InsertChannelInfoSub channelInfo お試しfn
func InsertChannelInfoSub() {
	db := ConnectGorm()
	defer db.Close()
	db.Create(&entity.ChannelInfos{})
}

// GetDBChannelInfo channelInfo の DB 全て取得
func GetDBChannelInfo() ([]entity.ChannelInfos, error) {
	db := ConnectGorm()
	defer db.Close()
	var channelInfo []entity.ChannelInfos
	err := db.Find(&channelInfo).Error
	if err != nil {
		return nil, err
	}
	return channelInfo, nil
}

// DeleteDBChannelInfo 選択したidをchannelInfo DB から削除
func DeleteDBChannelInfo(id int) {
	db := ConnectGorm()
	var channelInfo entity.ChannelInfos
	db.Find(&channelInfo, id)
	db.Delete(&channelInfo)
	db.Close()
}
