package server

import (
	"fmt"
	"io"
	"needmov/db"
	"needmov/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_defMiddleware(t *testing.T) {
	fn := func(key string) (entity.APIKEY, error) {
		var k entity.APIKEY
		db := db.ConnectGorm()
		defer db.Close()
		if err := db.Where("self_key = ?", key).Find(&k).Error; err != nil {
			return k, err
		}
		return k, nil
	}
	tests := []struct {
		args    string
		want    string
		wantErr bool
	}{
		{
			args:    "test-key-no1",
			want:    "test-key-no1",
			wantErr: false,
		},
		{
			args:    "test-false",
			want:    "test-key-no1",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		ke, err := fn(tt.args)
		fmt.Println(err)
		if (err != nil) != tt.wantErr {
			t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
		}
		if (ke.SelfKey == tt.want) == tt.wantErr {
			t.Errorf("fn() = %v, want %v", ke.SelfKey, tt.want)
		}
	}
}

func testform(method string, url string, body io.Reader) (w *httptest.ResponseRecorder, c *gin.Context) {
	router := router(false)
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(method, url, body)
	router.ServeHTTP(w, c.Request)
	return
}
func Test_ChInfo(t *testing.T) {
	w, _ := testform("GET", "/api/ch-info", nil)
	assert.Equal(t, w.Code, 200)
}
func Test_ChSel(t *testing.T) {
	// w, _ := testform("POST", "/reg?url=xxx", nil)
	w, c := testform("GET", "/api/ch-sel?who-ch=xxx", nil)

	assert.Equal(t, c.Query("who-ch"), "xxx")
	assert.Equal(t, w.Code, 200)
}
func Test_DateSel(t *testing.T) {
	w, c := testform("GET", "/api/date-sel?who-ch=xxx", nil)

	assert.Equal(t, c.Query("who-ch"), "xxx")
	assert.Equal(t, w.Code, 200)
}
func Test_LatestCh(t *testing.T) {
	w, c := testform("GET", "/api/latest-ch?who-ch=xxx", nil)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, c.Query("who-ch"), "xxx")
}

func Test_DateBetween(t *testing.T) {
	w, c := testform("GET", "/api/date-between?who-ch=xxx", nil)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, c.Query("who-ch"), "xxx")
}

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
