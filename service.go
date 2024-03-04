// Package whs @Author Bing
// @Date 2024/2/3 20:11:00
// @Desc
package whs

import (
	"context"
	"github.com/learnselfs/wlog"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// Service struct
type Service struct {
	host  string
	port  int
	close chan struct{}
	*Route
	*http.Server
	template *template.Template
}

// ServeHTTP for main processing function
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContent(r, w)
	c.template = s.template // http template
	c.middlewares = s.Router(c.RequestURI)
	c.param = s.Route.param // route parameters
	c.template = s.template
	// logger

	if len(c.middlewares) > 0 {
		c.Next()
	} else {
		NotFoundHandler(c)
	}
}

// Start for http server
func (s *Service) Start() {
	wlog.Info("Listening on " + s.host + ":" + strconv.Itoa(s.port))
	s.ListenAndServe()
}

// Stop for http server
func (s *Service) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := s.Shutdown(ctx)
	if err != nil {
		return
	}

}

func (s *Service) Static(url, path string) {
	h := fileServer(url, path)
	s.RegisterRouter(url+"*", h)
}

func (s *Service) Template(webPah string) {
	s.template = template.Must(s.template.ParseGlob(webPah))
}

func (s *Service) TemplateDelim(left, right, webPah string) {
	s.template = template.Must(s.template.Delims(left, right).ParseGlob(webPah))
}

func (s *Service) Func(fun template.FuncMap) {
	s.template = s.template.Funcs(fun)
}

func (s *Service) preloadMiddleware() {
	s.UseMiddleware(accessLogHandler)
}
