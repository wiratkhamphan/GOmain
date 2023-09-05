package main

import (
	"example/hello/views"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	views.WD(router)
	http.ListenAndServe(":8080", router)
}
