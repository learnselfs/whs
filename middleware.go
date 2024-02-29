// Package whs @Author Bing
// @Date 2024/2/4 20:58:00
// @Desc
package whs

import "net/http"

func NotFoundHandler(c *Context) {
	c.setState(404)
	accessLog.Errorf("[method]: %s,\t[url]: %#v,\t[state]: %d(%s), [remote]:%s", c.Request.Method, c.RequestURI, c.state, http.StatusText(c.state), c.RemoteAddr)
	_, err := c.ResponseWriter.Write([]byte("not found 404"))
	if err != nil {
		return
	}
}

func accessLogHandler(c *Context) {
	c.Next()
	if c.state == 200 || c.state == 0 {
		accessLog.Printf("[method]: %s,\t[url]: %#v,\t[remote]:%s", c.Request.Method, c.RequestURI, c.Request.RemoteAddr)
	} else {
		accessLog.Errorf("[method]: %s,\t[url]: %#v,\t[state]: %d(%s), [remote]:%s", c.Request.Method, c.RequestURI, c.state, http.StatusText(c.state), c.RemoteAddr)
	}
}
