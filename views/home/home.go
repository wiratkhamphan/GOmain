package home

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func Myhome(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("static/home.html"))
	tmpl.Execute(c.Writer, nil)
}
