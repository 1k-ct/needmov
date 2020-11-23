package server

import (
	user "needmov/controller"
	"net/http"

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

	ctrl := user.Controller{}

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "name_sel.html", nil)
	})
	// r.POST("/", ctrl.GetLikeAMsg)

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
			// api/comme/name_sel?name=xxx
			comme.GET("/name_sel", ctrl.CommeName)

			// name ? and video_id ? その人(name)が書いたコメント、動画内全て
			// api/comme/namecomme_sel?name=xxx&id=xxx(url)
			comme.GET("/namecomme_sel", ctrl.CommeNameCom)

			// type ? (superChat) そのチャンネルのsuperChat全て
			// api/comme/all_sc
			comme.GET("/all_sc", ctrl.CommeAllSC) // もし、他のチャンネルが登録されれば消す

			// type ? (superChat) and video_id(url) その動画内でのsuperChat全て
			// api/comme/video_sc?chid=xxx&id=xxx
			comme.GET("video_sc", ctrl.CommeVideoSC)

			// type ? (superChat) and name ? その人(name)のsuperChatチャンネル内全て
			// api/comme/name_sc?name=xxx&chid=xxx
			comme.GET("/name_sc", ctrl.CommeNameSC)

			// type ? (superChat) and name ? video_id(url) ?　その人(name)のsuperChat動画内全て
			// api/comme/namevideo_sc?name=xxx&id=xxx&chid=xxx
			comme.GET("/namevideo_sc", ctrl.CommeNameVideoSC)

			// message ? (like) そのチャンネルでコメントを検索、全て
			// api/comme/chmsg_simi?chid=xxx&msg=xxx
			comme.GET("/chmsg_simi", ctrl.CommeChMsg)
			// message ? (like) video_id ? その動画内でのコメント検索、全て
			// api/comme/chvimsg_simi?chid=xxx&id=xxx&msg=xxx
			comme.GET("/chvimsg_simi", ctrl.CommeChViMsg)

		}
	}
	return r
}
