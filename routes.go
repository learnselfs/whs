// Package whs @Author Bing
// @Date 2024/2/3 20:37:00
// @Desc
package whs

type Handler func(ctx *Content)

// route struct
type route struct {
	*route
	index   int
	prefix  string
	handler Handler
}
