// Package whs @Author Bing
// @Date 2024/2/10 20:46:00
// @Desc
package whs

import (
	"net/http"
	"os"
	"path/filepath"
)

func fileServer(url, Path string) Handler {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, Path)
	h := http.StripPrefix(url, http.FileServer(http.Dir(path)))
	return func(c *Context) {
		h.ServeHTTP(c.ResponseWriter, c.Request)
	}
}
