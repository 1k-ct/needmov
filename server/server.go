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
	r := router(true)
	r.Run()
	appengine.Main()
}

func router(gae bool) *gin.Engine {
	r := gin.Default()
	if !gae {
		r.LoadHTMLGlob("/home/sato/go/src/github.com/1k-ct/nomv/src/needmov/templates/**/*") //*/**
	} else {
		r.Static("/assets", "./assets")
		r.LoadHTMLGlob("templates/**/*")
	}
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

	r.GET("/logout", ctrl.PostLogout)
	r.POST("/regvideo", ctrl.CreateVideoInfo)
	r.POST("/regchannel", ctrl.CreateChannelInfo)
	r.POST("/shiromiyaregvideo", ctrl.ShiromiyaCreateVideoInfo)
	r.POST("/shiromiyaregchannel", ctrl.ShiromiyaCreateChannelInfo)
	r.POST("/hashibaregvideo", ctrl.HashibaCreateVideoInfo)
	r.POST("/hashibaregchannel", ctrl.HashibaCreateChannelInfo)

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
	api := r.Group("/api")
	{
		//"api/ch-info" apiで登録したデータベースを全部取る "api/ch-info"
		api.GET("/ch-info", ctrl.APIAllGetChannelInfo)

		// 選んだ人の、チャンネルを取る "api/ch-sel?who-ch="
		// "api/ch-sel?who-ch=UCxxxxxxxxxxxxxxxxxxxxxx"
		api.GET("/ch-sel", ctrl.APISelectWho)

		// 選んだ人と日付？ "api/date-sel?who-ch=&date="
		// "api/date-sel?who-ch=&date=2020-01-20"
		api.GET("/date-sel", ctrl.APISelectDate)

		// 選んだ人と最新の日付 "api/latest-ch?who-ch="
		// "api/latest-ch?who-ch=2020-02-05"
		api.GET("/latest-ch", ctrl.APISelectLatest)

		// 選んだ人とBETWEEN日付の選択 "api/date-between?who-ch=&a=&b="
		// "api/date-between?who-ch=UCxxxxxxxxxxxxxxxxxxxxxx&a=2020-10-10&b=2020-10-20"
		// 2020-10-10から2020-10-20の間でURLで指定したチャンネル情報です。
		api.GET("/date-between", ctrl.APISelectDateBetween)

		// urlを登録する１つだけ "api/reg?url="
		// "api/reg?url=UCxxxxxxxxxxxxxxxxxxxxxx"
		api.POST("/reg", ctrl.APIInsterChURL)

		// ch情報をjsonで受け取りdbに保存する "api/pri" "POST" bindJSON entity.ChannelInfos = ch
		api.POST("/pri", ctrl.APIInsterChInfo)
		comme := api.Group("comme")
		{
			// コメントデータをdbに保存する。"api/data" "POST" bindJSON entity.Data
			comme.POST("/data", ctrl.APIInsertCommentData)

			// name ? その人(name)が書いたコメント、チャンネル内全て
			// api/comme/namesel?name=xxx
			comme.GET("/name_sel", ctrl.CommeName)

			// name ? and video_id ? その人(name)が書いたコメント、動画内全て
			// api/comme/namecommesel?name=xxx&id=xxx(url)
			comme.GET("/namecomme_sel", ctrl.CommeNameCom)

			// type ? (superChat) そのチャンネルのsuperChat全て
			// api/comme/allsc
			comme.GET("/all_sc", ctrl.CommeAllSC) // もし、他のチャンネルが登録されれば消す

			// type ? (superChat) and video_id(url) その動画内でのsuperChat全て
			// api/comme/videosc?chid=xxx&id=xxx
			comme.GET("video_sc", ctrl.CommeVideoSC)

			// type ? (superChat) and name ? その人(name)のsuperChatチャンネル内全て
			// api/comme/namesc?name=xxx&chid=xxx
			comme.GET("/name_sc", ctrl.CommeNameSC)

			// type ? (superChat) and name ? video_id(url) ?　その人(name)のsuperChat動画内全て
			// api/comme/namesc?name=xxx&id=xxx&chid=xxx
			comme.GET("/namevideo_sc", ctrl.CommeNameVideoSC)

			// message ? (like) そのチャンネルでコメントを検索、全て
			// api/comme/chmsg?chid=xxx&msg=xxx
			comme.GET("/chmsg_simi", ctrl.CommeChMsg)
			// message ? (like) video_id ? その動画内でのコメント検索、全て
			// api/comme/chvimsg?chid=xxx&id=xxx&msg=xxx
			comme.GET("/chvimsg_simi", ctrl.CommeChViMsg)

		}
	}

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
			c.Redirect(http.StatusMovedPermanently, "/")
			c.Abort()
		} else {
			c.Set("ID", LoginInfo.ID)
			c.Next()
		}
		log.Println("ログインチェック終わり")
	}
}

var url string = "https://virtual-youtuber.userlocal.jp/lives"

func startCruise(url string) func() (string, bool) {
	dataLink := GetLivingVideo(url)
	lenDataLink := len(dataLink) // lenDataLink = 動画の本数
	n := -1
	return func() (string, bool) {
		n++
		if n == lenDataLink {
			return dataLink[0], false //errors.New("終了")
		}
		return dataLink[n], true
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
