package main

import (
	"TinyWeb/tiny"
	"net/http"
)

func main() {
	r := tiny.Default()
	r.GET("/", func(ctx *tiny.Context) {
		ctx.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(ctx *tiny.Context) {
		names := []string{"geektutu"}
		ctx.String(http.StatusOK, names[100])
	})

	_ = r.Run(":9999")
}
