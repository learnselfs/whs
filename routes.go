// Package whs @Author Bing
// @Date 2024/2/3 20:37:00
// @Desc
package whs

import (
	"github.com/learnselfs/wlog"
	"strings"
)

type Handler func(ctx *Content)

// route struct
type route struct {
	routes     map[string]*route // routes
	index      int               // index of route
	prefix     string            // prefix is the route url
	handler    Handler           // handler is processing the route
	isHandle   bool              // isHandle
	isAsterisk bool              // isAsterisk matching
	isColon    bool              // isColon matching
	param      map[string]string // params

}

// newRoute for return route
func newRoute() *route {
	r := &route{index: -1, routes: make(map[string]*route), param: make(map[string]string)}
	return r
}

// RegisterRouter for add router
func (r *route) RegisterRouter(url string, handler Handler) {
	urlList := parseUrl(url)
	recursionRegisterRouter(r, urlList, handler)
}

func (r *route) Router(url string) Handler {
	urlList := parseUrl(url)
	var router *route
	router = r
	router = recursionRouter(router, urlList, r.param)
	wlog.Debug.Printf("%#v", router)
	if router != nil && router.isHandle {
		return router.handler
	}
	return nil
}

func recursionRegisterRouter(r *route, urls []string, handler Handler) {
	if r.isAsterisk {
		r.handler = handler
		r.isHandle = true
		return
	} else if r.index >= len(urls)-1 {
		r.handler = handler
		r.isHandle = true
		return
	}
	var isColon bool
	var isAsterisk bool
	index := r.index
	index++

	url := urls[index]
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

	recursionRegisterRouter(router, urls, handler)

}

func recursionRouter(r *route, urls []string, param map[string]string) *route {
	if r.isAsterisk {
		return r
	} else if r.isHandle && r.index >= len(urls)-1 {
		wlog.Debug.Printf("%#v", r)
		return r
	}

	url := urls[r.index+1]
	router, ok := r.routes[url]
	if ok {
		wlog.Debug.Printf("---%#v", router)
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
