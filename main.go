package main

import (
	"fmt"
	"pack"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	pack.Printsl()
	r := gin.Default()
	// r.Static("/", "views")
	r.Static("/css", "views/css")
	r.Static("/js", "views/js")
	r.Static("/img", "views/img")

	r.LoadHTMLGlob("views/*.html")
	r.GET("/", indexHandler)
	r.GET("/index.html", indexHandler)
	r.GET("/form1.html", form1Handler)
	r.POST("/", formHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func form1Handler(c *gin.Context) {
	c.HTML(200, "form1.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	fmt.Print(fakeForm.Colors)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}
