package main

import (
	"context"
	"log"
	"needmov/db"
	"needmov/server"
	youtubeapi "needmov/youtubeAPI"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	//id, name, description, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo("U1xQX6BllRA")
	//fmt.Println(id, name, description, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)

	db.NewMakeDB()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go periodicLoop(ctx, 60*time.Second)
	server.Init()
}

func jsonCreateChannel(t time.Time, channelURL ...string) {
	//var channelInfos []entity.ChannelInfos
	//var c *gin.Context
	//var dbd *gorm.DB
	//dbd.NewRecord(channelInfos)
	//if dbd.NewRecord(channelInfos) == false {
	//	c.JSON(http.StatusOK, channelInfos)
	//}
	//if err := c.BindJSON(&channelInfos); err != nil {
	//	c.String(http.StatusBadRequest, "failed err:"+err.Error())
	//}
	for _, url := range channelURL {
		channelID, channelName, viewCount, subscriberCount, videoCount := youtubeapi.PrintChannelInfo(url)
		db.InsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
	}
	//c.Redirect(http.StatusFound, "/")
}
func periodicLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	date := time.Now()
	channelURL := []string{"UCL-2thbJ7grC9fmGF4OLuTg", "UCvUc0m317LWTTPZoBQV479A"}
	jsonCreateChannel(date, channelURL...)
	log.Println("go func api json create ch-info------")
	for {
		select {
		case <-ctx.Done():
			return
		case t := <-ticker.C:
			jsonCreateChannel(t, channelURL...)
		}
	}
}
