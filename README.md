## はじめに
https://needmov239087.df.r.appspot.com/api/ch-info  
まずは、このURLを見てきてください。  
これは登録してある人のチャンネルの簡単な情報を表示しています。
## 使い方
./server/server.go 88行目からを見てください  
補足で  ./controller/user_api.go　も見てみてください。
### /api/ch-infoにしていろいろ見てみください
- チャンネルURLを指定する  
例：  
https://www.youtube.com/channel/UCL-2thbJ7grC9fmGF4OLuTg  
/ch-sel?who-ch=UCL-2thbJ7grC9fmGF4OLuTg  
/ch-sel?who-ch=ここに上みたいにURLを入れてください。  
- 日付をしていする  
api/date-sel?who-ch=&date=ここに日(例：2020-01-03)を指定
- /api/ch-infoに欲しいチャンネル情報がないとき  
登録しよう！  
api/reg?url=UCxxxxxxxxxxxxxxxxxxxxxxのようにして登録する  
登録したら３時間待ってください。３時間後には、登録されてます。  
- ３時間ごとに更新しています。  

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

スライス  
var t []string -> nil  
t := []string{} -> 非nil(JSONオブジェクト)  
JSONオブジェクトをエンコードするときは非nilスライスが優先されます。  
