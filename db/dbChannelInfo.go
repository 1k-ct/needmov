package db

import (
	"errors"
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

// AllGetDBChannelInfo 選択したDB(channelInfo)の情報をすべて取得する
//"ShiromiyaChannelInfos","HashibaChannelInfos","ChannelInfos"
func AllGetDBChannelInfo(chInfo string) (interface{}, error) {
	db := ConnectGorm()
	defer db.Close()

	switch chInfo {
	case "ShiromiyaChannelInfos":
		var channelInfo []entity.ShiromiyaChannelInfos
		db.Find(&channelInfo)
		return channelInfo, nil
	case "HashibaChannelInfos":
		var channelInfo []entity.HashibaChannelInfos
		db.Find(&channelInfo)
		return channelInfo, nil
	case "ChannelInfos":
		var videoInfo []entity.VideoInfos
		db.Find(&videoInfo)
		return videoInfo, nil
	default:
		return nil, errors.New("そのdb_nameありません")
	}
	//f := func(){}
	//db.Find(&channelInfo)
	//db.Close()
	//return channelInfo
}

// DeleteDBChannelInfo 選択したidをchannelInfo DB から削除
func DeleteDBChannelInfo(id int) {
	db := ConnectGorm()
	var channelInfo entity.ChannelInfos
	db.Find(&channelInfo, id)
	db.Delete(&channelInfo)
	db.Close()
}
func ShiromiyaInsertChannelInfo(
	//id uint64,
	channelID string,
	channelName string,
	viewCount uint64,
	subscriberCount uint64,
	videoCount uint64,
) {
	db := ConnectGorm()
	defer db.Close()

	db.Create(&entity.ShiromiyaChannelInfos{
		//ID:              id,
		ChannelID:       channelID,
		ChannelName:     channelName,
		ViewCount:       viewCount,
		SubscriberCount: subscriberCount,
		VideoCount:      videoCount,
	})
}
func HashibaInsertChannelInfo(
	//id uint64,
	channelID string,
	channelName string,
	viewCount uint64,
	subscriberCount uint64,
	videoCount uint64,
) {
	db := ConnectGorm()
	defer db.Close()

	db.Create(&entity.HashibaChannelInfos{
		//ID:              id,
		ChannelID:       channelID,
		ChannelName:     channelName,
		ViewCount:       viewCount,
		SubscriberCount: subscriberCount,
		VideoCount:      videoCount,
	})
}
