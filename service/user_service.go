package user

import (
	"needmov/db"
	"needmov/entity"

	"github.com/gin-gonic/gin"
)

type Service struct{}

// User struct db Users
type User entity.Users
type ChannelInfo entity.ChannelInfos
type VideoInfos entity.VideoInfos

/*
//GetAllUserInfo entity.Users(db)内容を全部取る
func (s Service) GetAllUserInfo() ([]User, error) {
	db := db.ConnectGorm()
	var user []User
	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
*/
//GetSomeoneChannelInfo entity.ChannelInfos(db)から指定したURLの情報を全部取る
func (s Service) GetSomeoneChannelInfo(c *gin.Context, url string) ([]ChannelInfo, error) {
	db := db.ConnectGorm()
	defer db.Close()
	var channel []ChannelInfo
	if err := db.Where("channel_id = ?", url).Find(&channel).Error; err != nil {
		return nil, err
	}
	return channel, nil
}

//GetSomeoneVideoInfo entity.VideoInfos(db)から指定したvideoidの情報を全部取る
func (s Service) GetSomeoneVideoInfo(c *gin.Context, videoID string) ([]VideoInfos, error) {
	db := db.ConnectGorm()
	defer db.Close()
	var video []VideoInfos
	if err := db.Where("video_id = ?", videoID).Find(&video).Error; err != nil {
		return nil, err
	}
	return video, nil
}
