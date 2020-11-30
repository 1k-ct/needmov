package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"needmov/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is user controller
type Controller struct{}

// var d []entity.Data2

// var j string = "http://localhost:8080/api/comme/chmsg_simi?chid=UCvUc0m317LWTTPZoBQV479A&msg=草"

func myReqJSON(url string) ([]entity.Data3, error) {
	res, err := http.Get(url) //https://qiita.com/tutuz/items/fedb8e3a1137d046f418
	if err != nil {
		return d, err
	}
	if res.StatusCode != http.StatusOK {
		return d, errors.New("no statusOK")
	}
	defer res.Body.Close() //http.Getのリクエストをエラーチェックしてから
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return d, err
	}
	if err := json.Unmarshal(body, &d); err != nil {
		return d, err
	}
	return d, nil
}
func (pc Controller) GetLikeAMsg(c *gin.Context) {
	mChID := c.DefaultPostForm("chid", "UCvUc0m317LWTTPZoBQV479A")
	msg := c.PostForm("msg")
	vID := c.PostForm("id")
	url := "api/comme/chvimsg_simi?chid=" + mChID + "&id=" + vID + "&msg=" + msg

	dv, err := myReqJSON(url)
	if err != nil {
		c.HTML(http.StatusOK, "comme.html", gin.H{})
		return
	}
	c.HTML(http.StatusOK, "comme.html", gin.H{
		"commedata": dv,
	})
	log.Println(dv)
}
