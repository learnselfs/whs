// Package whs @Author Bing
// @Date 2024/2/3 20:37:00
// @Desc
package whs

import (
	"strings"
)

type Handler func(ctx *Context)

// Route struct
type Route struct {
	routes      map[string]*Route // routes
	index       int               // index of Route
	prefix      string            // prefix is the Route url
	pattern     string            // prefix is the Route url
	method      string            // method
	handler     Handler           // handler is processing the Route
	isHandle    bool              // isHandle
	isAsterisk  bool              // isAsterisk matching
	isColon     bool              // isColon matching
	param       map[string]string // params
	middlewares []Handler
	handlers    map[string][]Handler
}

// newRoute for return Route
func newRoute() *Route {
	r := &Route{index: -1, routes: make(map[string]*Route), param: make(map[string]string), middlewares: make([]Handler, 0), handlers: make(map[string][]Handler)}
	return r
}

// RegisterRouter for add router
func (r *Route) RegisterRouter(method, url string, handler Handler) {
	urlList := parseUrl(url)
	recursionRegisterRouter(r, 0, urlList, handler, method)
}

func (r *Route) GET(url string, handler Handler) {
	r.RegisterRouter("GET", url, handler)
}

func (r *Route) POST(url string, handler Handler) {
	r.RegisterRouter("POST", url, handler)
}

func (r *Route) DELETE(url string, handler Handler) {
	r.RegisterRouter("DELETE", url, handler)
}
func (r *Route) PUT(url string, handler Handler) {
	r.RegisterRouter("PUT", url, handler)
}
func (r *Route) Router(method, url string) *Route {
	urlList := parseUrl(url)
	var router *Route
	router = r
	router = recursionRouter(router, urlList, r.param, method)
	if router != nil && router.isHandle {
		return router
	}
	return nil
}

func recursionRegisterRouter(r *Route, count int, urls []string, handler Handler, method string) {
	if r.isAsterisk || count >= len(urls) {
		r.handler = handler
		if handler != nil {
			r.isHandle = true
			r.handlers[method] = append(r.middlewares, r.handler)
			p := strings.Join(urls, "/")
			p = r.pattern + "/" + p
			pattern = append(pattern, p)
		} else {
			r.pattern = strings.Join(urls, "/")
		}
		return
	}
	var isColon bool
	var isAsterisk bool
	index := r.index
	index++

	url := urls[count]
	tmpUrl := url
	if strings.HasPrefix(url, `:`) {
		url = url[1:]
		isColon = true
	}
	if strings.HasPrefix(url, `*`) {
		isAsterisk = true
		url = "*"
	}

	router, ok := r.routes[tmpUrl]
	if !ok {
		router = newRoute()
		r.routes[tmpUrl] = router
	}
	router.pattern = r.pattern
	router.index = index
	router.prefix = url
	router.isColon = isColon
	router.isAsterisk = isAsterisk

	router.middlewares = append(router.middlewares, r.middlewares...)
	recursionRegisterRouter(router, count+1, urls, handler, method)

}

func recursionRouter(r *Route, urls []string, param map[string]string, method string) *Route {
	if r.isAsterisk || (r.isHandle && r.index >= len(urls)-1) {
		return r
	}

	url := urls[r.index+1]
	tempUrl := url
	router, ok := r.routes[tempUrl]
	if ok {
		return recursionRouter(router, urls, param, method)
	}
	for k, v := range r.routes {
		if strings.HasPrefix(k, `:`) {
			param[v.prefix] = url
			return recursionRouter(v, urls, param, method)
		}
		if strings.HasPrefix(k, `*`) {
			return recursionRouter(v, urls, param, method)
		}
	}
	return nil
}

// Group route
func (r *Route) Group(prefix string) *Route {
	prefix = parseUrlExcludeSpecialSymbol(prefix)
	r.GET(prefix, nil)
	return r.routes[prefix]
}

// UseMiddleware complete
func (r *Route) UseMiddleware(middleware Handler) {
	r.middlewares = append(r.middlewares, middleware)
}
