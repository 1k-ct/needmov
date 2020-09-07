package server

import (
	"log"
	user "needmov/controller"
	"needmov/entity"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Init is server run
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*") //*/**
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	ctrl := user.Controller{}

	r.GET("/", ctrl.Start)
	r.GET("/admin", ctrl.Adimn)

	r.GET("/signup", ctrl.SignUpGet)
	r.POST("/signup", ctrl.SignUpPost)
	r.GET("/login", ctrl.LoginGet)
	r.POST("/login", ctrl.LoginPost)

	hashiba := r.Group("/hashiba")
	hashiba.Use(sessionCheck())
	{
		hashiba.GET("/", ctrl.HashibaDeteil)
		hashiba.GET("/home", ctrl.HashibaHome)
	}
	/*
		menu := r.Group("/menu")
		menu.Use(sessionCheck())
		{
			menu.GET("/", ctrl.HashibaDeteil)
			menu.GET("/top", ctrl.HashibaHome)
		}
	*/
	r.GET("/logout", ctrl.PostLogout) //r.POST("/logout", ctrl.PostLogout)
	r.POST("/regvideo", ctrl.CreateVideoInfo)
	r.POST("/regchannel", ctrl.CreateChannelInfo)
	return r
}

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
