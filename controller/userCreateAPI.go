package user

import (
	"needmov/db"
	youtubeapi "needmov/youtubeAPI"
	"net/http"

	"github.com/gin-gonic/gin"
)

func viSome(c *gin.Context, videoURL string) {
	viInfo, err := youtubeapi.PrintVideoInfo(videoURL)
	if err != nil {
		c.Redirect(http.StatusBadRequest, "/")
		return
	}
	db.InsertVideoInfo(viInfo)
	c.Redirect(http.StatusFound, "/")
}
func chSome(c *gin.Context, channelURL string) {
	chInfo, err := youtubeapi.PrintChannelInfo(channelURL)
	if err != nil {
		c.Redirect(http.StatusBadRequest, "/")
		return
	}
	db.InsertChannelInfo(chInfo)
	c.Redirect(http.StatusFound, "/")
}

// CreateVideoInfo "/regVideo" PrintVideoInfoをdbに登録
func (pc Controller) CreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("videoURL")
	viSome(c, videoURL)
}

// CreateChannelInfo "/regchannel" PrintChannelInfoをdbに登録
func (pc Controller) CreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("channelURL")
	chSome(c, channelURL)
}

//ShiromiyaCreateVideoInfo "shiromiyaregvideo" PostForm => "shiromiyaVideoURL"白宮の動画情報をDB(ShiromiyaVideoInfo)に登録
func (pc Controller) ShiromiyaCreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("shiromiyaVideoURL")
	viSome(c, videoURL)
}

// ShiromiyaCreateChannelInfo "shiromiyaregchannel" PostForm => "shiromiyaChannelURL"白宮のチャンネル情報をDB(ShiromiyaChannelInfo)に登録
func (pc Controller) ShiromiyaCreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("shiromiyaChannelURL")
	chSome(c, channelURL)
}

// HashibaCreateVideoInfo "hashibaregvideo" PostForm => "HashibaVideoURL"羽柴の動画情報をDB(HashibaVideoInfo)に登録
func (pc Controller) HashibaCreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("hashibaVideoURL")
	viSome(c, videoURL)
}

//HashibaCreateChannelInfo "hashibaregchannel" PostForm => "HashibaChannelURL"羽柴のチャンネル情報をDB(HashibaChannelInfo)に登録
func (pc Controller) HashibaCreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("hashibaChannelURL")
	chSome(c, channelURL)
}
