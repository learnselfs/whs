// Package whs @Author Bing
// @Date 2024/2/3 20:37:00
// @Desc
package whs

import "github.com/learnselfs/wlog"

type Handler func(ctx *Content)

// route struct
type route struct {
	routes   map[string]*route // routes
	index    int               // index of route
	prefix   string            // prefix is the route url
	handler  Handler           // handler is processing the route
	isHandle bool              // isHandle
	isWail   bool              // isWail
}

// newRoute for return route
func newRoute() *route {
	r := &route{index: -1, routes: make(map[string]*route)}
	return r
}

// RegisterRouter for add router
func (r *route) RegisterRouter(url string, handler Handler) {
	urlList := parseUrl(url)
	recursionRegisterRouter(r, urlList, handler)
}

func (r *route) Router(url string) Handler {
	urlList := parseUrl(url)
	var ok bool
	var router *route
	router = r

	for _, url := range urlList {
		router, ok = router.routes[url]
		if !ok {
			wlog.Error.Printf(" routes map not exculde: %s", url)
		}
	}
	if router.isHandle {
		return router.handler
	}
	return nil
}

func recursionRegisterRouter(r *route, urls []string, handler Handler) {
	if r.isWail {
		return
	} else if r.index >= len(urls)-1 {
		r.handler = handler
		r.isHandle = true
		return
	}

	index := r.index
	index++
	url := urls[index]
	var router *route
	var ok bool
	router, ok = r.routes[url]
	if !ok {
		router = newRoute()
		r.routes[url] = router
	}
	router.index = index
	router.prefix = url

	recursionRegisterRouter(router, urls, handler)

}
