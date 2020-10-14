package user

import (
	"needmov/db"
	youtubeapi "needmov/youtubeAPI"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateVideoInfo "/regVideo" PrintVideoInfoをdbに登録
func (pc Controller) CreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("videoURL")

	videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo(videoURL) //videoDescription
	db.InsertVideoInfo(videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)                    //, videoDescription
	c.Redirect(http.StatusFound, "/")                                                                                                     // http.StatusFound = 302
}

// CreateChannelInfo "/regchannel" PrintChannelInfoをdbに登録
func (pc Controller) CreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("channelURL")
	channelID, channelName, viewCount, subscriberCount, videoCount := youtubeapi.PrintChannelInfo(channelURL)
	db.InsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
	c.Redirect(http.StatusFound, "/") // http.StatusFound = 302
}

//ShiromiyaCreateVideoInfo "shiromiyaregvideo" PostForm => "shiromiyaVideoURL"白宮の動画情報をDB(ShiromiyaVideoInfo)に登録
func (pc Controller) ShiromiyaCreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("shiromiyaVideoURL")

	videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo(videoURL) //videoDescription
	db.ShiromiyaInsertVideoInfo(videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)           //, videoDescription
	c.Redirect(http.StatusFound, "/")                                                                                                     // http.StatusFound = 302
}

// ShiromiyaCreateChannelInfo "shiromiyaregchannel" PostForm => "shiromiyaChannelURL"白宮のチャンネル情報をDB(ShiromiyaChannelInfo)に登録
func (pc Controller) ShiromiyaCreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("shiromiyaChannelURL")
	channelID, channelName, viewCount, subscriberCount, videoCount := youtubeapi.PrintChannelInfo(channelURL)
	db.ShiromiyaInsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
	c.Redirect(http.StatusFound, "/") // http.StatusFound = 302
}

// HashibaCreateVideoInfo "hashibaregvideo" PostForm => "HashibaVideoURL"羽柴の動画情報をDB(HashibaVideoInfo)に登録
func (pc Controller) HashibaCreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("hashibaVideoURL")

	videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo(videoURL) //videoDescription
	db.HashibaInsertVideoInfo(videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)             //, videoDescription
	c.Redirect(http.StatusFound, "/")                                                                                                     // http.StatusFound = 302
}

//HashibaCreateChannelInfo "hashibaregchannel" PostForm => "HashibaChannelURL"羽柴のチャンネル情報をDB(HashibaChannelInfo)に登録
func (pc Controller) HashibaCreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("hashibaChannelURL")
	channelID, channelName, viewCount, subscriberCount, videoCount := youtubeapi.PrintChannelInfo(channelURL)
	db.HashibaInsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
	c.Redirect(http.StatusFound, "/") // http.StatusFound = 302
}
