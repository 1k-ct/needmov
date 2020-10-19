package user

import (
	"log"
	"needmov/db"
	"needmov/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIAllGetChannelInfo "api/ch-info"apiで登録したデータベースを全部取る
// "api/ch-info"
func (pc Controller) APIAllGetChannelInfo(c *gin.Context) {
	channelInfos, err := db.AllGetDBChannelInfo("ChannelInfos")
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
	db.Where("channel_id = ?", who).Find(&chInfos)
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
