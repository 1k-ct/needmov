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
	videoDescription string,
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
		VideoID:          videoID,
		VideoName:        videoName,
		VideoDescription: videoDescription,
		ThumbnailURL:     thumbnailURL,
		ViewCount:        viewCount,
		CommentCount:     commentCount,
		LikeCount:        likeCount,
		DislikeCount:     dislikeCount,
		UploadDate:       uploadDate,
	})
	defer db.Close()
}

// GetDBVideoInfo databaseから、Video_Infoの情報を取得
func GetDBVideoInfo() []entity.VideoInfos {
	db := ConnectGorm()
	var videoInfo []entity.VideoInfos
	db.Find(&videoInfo)
	db.Close()
	return videoInfo
}

// DeleteDBVideoInfo 選択したidをVideoInfo DB　から削除
func DeleteDBVideoInfo(id int) {
	db := ConnectGorm()
	var videoInfo entity.VideoInfos
	db.Find(&videoInfo, id)
	db.Delete(&videoInfo)
	db.Close()
}
