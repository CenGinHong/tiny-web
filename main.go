package main

import (
	"TinyWeb/tiny"
	"net/http"
)

func main() {
	r := tiny.New()
	r.GET("/", func(ctx *tiny.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *tiny.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *tiny.Context) {
		c.JSON(http.StatusOK, tiny.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	_ = r.Run(":9999")
}
