package user

import (
	"log"
	apierrors "needmov/APIerrors"
	"needmov/db"
	"needmov/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIInsterChInfo ch情報をjsonで受け取りdbに保存する
// "api/pri" "POST" bindJSON entity.ChannelInfos = ch
func (pc Controller) APIInsterChInfo(c *gin.Context) {
	var ch entity.ChannelInfos
	if err := c.BindJSON(&ch); err != nil {
		c.JSON(http.StatusBadRequest, *apierrors.ErrJson)
		return
	}
	// ch.CreatedAt(time.Time) は、登録しない
	db.InsertChannelInfo(ch.ChannelID, ch.ChannelName, ch.ViewCount, ch.SubscriberCount, ch.VideoCount)
	// msg => ch で返ってくる"create_at": "0001-01-01T00:00:00Z"　エラーではない
	c.JSON(http.StatusOK, gin.H{"msg": ch})
}

// APIInsterChURL urlを登録する１つだけ
// "api/reg?url="
func (pc Controller) APIInsterChURL(c *gin.Context) {
	url := c.Query("url")
	// urlが正しいかチェック
	if len(url) == 24 && url[0:2] == "UC" { // urlが24文字でUCで始まるのはOK
		_, err := db.InsterRegChannel(url) // urlの登録。err は、urlが重複するかチェック
		if err != nil {
			//Error msg その、URLはすでにあります。
			c.JSON(http.StatusOK, *apierrors.ErrDuplicateURL)
		} else {
			//Success msg 追加しました！
			ch := entity.RegChannel{ChannelID: url}
			c.JSON(http.StatusOK, ch)
		}
	} else {
		//Error msg そのURLは正しくありません。もう一度確認お願いします！
		c.JSON(http.StatusOK, *apierrors.ErrInvalidURL)
	}
}

// APIAllGetChannelInfo "api/ch-info"apiで登録したデータベースを全部取る
// "api/ch-info"
func (pc Controller) APIAllGetChannelInfo(c *gin.Context) {
	channelInfos, err := db.GetDBChannelInfo()
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, channelInfos)
}

// APISelectWho 選んだ人の、チャンネルを取る
// "api/ch-sel?who-ch="
func (pc Controller) APISelectWho(c *gin.Context) {
	db := db.ConnectGorm()
	var chInfos []entity.ChannelInfos
	who := c.Query("who-ch")

	err := db.Where("channel_id = ?", who).Find(&chInfos).Error
	if err != nil {
		log.Printf(":%v", err)
	}
	c.JSON(http.StatusOK, chInfos)
}

var chInfos []entity.ChannelInfos

//APISelectDate 選んだ人と日付？
// "api/date-sel?who-ch=&date="
func (pc Controller) APISelectDate(c *gin.Context) {
	db := db.ConnectGorm()
	date := c.Query("date")
	id := c.Query("who-ch")
	db.Where("created_at LIKE ? AND channel_id = ?", "%"+date+"%", id).Find(&chInfos)
	c.JSON(http.StatusOK, chInfos)
}

// APISelectLatest 選んだ人と最新の日付
// "api/latest-ch?who-ch="
func (pc Controller) APISelectLatest(c *gin.Context) {
	db := db.ConnectGorm()
	id := c.Query("who-ch")
	db.Where("channel_id = ?", id).First(&chInfos)
	c.JSON(http.StatusOK, chInfos)
}

// APISelectDateBetween 選んだ人とBETWEEN日付の選択
// "api/date-between?who-ch=&a=&b="
func (pc Controller) APISelectDateBetween(c *gin.Context) {
	db := db.ConnectGorm()
	id := c.Query("who-ch")
	past1 := c.Query("a")
	past2 := c.Query("b")
	db.Where("channel_id = ? AND created_at BETWEEN ? AND ?", id, past1, past2).Find(&chInfos)
	c.JSON(http.StatusOK, chInfos)
}

//func (pc Controller) API (c *gin.Context) {}
