package main

import (
	"log"
	"needmov/db"
	"needmov/server"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var url string = "https://virtual-youtuber.userlocal.jp/lives"

func startCruise(url string) func() (string, bool) {
	dataLink := GetLivingVideo(url) //動画をスクレイピングしてくる
	log.Println("スクレイピング出来たよ！")
	lenDataLink := len(dataLink) // 動画の本数
	//fmt.Println(lenDataLink)
	n := -1
	return func() (string, bool) {
		n++
		if n == lenDataLink {
			return dataLink[0], false //errors.New("終了")
		}
		return dataLink[n], true //, "mada"
	}
}

// GetLivingVideo 指定されたLIVE配信中の動画のURLを取得する -> return slice
func GetLivingVideo(url string) []string {
	var dataLink []string
	doc, _ := goquery.NewDocument(url)
	doc.Find("div").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("data-link")
		if len(url) > 10 {
			dataLink = append(dataLink, url)
		}
	})
	return dataLink
}

// GetChannelName 羽柴チャンネルの視聴回数をスクレイピングする
func GetChannelName() string {
	doc, err := goquery.NewDocument("https://www.youtube.com/channel/UC_BlXOQe5OcRC7o0GX8kp8A/about")
	if err != nil {
		panic(err)
	}
	selection := doc.Find("#right-column > yt-formatted-string:nth-child(3)")
	innerSelection := selection.Text()

	return innerSelection
}

func main() {
	//db.NewMakeDB()
	//server.Init()
	//id, name, viewCount, subscriberCount, videoCount := youtubeapi.PrintChannelInfo("UC_BlXOQe5OcRC7o0GX8kp8A")
	//fmt.Println(id, name, viewCount, subscriberCount, videoCount)

	//id, name, description, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate := youtubeapi.PrintVideoInfo("pV1GoNi5mFs")
	//fmt.Println(id, name, description, thumbnailURL, viewCount, commentCount, likeCount, dislikeCount, uploadDate)
	db.NewMakeDB()
	server.Init()
}
