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
	ctrl := user.Controller{}
	r.GET("/", ctrl.Index)
	r.GET("/admin", ctrl.Adimn)
	u := r.Group("/hashiba")
	{
		u.GET("/", ctrl.HashibaDeteil)
		u.GET("/home", ctrl.HashibaHome)
	}
	return r
}
