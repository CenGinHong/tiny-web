## TinyWeb

### 简介

练手项目，使用Go开发的简易web框架，支持快速简易搭建一个Web应用。

### 快速入门

```go
package main

import (
	"github.com/CenGinHong/TinyWeb/tiny"
	"net/http"
)

func main() {
	r := tiny.Default()
	r.GET("/", func(ctx *tiny.Context) {
		ctx.String(http.StatusOK, "Hello World！\n")
	})
	_ = r.Run(":9999")
}
```

运行以上代码并在浏览器中访问 [http://localhost:9999](http://localhost:9999)，即可在浏览器中访问到`Hello World！`

### 使用RESTful API

```go
package main

import (
	"github.com/CenGinHong/TinyWeb/tiny"
	"net/http"
)

func main() {
	r := tiny.Default()
	r.GET("/", func(ctx *tiny.Context) {
        ctx.String(http.StatusOK, "GET:Hello World！\n")
	})
	r.POST("/", func(ctx *tiny.Context) {
        ctx.String(http.StatusOK, "GET:Hello World！\n")
	})
    r.DELETE("/", func(ctx *tiny.Context) {
        ctx.String(http.StatusOK, "GET:Hello World！\n")
	})
    r.PUT("/", func(ctx *tiny.Context) {
        ctx.String(http.StatusOK, "GET:Hello World！\n")
	})
	_ = r.Run(":9999")
}
```



### 模糊匹配路径参数

```go
package main

import (
	"github.com/CenGinHong/TinyWeb/tiny"
	"net/http"
)

func main() {
	r := tiny.Default()
	r.GET("/:id", func(ctx *tiny.Context) {
		ctx.String(http.StatusOK, "Hello World! %s\n", ctx.Param("id"))
	})
	_ = r.Run(":9999")
}
```

运行以上代码并在浏览器中访问 [http://localhost:9999/123](http://localhost:9999/123)，即可在浏览器中访问到`Hello World！ 123`



### 路由分组

```go
package main

import (
	"github.com/CenGinHong/TinyWeb/tiny"
	"net/http"
)

func main() {
	r := tiny.Default()
	r.GET("/index", func(c *tiny.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *tiny.Context) {
			c.HTML(http.StatusOK, "<h1>Hello v1</h1>")
		})

		v1.GET("/hello", func(c *tiny.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *tiny.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *tiny.Context) {
			c.JSON(http.StatusOK, tiny.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
```

实现路由分组，同一分组共享路由前缀



### 使用中间件

```go
package main

import (
	"github.com/CenGinHong/TinyWeb/tiny"
	"net/http"
)

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := tiny.Default()
	r.Use(gee.Logger()) 
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
```

`Router.Use`用于添加中间件，中间件以如下形式组织，part1将作为handler前置中间件部分，part2作为后置中间件部分

```go
func A(c *Context) {
    part1
    c.Next()
    part2
}
```

### 模板渲染

建立结构目录如下

```
---static/
   |---css/
        |---tiny.css
---templates/
   |---css.tmpl
---main.go
```

```css
<!-- tiny.css -->
p {
    color: orange;
    font-weight: 700;
    font-size: 20px;
}
```

```html
<!-- css.tmpl -->
<html>
    <link rel="stylesheet" href="/assets/css/tiny.css">
    <p>tiny.css is loaded</p>
</html>
```

```go
func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.Run(":9999")
}
```

运行以上代码并在浏览器中访问 [http://localhost:9999](http://localhost:9999)，即可在浏览器中访问到模板文件