package welcome

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyFunction(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("static/login.html"))
	tmpl.Execute(c.Writer, nil)
}
func LoginHandler(c *gin.Context) {

	c.Redirect(http.StatusSeeOther, "/Home")
}
