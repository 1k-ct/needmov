package user

import (
	"log"
	apierrors "needmov/APIerrors"
	"needmov/db"
	youtubeapi "needmov/youtubeAPI"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateVideoInfo "/regVideo" PrintVideoInfoをdbに登録
func (pc Controller) CreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("videoURL")

	videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo(videoURL)
	db.InsertVideoInfo(videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)
	c.Redirect(http.StatusFound, "/")
}

// CreateChannelInfo "/regchannel" PrintChannelInfoをdbに登録
func (pc Controller) CreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("channelURL")
	channelID, channelName, viewCount, subscriberCount, videoCount, err := youtubeapi.PrintChannelInfo(channelURL)
	if err != nil {
		c.AbortWithStatusJSON(404, apierrors.ErrDB)

	} else {
		db.InsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
		c.Redirect(http.StatusFound, "/")
	}
}

//ShiromiyaCreateVideoInfo "shiromiyaregvideo" PostForm => "shiromiyaVideoURL"白宮の動画情報をDB(ShiromiyaVideoInfo)に登録
func (pc Controller) ShiromiyaCreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("shiromiyaVideoURL")

	videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo(videoURL)
	db.InsertVideoInfo(videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)
	c.Redirect(http.StatusFound, "/")
}

// ShiromiyaCreateChannelInfo "shiromiyaregchannel" PostForm => "shiromiyaChannelURL"白宮のチャンネル情報をDB(ShiromiyaChannelInfo)に登録
func (pc Controller) ShiromiyaCreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("shiromiyaChannelURL")
	channelID, channelName, viewCount, subscriberCount, videoCount, err := youtubeapi.PrintChannelInfo(channelURL)
	if err != nil {
		c.String(http.StatusOK, "%v\n", err)
		log.Println(err)
	} else {
		db.InsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
		c.Redirect(http.StatusFound, "/")
	}
}

// HashibaCreateVideoInfo "hashibaregvideo" PostForm => "HashibaVideoURL"羽柴の動画情報をDB(HashibaVideoInfo)に登録
func (pc Controller) HashibaCreateVideoInfo(c *gin.Context) {
	videoURL := c.PostForm("hashibaVideoURL")

	videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo(videoURL)
	db.InsertVideoInfo(videoID, videoName, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)
	c.Redirect(http.StatusFound, "/")
}

//HashibaCreateChannelInfo "hashibaregchannel" PostForm => "HashibaChannelURL"羽柴のチャンネル情報をDB(HashibaChannelInfo)に登録
func (pc Controller) HashibaCreateChannelInfo(c *gin.Context) {
	channelURL := c.PostForm("hashibaChannelURL")
	channelID, channelName, viewCount, subscriberCount, videoCount, err := youtubeapi.PrintChannelInfo(channelURL)
	if err != nil {
		log.Printf("%v\n", err)
		c.String(http.StatusOK, "%v\n", err)
	} else {
		db.InsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
		c.Redirect(http.StatusFound, "/")
	}
}
