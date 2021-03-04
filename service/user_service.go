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

//GetSomeoneChannelInfo entity.ChannelInfos(db)から指定したURLの情報を全部取る
func (s Service) GetSomeoneChannelInfo(c *gin.Context, url string) ([]ChannelInfo, error) {
	db := db.ConnectGorm()
	defer db.Close()
	var ch []ChannelInfo
	err := db.Where("channel_id = ?", url).Find(&ch).Error
	if err != nil {
		return nil, err
	}
	return ch, nil
}

//GetSomeoneVideoInfo entity.VideoInfos(db)から指定したvideoidの情報を全部取る
func (s Service) GetSomeoneVideoInfo(c *gin.Context, videoID string) ([]VideoInfos, error) {
	db := db.ConnectGorm()
	defer db.Close()
	var vi []VideoInfos
	err := db.Where("video_id = ?", videoID).Find(&vi).Error
	if err != nil {
		return nil, err
	}
	return vi, nil
}
