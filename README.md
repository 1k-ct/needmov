heroku addons:create heroku-postgresql:<PLAN_NAME>
error になる場合アプリ選択

heroku addons:create heroku-postgresql:hobby-dev -a <heroku アプリ名>

func router() *gin.Engine {
    r := gin.Default()

    r.LoadHTMLGlob("templates/**/*") // どっちか、ファイルによる
    r.LoadHTMLGlob("templates/*.html") // あと、忘れやすい!
    
    r.GET("/", xxx.yyy)
}  
参考
https://github.com/Doarakko/api-challenge.git
```
db.DeleteDBChannelInfo(id)
```
データーベースにないIDを選択した場合  
errorになる => データベース消えた  
確認がしないと
接続確認ok
  
docker/api/Dockerfile  
下の無くてもいい
#RUN go get github.com/gin-contrib/sessions
#RUN go get github.com/gin-contrib/sessions/cookie
#RUN go get golang.org/x/crypto/bcrypt
#RUN go get google.golang.org/api/googleapi/transport
#RUN go get google.golang.org/api/youtube/v3