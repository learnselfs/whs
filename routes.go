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
	handler     Handler           // handler is processing the Route
	isHandle    bool              // isHandle
	isAsterisk  bool              // isAsterisk matching
	isColon     bool              // isColon matching
	param       map[string]string // params
	middlewares []Handler
}

// newRoute for return Route
func newRoute() *Route {
	r := &Route{index: -1, routes: make(map[string]*Route), param: make(map[string]string), middlewares: make([]Handler, 0)}
	return r
}

// RegisterRouter for add router
func (r *Route) RegisterRouter(url string, handler Handler) {
	urlList := parseUrl(url)
	recursionRegisterRouter(r, 0, urlList, handler)
}

func (r *Route) Router(url string) []Handler {
	urlList := parseUrl(url)
	var router *Route
	router = r
	router = recursionRouter(router, urlList, r.param)
	if router != nil && router.isHandle {
		return router.middlewares
	}
	return nil
}

func recursionRegisterRouter(r *Route, count int, urls []string, handler Handler) {
	if r.isAsterisk || count >= len(urls) {
		r.handler = handler
		if handler != nil {
			r.isHandle = true
			r.middlewares = append(r.middlewares, r.handler)
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
	router.index = index
	router.prefix = url
	router.isColon = isColon
	router.isAsterisk = isAsterisk

	router.middlewares = append(router.middlewares, r.middlewares...)
	recursionRegisterRouter(router, count+1, urls, handler)

}

func recursionRouter(r *Route, urls []string, param map[string]string) *Route {
	if r.isAsterisk || r.isHandle || r.index >= len(urls)-1 {
		return r
	}

	url := urls[r.index+1]
	router, ok := r.routes[url]
	if ok {
		return recursionRouter(router, urls, param)
	}
	for k, v := range r.routes {
		if strings.HasPrefix(k, `:`) {
			param[v.prefix] = url
			return recursionRouter(v, urls, param)
		}
		if strings.HasPrefix(k, `*`) {
			return recursionRouter(v, urls, param)
		}
	}
	return nil
}

// Group route
func (r *Route) Group(prefix string) *Route {
	prefix = parseUrlExcludeSpecialSymbol(prefix)
	r.RegisterRouter(prefix, nil)
	return r.routes[prefix]
}

// UseMiddleware complete
func (r *Route) UseMiddleware(middleware Handler) {
	r.middlewares = append(r.middlewares, middleware)
}
