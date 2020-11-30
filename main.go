package main

import (
	"context"
	"log"
	"needmov/db"
	"needmov/entity"
	"needmov/server"
	youtubeapi "needmov/youtubeAPI"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// db.DropDBGorm(&entity.Data2{})
	db.NewMakeDB()

	db := db.ConnectGorm()
	defer db.Close()

	db.AutoMigrate(&entity.Data3{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go periodicLoop(ctx, 8*time.Hour) //24*time.Hour
	server.Init()
	// s := "🥺" // encoding Base64 => 8J+lug==
}

func jsonCreateChannel(t time.Time) {
	for _, url := range db.AllGetRegCh() {
		channelID, channelName, viewCount, subscriberCount, videoCount, err := youtubeapi.PrintChannelInfo(url.ChannelID)
		if err != nil {
			log.Println(err) //youtube-API_KEYが何らかのエラーのときの処理
			continue
		}
		db.InsertChannelInfo(channelID, channelName, viewCount, subscriberCount, videoCount)
	}
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
