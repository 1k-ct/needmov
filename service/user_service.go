package user

import (
	"needmov/db"
	"needmov/entity"

	"github.com/gin-gonic/gin"
)

type Service struct{}

type ChannelInfo entity.ChannelInfos

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
