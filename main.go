package main

import (
	"TinyWeb/tiny"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func FormAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := tiny.New()
	r.Use(tiny.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormAsDate,
	})
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")
	r.GET("/index", func(ctx *tiny.Context) {
		ctx.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *tiny.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", tiny.H{
			"title":  "tiny",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *tiny.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", tiny.H{
			"title": "tiny",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})
	_ = r.Run(":9999")
}
