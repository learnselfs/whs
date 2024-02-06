// Package whs @Author Bing
// @Date 2024/2/3 20:11:00
// @Desc
package whs

import (
	"context"
	"github.com/learnselfs/wlog"
	"net/http"
	"time"
)

// Service struct
type Service struct {
	host  string
	port  int
	close chan struct{}
	*Route
	*http.Server
}

// ServeHTTP for main processing function
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContent(r, w)
	c.middlewares = s.Router(c.RequestURI)
	c.param = s.Route.param
	if len(c.middlewares) > 0 {
		c.Next()
	} else {
		NotFoundHandler(c)
	}
}

// Start for http server
func (s *Service) Start() {
	wlog.Info.Println(s.ListenAndServe())
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
