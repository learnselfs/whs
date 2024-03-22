// Package whs @Author Bing
// @Date 2024/2/4 20:58:00
// @Desc
package whs

import (
	"errors"
	"net/http"
)

func NotFoundHandler(c *Context) {
	ErrorHandler(c, http.StatusNotFound, errors.New("not found 404"))
}

func ErrorHandler(c *Context, status int, err error) {
	c.setState(status)
	accessLog.Errorf("[method]: %s,\t[url]: %#v,\t[state]: %d(%s), [remote]:%s", c.Request.Method, c.RequestURI, c.state, http.StatusText(c.state), c.RemoteAddr)
	c.ResponseWriter.Write([]byte(err.Error()))
	return
}

func accessLogHandler(c *Context) {
	c.Next()
	if c.state == 200 || c.state == 0 {
		accessLog.Infof("[method]: %s,\t[url]: %#v,\t[remote]:%s", c.Request.Method, c.RequestURI, c.Request.RemoteAddr)
	} else {
		accessLog.Errorf("[method]: %s,\t[url]: %#v,\t[state]: %d(%s), [remote]:%s", c.Request.Method, c.RequestURI, c.state, http.StatusText(c.state), c.RemoteAddr)
	}
}
