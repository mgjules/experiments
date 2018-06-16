package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFS("/res", http.Dir("./res"))
	r.LoadHTMLGlob("tmpls/*")

	r.GET("/", indexPage)
	r.GET("/about", aboutPage)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func indexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Index page",
		"content": "Placeholder for index page",
	})
}

func aboutPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "About page",
		"content": "Placeholder for about page",
	})
}
