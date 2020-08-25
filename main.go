package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var url string = "https://virtual-youtuber.userlocal.jp/lives"

/*
func main() {
	//connection := "host=db port=5432 user=aaaa password= dbname= sslmode=disable"
	//db, err := gorm.Open("postgres", connection)

	//url := os.Getenv("")
	//connection, err := pq.ParseURL(url)
	//if err != nil {
	//	panic(err.Error())
	//}
	//connection += " sslmode=require"

	//db, err := gorm.Open("postgres", connection)
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer db.Close()

	r := gin.Default()*/
//r.LoadHTMLGlob("templates/*/**")
/*
	r.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "start.html", gin.H{}) })
	//r.GET("/new", ctrl.VideoStart)
	r.GET("/ggnew", func(c *gin.Context) { c.Redirect(302, "/new") })
	r.GET("/stoppoint", func(c *gin.Context) { c.HTML(http.StatusOK, "stoppoint.html", gin.H{}) })
	hashiba := r.Group("/hashiba")
	views := GetChannelName()
	{
		hashiba.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "hashibadeteil.html", gin.H{}) })

		hashiba.GET("/home", func(c *gin.Context) { c.HTML(http.StatusOK, "hashibahome.html", gin.H{"views": views}) })
	}
	r.GET("/shiromiya", func(c *gin.Context) { c.HTML(http.StatusOK, "shiromiyahome.html", gin.H{}) })
	sc := startCruise(url)
	r.GET("/new", func(c *gin.Context) {
		dataLink, ok := sc()
		if ok {
			c.HTML(200, "index.html", gin.H{"dataLink": dataLink})
		} else if !ok {
			sc = startCruise(url)
			c.Redirect(302, "/ggnew")
		}
	})
	r.Run()
}
*/
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
	doc, err := goquery.NewDocument("https://www.youtube.com/channel/UC_BlXOQe5OcRC7o0GX8kp8A/about")
	if err != nil {
		panic(err)
	}
	doc.Find("#right-column").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
	})
}
