package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func testform(method string, url string, body io.Reader) (w *httptest.ResponseRecorder, c *gin.Context) {
	router := router(false)
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(method, url, body)
	router.ServeHTTP(w, c.Request)
	return
}
func TestChInfo(t *testing.T) {
	// router := router(false)
	// w := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(w)
	// // body := bytes.NewBufferString("{\"name\":\"foo\"}")
	// c.Request, _ = http.NewRequest("GET", "/api/ch-info", nil)
	// // req, _ := http.NewRequest("POST", "/ps", body)
	// router.ServeHTTP(w, c.Request)
	w, _ := testform("GET", "/api/ch-info", nil)
	// log.Println("-------", w.Body.String())
	assert.Equal(t, w.Code, 200)
}
func TestChSel(t *testing.T) {
	// w, _ := testform("POST", "/reg?url=xxx", nil)
	w, c := testform("GET", "/api/ch-sel?who-ch=xxx", nil)

	assert.Equal(t, c.Query("who-ch"), "xxx")
	assert.Equal(t, w.Code, 200)
}
func TestDateSel(t *testing.T) {
	w, c := testform("GET", "/api/date-sel?who-ch=xxx", nil)

	assert.Equal(t, c.Query("who-ch"), "xxx")
	assert.Equal(t, w.Code, 200)
}
func TestLatestCh(t *testing.T) {
	w, c := testform("GET", "/api/latest-ch?who-ch=xxx", nil)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, c.Query("who-ch"), "xxx")
}

func TestDateBetween(t *testing.T) {
	w, c := testform("GET", "/api/date-between?who-ch=xxx", nil)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, c.Query("who-ch"), "xxx")
}

// postわからんくなったPOSTはテストあとで
// func TestReg(t *testing.T) {
// 	url := "/api/reg?url=UCxxxxxxxxxxxxxxxxxxxxxx"
// 	router := router(false)
// 	w := httptest.NewRecorder()
// 	//c, _ := gin.CreateTestContext(w)
// 	req, _ := http.NewRequest("POST", url, nil)
// 	router.ServeHTTP(w, req)
// 	var c *gin.Context
// 	assert.Equal(t, 200, w.Code)

// 	n := strings.Index(url, "=")
// 	log.Println(url[n+1:], "-------")
// 	// assert.Equal(t, w.Code, 200)
// 	assert.Equal(t, c.Query("url"), url[n+1:])
// }
