package server

import (
	"log"
	user "needmov/controller"
	"needmov/entity"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

// Init is server run
func Init() {
	r := router()
	r.Run()
	appengine.Main()
}

func router() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/**/*") //*/**

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	ctrl := user.Controller{}

	r.GET("/", ctrl.Start)

	r.GET("/ping", ctrl.Connection)

	u := r.Group("/admin")
	u.Use(sessionCheck())
	{
		u.GET("/", ctrl.Adimn)
	}

	r.GET("/signup", ctrl.SignUpGet)
	r.POST("/signup", ctrl.SignUpPost)
	r.GET("/login", ctrl.LoginGet)
	r.POST("/login", ctrl.LoginPost)

	hashiba := r.Group("/hashiba")
	{
		hashiba.GET("/", ctrl.HashibaHome)
		hashiba.GET("/reg", ctrl.HashibaDeteil)
	}

	shiromiya := r.Group("/shiromiya")
	{
		shiromiya.GET("/", ctrl.ShiromiyaHome)
		shiromiya.GET("/reg", ctrl.ShiromiyaRegVideo)
	}

	r.GET("/logout", ctrl.PostLogout) //r.POST("/logout", ctrl.PostLogout)
	r.POST("/regvideo", ctrl.CreateVideoInfo)
	r.POST("/regchannel", ctrl.CreateChannelInfo)

	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "start.html", gin.H{})
	//})
	//r.GET("/new", ctrl.VideoStart)
	r.GET("/ggnew", ctrl.RedirectGGNew)
	r.GET("/stoppoint", ctrl.Stoppoint)

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

	return r
}

// LoginInfo cookie 関係
var LoginInfo entity.SessionInfo

func sessionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		LoginInfo.ID = session.Get("ID")

		// セッションがない場合、ログインフォームをだす
		if LoginInfo.ID == nil {
			log.Println("ログインしていません")
			c.Redirect(http.StatusMovedPermanently, "/") // /signup
			c.Abort()                                    // これがないと続けて処理されてしまう
		} else {
			c.Set("ID", LoginInfo.ID) // ユーザidをセット
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}

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
