heroku addons:create heroku-postgresql:<PLAN_NAME>
error になる場合アプリ選択

heroku addons:create heroku-postgresql:hobby-dev -a <heroku アプリ名>

func router() *gin.Engine {
    r := gin.Default()

    r.LoadHTMLGlob("templates/**/*") // どっちか、ファイルによる
    r.LoadHTMLGlob("templates/*.html") // あと、忘れやすい!
    
    r.GET("/", xxx.yyy)
}