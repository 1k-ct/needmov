package user

import (
	"encoding/base64"
	"log"
	apierrors "needmov/APIerrors"
	"needmov/db"
	"needmov/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 共通部分ですどうするか
// var d []entity.Data
// db := db.ConnectGorm()
func cone(obj ...interface{}) *gorm.DB {
	db := db.ConnectGorm()
	return db
}
func encodingMsg(b string) string {
	// Base64 エンコード
	data := []byte(b)

	enc := base64.StdEncoding.EncodeToString(data)
	return enc
}
func decodingMsg(enc string) (string, error) {
	// Base64 decoding
	dec, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return "", err
	}
	return string(dec), nil
}

// APIInsertCommentData データベースにentity.Dataを登録する。
// "api/data" "POST" binding database
func (pc Controller) APIInsertCommentData(c *gin.Context) {
	var d entity.Data
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "error"})
		return
	}
	d.Name = encodingMsg(d.Name)
	d.Message = encodingMsg(d.Message)
	// d.Name d.Message encoding start
	if err := db.CreateDBSelfFn(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	name, err := decodingMsg(d.Name)
	if err != nil {
		log.Println(err)
	}
	msg, err := decodingMsg(d.Message)
	if err != nil {
		log.Println(err)
	}
	d.Name, d.Message = name, msg
	c.JSON(http.StatusOK, &d)
}

// CommeName nameだけ
// "api/comme/namesel?name=xxx"
func (pc Controller) CommeName(c *gin.Context) {
	var d []entity.Data
	db := db.ConnectGorm()
	name := c.Query("name")
	name = encodingMsg(name)
	err := db.Where("name = ?", name).Find(&d).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.ErrDB)
		return
	}
	c.JSON(http.StatusOK, d)
}

// CommeNameCom name and video_id
// api/comme/namecommesel?name=xxx&id=xxx(url)
func (pc Controller) CommeNameCom(c *gin.Context) {
	var d []entity.Data
	db := db.ConnectGorm()
	name := c.Query("name")
	name = encodingMsg(name)
	vID := c.Query("id")
	err := db.Where("name = ? AND video_url = ?", name, vID).Find(&d).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.ErrDB)
		return
	}
	c.JSON(http.StatusOK, d)
}

// CommeAllSC 全てのsuperChat
// "api/comme/allsc?chid=xxxxx"
func (pc Controller) CommeAllSC(c *gin.Context) {
	var d []entity.Data
	db := db.ConnectGorm()
	ch := c.Query("chid")
	s := "superChat"
	err := db.Where("type = ? AND master_channel_id = ?", s, ch).Find(&d).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.ErrDB)
		return
	}
	c.JSON(http.StatusOK, d)
}

// CommeVideoSC そのチャンネルの動画のsuperChat全部
// api/comme/videosc?chid=xxx&id=xxx
func (pc Controller) CommeVideoSC(c *gin.Context) {
	var d []entity.Data
	db := db.ConnectGorm()
	s := "superChat"
	mChID := c.Query("chid")
	vID := c.Query("id")
	err := db.Where("type = ? AND master_channel_id = ? AND video_url = ?", s, mChID, vID).Find(&d).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.ErrDB)
		return
	}
	c.JSON(http.StatusOK, d)
}

// CommeNameSC そのチャンネルでその人(name)がsuperChatした全部
// api/comme/namesc?name=xxx&chid=xxx
func (pc Controller) CommeNameSC(c *gin.Context) {
	var d []entity.Data
	db := db.ConnectGorm()
	s := "superChat"
	name := encodingMsg(c.Query("name"))
	mChID := c.Query("chid")
	err := db.Where("type = ? AND name = ? AND master_channel_id = ?", s, name, mChID).Find(&d).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.ErrDB)
		return
	}
	c.JSON(http.StatusOK, d)
}

// CommeNameSC チャンネルの動画内のsuperChatしたname選択
// "api/comme/namesc?name=xxx&"
func (pc Controller) CommeNameVideoSC(c *gin.Context) {
	var d []entity.Data
	db := db.ConnectGorm()
	s := "superChat"
	// name := c.Query("name")
	// name = encodingMsg(name)
	name := encodingMsg(c.Query("name"))
	vID := c.Query("id")
	mChID := c.Query("chid")
	err := db.Where("type = ? AND name = ? AND video_url = ? AND master_channel_id = ?", s, name, mChID, vID).Find(&d).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.ErrDB)
		return
	}
	c.JSON(http.StatusOK, d)
}

// CommeChMsg そのチャンネルでコメント類似検索、全て
// api/comme/chmsg?chid=xxx&msg=xxx
func (pc Controller) CommeChMsg(c *gin.Context) {
	var d []entity.Data
	db := db.ConnectGorm()
	mChID := c.Query("chid")
	msg := encodingMsg(c.Query("msg"))
	err := db.Where("master_channel_id = ? AND message LIKE ?", mChID, "%"+msg+"%").Find(&d).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.ErrDB)
		return
	}
	c.JSON(http.StatusOK, d)
}

// CommeChViMsg そのチャンネルの動画内でのコメントを類似検索、全て
// api/comme/chvimsg?chid=xxx&id=xxx&msg=xxx
func (pc Controller) CommeChViMsg(c *gin.Context) {
	var d []entity.Data
	db := db.ConnectGorm()
	mChID := c.Query("chid")
	msg := encodingMsg(c.Query("msg"))
	vID := c.Query("id")
	err := db.Where("master_channel_id = ? AND message LIKE ? AND video_url = ?", mChID, "%"+msg+"%", vID).Find(&d).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apierrors.ErrDB)
		return
	}
	c.JSON(http.StatusOK, d)
}
