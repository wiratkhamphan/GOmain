package views

import (
	welcome "example/hello/views/Welcome"
	home "example/hello/views/home"

	"github.com/gin-gonic/gin"
)

func WD(router *gin.Engine) {
	router.GET("/", welcome.MyFunction)
	router.POST("/login", welcome.LoginHandler)
	router.GET("/Home", home.Myhome)
}
