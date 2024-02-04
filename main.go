// Package whs @Author Bing
// @Date 2024/2/2 17:07:00
// @Desc
package whs

import (
	"fmt"
	"net/http"
	"time"
)

// New function return http server
func New(host string, port int) *Service {
	return &Service{host: host, port: port, Server: &http.Server{Addr: fmt.Sprintf("%s:%d", host, port), ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}}
}
