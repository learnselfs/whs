// Package whs @Author Bing
// @Date 2024/2/4 20:58:00
// @Desc
package whs

func NotFoundHandler(c *Context) {
	_, err := c.ResponseWriter.Write([]byte("not found 404"))
	if err != nil {
		return
	}
}
