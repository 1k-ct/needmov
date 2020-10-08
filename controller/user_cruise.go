package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RedirectGGNew "/ggnew" リダイレクト302 -> "/new"
func (pc Controller) RedirectGGNew(c *gin.Context) {
	c.Redirect(302, "/new")
}

// Stoppoint "/stoppoint"
func (pc Controller) Stoppoint(c *gin.Context) {
	c.HTML(http.StatusOK, "stoppoint.html", gin.H{})
}

/*
var url string = "https://virtual-youtuber.userlocal.jp/lives"

func startCruise(url string) func() (string, bool) {
	dataLink, err := GetLivingVideo(url) //動画をスクレイピングしてくる
	if err != nil {
		log.Fatal("urlが無効")
	}
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
func GetLivingVideo(url string) ([]string, error) {
	var dataLink []string
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	doc.Find("div").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("data-link")
		if len(url) > 10 {
			dataLink = append(dataLink, url)
		}
	})
	return dataLink, nil
}
*/
