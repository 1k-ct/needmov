## はじめに
https://needmov239087.df.r.appspot.com/api/ch-info  
まずは、このURLを見てきてください。  
これは登録してある人のチャンネルの簡単な情報を表示しています。
## 使い方
./server/server.go 88行目[ここ](https://github.com/1k-ct/needmov/blob/db%E3%81%BE%E3%81%A8%E3%82%81%E3%81%9F/server/server.go#L88)からを見てください  
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

