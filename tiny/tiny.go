package tiny

import (
	"log"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Engine struct {
	*RouterGroup
	router *router        // 管理路由前缀树
	groups []*RouterGroup // 管理所有RouterGroup
}

func New() *Engine {
	e := &Engine{router: newRouter()}
	e.RouterGroup = &RouterGroup{engine: e}
	e.groups = []*RouterGroup{e.RouterGroup}
	return e
}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	e := g.engine
	newGroup := &RouterGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: e,
	}
	e.groups = append(e.groups, newGroup)
	return newGroup
}

func (e *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	e.router.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

// RouterGroup 路由组，使用自底向上的方式进行连接
type RouterGroup struct {
	prefix      string       // 前缀
	middlewares []HandleFunc // 应用在此上的中间件
	parent      *RouterGroup // 父级路由组
	engine      *Engine      // 所有的RouterGroup均持有同一个engine
}

func (g *RouterGroup) addRoute(method string, comp string, handler HandleFunc) {
	pattern := g.prefix + comp
	log.Printf("Route %4s -%s", method, pattern)
	g.engine.router.addRoute(method, pattern, handler)
}

func (g *RouterGroup) GET(pattern string, handler HandleFunc) {
	g.addRoute("GET", pattern, handler)
}

func (g *RouterGroup) POST(pattern string, handler HandleFunc) {
	g.addRoute("POST", pattern, handler)
}
