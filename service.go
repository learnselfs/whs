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
	*route
	*http.Server
}

// ServeHTTP for main processing function
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContent(r, w)
	wlog.Info.Printf("%#v", c.Request)
	c.ResponseWriter.Write([]byte(c.RequestURI))

}

// Start for http server
func (s *Service) Start() {
	err := http.ListenAndServe(s.Addr, s)
	if err != nil {
		wlog.Error.Println(err)
		return
	}
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
