package main

import (
	"fmt"
	"pack"
	"strings"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	pack.Printsl()
	r := gin.Default()
	r.Static("/views/static/", "views/static")
	// r.Static("/views/static/css/", "views/static/css")
	// r.Static("/views/static/js/", "/views/static/js")
	// r.Static("/views/static/img/", "views/static/img")

	r.LoadHTMLGlob("views/template/*.html")
	r.GET("/", indexHandler)
	r.GET("/views/template/*.html", otherHandler)
	// fmt.Println(abc)
	r.POST("/", formHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func otherHandler(c *gin.Context) {
	url1 := c.Request.URL.Path
	fmt.Println(url1)
	x := strings.Split(url1, "/")
	// s1 := "views/template/" + x[2] + ".html"
	s1 := x[3]
	fmt.Println(s1)
	c.HTML(200, s1, nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	fmt.Print(fakeForm.Colors)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}
