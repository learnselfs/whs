// Package whs @Author Bing
// @Date 2024/2/3 21:01:00
// @Desc
package whs

import (
	"net/http"
)

// Content for encapsulating request data and processing responses
type Content struct {
	// request
	*http.Request

	// response
	http.ResponseWriter
	param map[string]string
}

// NewContent returns a new Content
func NewContent(r *http.Request, w http.ResponseWriter) *Content {
	c := &Content{Request: r, ResponseWriter: w, param: make(map[string]string)}
	return c
}

func (c *Content) parse() {

}
