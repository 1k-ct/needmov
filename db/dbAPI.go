package db

import (
	"needmov/entity"
	"time"
)

// InsertVideoInfo videoInfo db に、情報を書き込み
func InsertVideoInfo(
	videoID string,
	videoName string,
	thumbnailURL string,
	viewCount uint64,
	commentCount uint64,
	likeCount uint64,
	dislikeCount uint64,
	uploadDate time.Time,
) {
	db := ConnectGorm()
	db.Create(&entity.VideoInfos{
		VideoID:      videoID,
		VideoName:    videoName,
		ThumbnailURL: thumbnailURL,
		ViewCount:    viewCount,
		CommentCount: commentCount,
		LikeCount:    likeCount,
		DislikeCount: dislikeCount,
		UploadDate:   uploadDate,
	})
	defer db.Close()
}

// CreateDBSelfFn DB Insert create 自作関数
// var d Data
// db := db.ConnectGorm()
// db.Create(&d) 同じこと
func CreateDBSelfFn(d interface{}) error {
	db := ConnectGorm()
	defer db.Close()
	err := db.Create(d).Error
	if err != nil {
		return err
	}
	return nil
}

// GetDBVideoInfo databaseから、Video_Infoの情報を取得
func GetDBVideoInfo() ([]entity.VideoInfos, error) {
	db := ConnectGorm()
	var videoInfo []entity.VideoInfos
	err := db.Find(&videoInfo).Error
	if err != nil {
		return nil, err
	}
	db.Close()
	return videoInfo, nil
}

// DeleteDBVideoInfo 選択したidをVideoInfo DB　から削除
func DeleteDBVideoInfo(id int) {
	db := ConnectGorm()
	var videoInfo entity.VideoInfos
	db.Find(&videoInfo, id)
	db.Delete(&videoInfo)
	db.Close()
}
