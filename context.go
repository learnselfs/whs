// Package whs @Author Bing
// @Date 2024/2/3 21:01:00
// @Desc
package whs

import (
	"net/http"
)

// Context for encapsulating request data and processing responses
type Context struct {
	// request
	*http.Request

	// response
	http.ResponseWriter
	param       map[string]string
	middlewares []Handler
	index       int
}

// NewContent returns a new Context
func NewContent(r *http.Request, w http.ResponseWriter) *Context {
	c := &Context{Request: r, ResponseWriter: w, param: make(map[string]string), middlewares: make([]Handler, 0), index: -1}
	return c
}

func (c *Context) parse() {

}

// Next complete middleware
func (c *Context) Next() {
	c.index++
	c.middlewares[c.index](c)
}
