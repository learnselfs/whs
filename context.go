// Package whs @Author Bing
// @Date 2024/2/3 21:01:00
// @Desc
package whs

import (
	"encoding/json"
	"html/template"
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
	state       int
	// html templates
	template *template.Template
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

func (c *Context) setState(state int) {
	c.state = state
	c.ResponseWriter.WriteHeader(state)
}

func (c *Context) setHeader(key string, value string) {
	c.ResponseWriter.Header().Set(key, value)
}

func (c *Context) Html(state int, tmpl string, data map[string]any) {
	c.setHeader("Content-Type", "text/html")
	c.setState(state)
	err := c.template.ExecuteTemplate(c.ResponseWriter, tmpl, data)
	if err != nil {
		return
	}
}

func (c *Context) Json(state int, i any) {
	c.setHeader("Content-Type", "application/json")
	c.setState(state)
	j := json.NewEncoder(c.ResponseWriter)
	err := j.Encode(i)
	if err != nil {
		return
	}
}
