package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller is user controller
type Controller struct{}

// Index is start page "/"
func (pc Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "start.html", gin.H{})
}

// HashibaDeteil hashibadeteil page "/hashiba/"
func (pc Controller) HashibaDeteil(c *gin.Context) {
	c.HTML(http.StatusOK, "hashibadeteil.html", gin.H{})
}

// HashibaHome hashiba home page "/hashiba/home"
func (pc Controller) HashibaHome(c *gin.Context) {
	c.HTML(http.StatusOK, "hashibahome.html", gin.H{})
}

// Adimn adimn page "/adimn"
func (pc Controller) Adimn(c *gin.Context) {
	c.HTML(http.StatusOK, "adim.html", gin.H{})
}
