// Package whs @Author Bing
// @Date 2024/2/3 22:02:00
// @Desc
package whs

import (
	"bytes"
	"github.com/learnselfs/wlog"
	"net/http"
	"testing"
)

func TestBaseService(t *testing.T) {
	s := New("127.0.0.1", 80)
	go func() {
		s.Start()
	}()

	c := &http.Client{}
	response, err := c.Get("http://127.0.0.1/123a")
	if err != nil {
		return
	}
	var url []byte
	response.Body.Read(url)
	var str bytes.Buffer
	str.Write(url)
	wlog.Info.Println(str)
	if str.String() == "/123a" {
		t.Log("ok!!!")
	}
	s.Stop()
}
