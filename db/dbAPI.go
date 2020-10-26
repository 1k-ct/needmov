package db

import (
	"needmov/entity"
	"time"
)

// InsertVideoInfo videoInfo db に、情報を書き込み
func InsertVideoInfo(
	//id int,
	videoID string,
	videoName string,
	//videoDescription string,
	thumbnailURL string,
	viewCount uint64,
	commentCount uint64,
	likeCount uint64,
	dislikeCount uint64,
	uploadDate time.Time,
) {
	db := ConnectGorm()
	db.Create(&entity.VideoInfos{
		//ID:               id,
		VideoID:   videoID,
		VideoName: videoName,
		//VideoDescription: videoDescription,
		ThumbnailURL: thumbnailURL,
		ViewCount:    viewCount,
		CommentCount: commentCount,
		LikeCount:    likeCount,
		DislikeCount: dislikeCount,
		UploadDate:   uploadDate,
	})
	defer db.Close()
}

// GetDBVideoInfo databaseから、Video_Infoの情報を取得
func GetDBVideoInfo() ([]entity.VideoInfos, error) {
	db := ConnectGorm()
	var videoInfo []entity.VideoInfos
	if err := db.Find(&videoInfo).Error; err != nil {
		return nil, err
	}
	db.Close()
	return videoInfo, nil
}

/*
// AllDBGetVideoInfo 選択したDB(videoInfo)の情報をすべて取得する
//"HashibaVideoInfos","ShiromiyaVideoInfos","VideoInfos"
func AllDBGetVideoInfo(videoInfo string) (interface{}, error) {
	db := ConnectGorm()
	defer db.Close()

	switch videoInfo {
	case "HashibaVideoInfos":
		var channelInfo []entity.HashibaVideoInfos
		db.Find(&channelInfo)
		return channelInfo, nil
	case "ShiromiyaVideoInfos":
		var channelInfo []entity.ShiromiyaVideoInfos
		db.Find(&channelInfo)
		return channelInfo, nil
	case "VideoInfos":
		var videoInfo []entity.VideoInfos
		db.Find(&videoInfo)
		return videoInfo, nil
	default:
		return nil, errors.New("そのdb_nameありません")
	}
}
*/
// DeleteDBVideoInfo 選択したidをVideoInfo DB　から削除
func DeleteDBVideoInfo(id int) {
	db := ConnectGorm()
	var videoInfo entity.VideoInfos
	db.Find(&videoInfo, id)
	db.Delete(&videoInfo)
	db.Close()
}

/*
// ShiromiyaInsertVideoInfo 白宮の動画情報をDBに登録
func ShiromiyaInsertVideoInfo(
	//id int,
	videoID string,
	videoName string,
	//videoDescription string,
	thumbnailURL string,
	viewCount uint64,
	commentCount uint64,
	likeCount uint64,
	dislikeCount uint64,
	uploadDate time.Time,
) {
	db := ConnectGorm()
	defer db.Close()

	db.Create(&entity.ShiromiyaVideoInfos{
		//ID:               id,
		VideoID:   videoID,
		VideoName: videoName,
		//VideoDescription: videoDescription,
		ThumbnailURL: thumbnailURL,
		ViewCount:    viewCount,
		CommentCount: commentCount,
		LikeCount:    likeCount,
		DislikeCount: dislikeCount,
		UploadDate:   uploadDate,
	})
}

// HashibaInsertVideoInfo 羽柴の動画情報をDBに登録
func HashibaInsertVideoInfo(
	//id int,
	videoID string,
	videoName string,
	//videoDescription string,
	thumbnailURL string,
	viewCount uint64,
	commentCount uint64,
	likeCount uint64,
	dislikeCount uint64,
	uploadDate time.Time,
) {
	db := ConnectGorm()
	defer db.Close()

	db.Create(&entity.HashibaVideoInfos{
		//ID:               id,
		VideoID:   videoID,
		VideoName: videoName,
		//VideoDescription: videoDescription,
		ThumbnailURL: thumbnailURL,
		ViewCount:    viewCount,
		CommentCount: commentCount,
		LikeCount:    likeCount,
		DislikeCount: dislikeCount,
		UploadDate:   uploadDate,
	})
}
*/
