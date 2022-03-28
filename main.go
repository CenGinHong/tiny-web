package main

import (
	"TinyWeb/tiny"
	"net/http"
)

func main() {
	r := tiny.New()
	r.GET("/index", func(ctx *tiny.Context) {
		ctx.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(ctx *tiny.Context) {
			ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(ctx *tiny.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(ctx *tiny.Context) {
			ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
		})

		v2.POST("/login", func(ctx *tiny.Context) {
			ctx.JSON(http.StatusOK, tiny.H{
				"username": ctx.PostForm("username"),
				"password": ctx.PostForm("password"),
			})
		})
	}
	_ = r.Run(":9999")
}
