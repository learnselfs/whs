// Package whs @Author Bing
// @Date 2024/2/10 20:46:00
// @Desc
package whs

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var htmlFile string

func Html(webPath string) (*template.Template, error) {
	return template.New(webPath).ParseFiles(htmlFile)
}

func fileServer(webPath string) Handler {
	wd, _ := os.Getwd()
	path := filepath.Join(wd, webPath)
	h := http.StripPrefix(webPath, http.FileServer(http.Dir(path)))
	return func(c *Context) {
		h.ServeHTTP(c.ResponseWriter, c.Request)
	}
}
