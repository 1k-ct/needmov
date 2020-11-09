package user

import (
	"log"
	apierrors "needmov/APIerrors"
	"needmov/crypto"
	"needmov/db"
	"needmov/entity"
	user "needmov/service"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Controller is user controller
type Controller struct{}

// Connection DB接続確認テスト
func (pc Controller) Connection(c *gin.Context) {
	db.NewMakeDB()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// Start is start page "/"
func (pc Controller) Start(c *gin.Context) {
	var sessioninfo entity.SessionInfo
	log.Println(sessioninfo.ID)
	c.HTML(http.StatusOK, "start.html", gin.H{})
}

// HashibaDeteil 羽柴さんのvideoDB, channelDBの情報を全て表示する
// "/hashiba/reg"
func (pc Controller) HashibaDeteil(c *gin.Context) {
	var s user.Service
	var url string = "UC_BlXOQe5OcRC7o0GX8kp8A" //羽柴チャンネル　特別です！
	//videoInfos := db.GetDBVideoInfo()
	channelInfos, err := s.GetSomeoneChannelInfo(c, url)
	if err != nil {
		c.AbortWithStatusJSON(404, apierrors.ErrDB)
	}
	/*
		channelInfos, err := db.GetDBChannelInfo()
		if err != nil {
			log.Println(err)
		}
	*/
	c.HTML(http.StatusOK, "hashibadeteil.html", gin.H{
		"channelInfos": channelInfos,
	})
	c.JSON(http.StatusOK, channelInfos)
}

// HashibaHome 羽柴ホーム "/hashiba"
func (pc Controller) HashibaHome(c *gin.Context) {
	//c.AbortWithStatusJSON(404, apierrors.ErrInvalidURL)
	c.HTML(http.StatusOK, "hashibahome.html", gin.H{})
}

// ShiromiyaHome 白宮ホーム　"/shiromiya/"
func (pc Controller) ShiromiyaHome(c *gin.Context) {
	c.HTML(http.StatusOK, "shiromiyahome.html", gin.H{})
}

// ShiromiyaRegVideo 白宮さんのvideoDBの情報を全て表示する
func (pc Controller) ShiromiyaRegVideo(c *gin.Context) {
	var s user.Service
	var url string = "UCtzCQnCT9E4o6U3mHHSHbQQ" //白宮チャンネル　特別です！

	/*
		videoInfos, err := db.GetDBVideoInfo()
		if err != nil {
			log.Println(err)
		}
	*/
	channelInfos, err := s.GetSomeoneChannelInfo(c, url)
	if err != nil {
		c.AbortWithStatusJSON(404, apierrors.ErrDB)
	}
	c.HTML(http.StatusOK, "shiromiyadeteil.html", gin.H{
		"channelInfos": channelInfos,
	})
	c.JSON(http.StatusOK, channelInfos)
}

// Adimn adimn page "/adimn"
func (pc Controller) Adimn(c *gin.Context) {
	c.HTML(http.StatusOK, "admin.html", gin.H{})
}

// SignUpGet "/signup" "signup.html"　ユーザー登録画面
func (pc Controller) SignUpGet(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", gin.H{})
}

// SignUpPost "/signup" "signup.html" ユーザー登録
func (pc Controller) SignUpPost(c *gin.Context) {
	var form entity.Users
	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		c.Abort()
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 登録ユーザーが重複していた場合にはじく処理
		if err := db.CreateUser(username, password); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
		}
		c.Redirect(302, "/") // c.Redirect(302, "/login")
	}
}

// LoginGet "/login" ユーザーログイン画面
func (pc Controller) LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// LoginPost "/login" ユーザーログイン
func (pc Controller) LoginPost(c *gin.Context) {
	// DBから取得したユーザーパスワード(Hash)
	//dbPassword := db.GetUser(c.PostForm("username")).Password
	dbPassword := db.GetUser(c.PostForm("username")).Password
	//log.Println("some numbers1")
	//log.Println(dbPassword)
	// フォームから取得したユーザーパスワード
	formPassword := c.PostForm("password")
	//log.Println(formPassword)
	// ユーザーパスワードの比較
	if err := crypto.CompareHashAndPassword(dbPassword, formPassword); err != nil {
		log.Println("ログインできませんでした")
		c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
		c.Abort()
	} else {
		login(c, formPassword)
		log.Println("ログインできました")
		c.Redirect(http.StatusMovedPermanently, "/") // "/" "/hashiba/home"
		//c.Redirect(302, "/")
	}
}

// PostLogout logout処理
func (pc Controller) PostLogout(c *gin.Context) {
	log.Println("ログアウト処理")
	//セッションからデータを破棄する
	session := sessions.Default(c)
	log.Println("セッション取得")
	session.Clear()
	log.Println("クリア処理")
	session.Save()

	// ログインフォームに戻す
	//var user entity.UsersMig
	c.HTML(http.StatusOK, "start.html", gin.H{})
}
func login(c *gin.Context, ID string) {
	//セッションにデータを格納する
	session := sessions.Default(c)
	session.Set("ID", ID)
	session.Save()
}
