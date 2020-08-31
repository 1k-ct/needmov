package server

import (
	user "needmov/controller"

	"github.com/gin-gonic/gin"
)

// Init is server run
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*") //*/**
	ctrl := user.Controller{}

	r.GET("/", ctrl.Index)
	r.GET("/admin", ctrl.Adimn)

	r.GET("/signup", ctrl.SignUpGet)
	r.POST("/signup", ctrl.SignUpPost)
	r.GET("/login", ctrl.LoginGet)
	r.POST("/login", ctrl.LoginPost)

	hashiba := r.Group("/hashiba")
	{
		hashiba.GET("/", ctrl.HashibaDeteil)
		hashiba.GET("/home", ctrl.HashibaHome)
	}
	return r
}
