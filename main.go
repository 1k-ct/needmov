package main

import (
	"bytes"
	"context"
	"log"
	"needmov/db"
	"needmov/entity"
	"needmov/server"
	youtubeapi "needmov/youtubeAPI"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/binding"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	/*
		id, name, viewCount, subscriberCount, videoCount, err := youtubeapi.PrintChannelInfo("UF4OLuTg") //UCL-2thbJ7grC9fmGF4OLuTg
		if err != nil {
			log.Println(err)
		}
		fmt.Println(id, name, viewCount, subscriberCount, videoCount)
	*/

	db.DropDBGorm(&entity.ChannelInfos{})
	db.NewMakeDB()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go periodicLoop(ctx, 3*time.Hour) //24*time.Hour
	server.Init()

}
func HttpPost(url string) error {
	jsonStr := ``

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", binding.MIMEJSON)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
func jsonCreateChannel(t time.Time) {

	// channelInfos = entity.ChannelInfos{}
	// dbd.NewRecord(channelInfos)
	// if dbd.NewRecord(channelInfos) == false {
	// 	c.JSON(http.StatusOK, channelInfos)
	// }
	//if err := c.BindJSON(&channelInfos); err != nil {
	//	c.String(http.StatusBadRequest, "failed err:"+err.Error())
	//}

	for _, url := range db.AllGetRegCh() {
		channelID, channelName, viewCount, subscriberCount, videoCount, err := youtubeapi.PrintChannelInfo(url.ChannelID)
		if err != nil {
			log.Println(err) //youtube-APIがエラーのときの処理
			continue
		}
		db.InsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
	}
	//c.Redirect(http.StatusFound, "/")
}
func periodicLoop(ctx context.Context, interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	date := time.Now()
	jsonCreateChannel(date)
	log.Println("go func api json create ch-info------")
	for {
		select {
		case <-ctx.Done():
			return
		case t := <-ticker.C:
			jsonCreateChannel(t)
		}
	}
}
