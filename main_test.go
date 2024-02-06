// Package whs @Author Bing
// @Date 2024/2/3 22:02:00
// @Desc
package whs

import (
	"bytes"
	"github.com/learnselfs/wlog"
	"io"
	"net/http"
	"strconv"
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

func TestParserUrl(t *testing.T) {
	urlTests := []struct {
		url    string
		result []string
	}{
		{"/user", []string{"user"}},
		{"/user/admin", []string{"user", "admin"}},
	}

	for i, test := range urlTests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			result := parseUrl(test.url)
			if !(len(result) == len(test.result)) {
				t.Errorf("[result] %s: %d, [test.result] %s: %d)", test.result, len(test.result), result, len(result))
			}
		})
	}
}

func TestRoutes(t *testing.T) {
	r := newRoute()
	r.RegisterRouter("/user", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user"))
	})
	r.RegisterRouter("/user/admin", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user/admin"))
	})

	r.Router("/user")
	r.Router("/user/admin")
}

func serverStart() *Service {

	s := New("127.0.0.1", 80)
	s.RegisterRouter("/user", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user"))
	})
	s.RegisterRouter("/user/admin", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user/admin"))
	})

	s.RegisterRouter("/users/*", func(c *Context) {
		c.ResponseWriter.Write([]byte("/users"))
	})
	s.RegisterRouter("/user/:user/info", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user/" + c.param["user"] + "/info"))
	})

	//go func() {
	//	s.Start()
	//}()

	return s
}

func ClientGet(t *testing.T, url string, result string) {
	c := &http.Client{}
	res1, _ := c.Get("http://127.0.0.1" + url)
	b1, err := io.ReadAll(res1.Body)
	if err != nil {
		wlog.Error.Println(err)
		t.Error(err)
		return
	}
	t.Run(string(b1), func(t *testing.T) {
		if string(b1) != result {
			t.Errorf("\n\tExpected result: %s, \n\tActual result: %s", result, b1)
		}
	})
}

func TestBaseRoutes(t *testing.T) {
	s := serverStart()
	c := http.Client{}
	res1, _ := c.Get("http://127.0.0.1/user/admin")
	b1 := make([]byte, 20)
	res1.Body.Read(b1)
	wlog.Info.Println(string(b1))

	res2, _ := c.Get("http://127.0.0.1/user")
	b2 := make([]byte, 20)
	res2.Body.Read(b2)
	wlog.Info.Println(string(b2))

	s.Stop()
}

func TestUrlParams(t *testing.T) {
	s := serverStart()
	t.Run("name", func(t *testing.T) {
		ClientGet(t, "/user/admins/info", "/user/admins/info")
		ClientGet(t, "/user/adminaaa/info", "/user/adminaaa/info")
		ClientGet(t, "/users/adminaaa/info", "/users/adminaa/info")
		ClientGet(t, "/users/admina/nfo", "/users/admina/nfo")
		ClientGet(t, "/use/nfo", "")

	})
	s.Stop()
	wlog.Info.Println(s)
}

func TestRouteGroup(t *testing.T) {
	s := serverStart()

	home := s.Group("/home")
	{
		home.RegisterRouter("/info", func(c *Context) { c.ResponseWriter.Write([]byte("/home/info")) })
		home.RegisterRouter("/*", func(c *Context) { c.ResponseWriter.Write([]byte("/home/*")) })
	}

	admin := s.Group("admin")
	{
		admin.RegisterRouter("info", func(c *Context) { c.ResponseWriter.Write([]byte("/admin/info")) })
		admin.RegisterRouter("/:info", func(c *Context) { c.ResponseWriter.Write([]byte("/admin/" + c.param["info"])) })
	}
	go func() {
		s.Start()
	}()

	ClientGet(t, "/home/info", "/home/info")
	ClientGet(t, "/home/info*", "/home/info*")
	ClientGet(t, "/admin/info", "/admin/info")
	ClientGet(t, "/admin/infos", "/admin/infos")
	s.Stop()
}

func TestMiddleware(t *testing.T) {
	s := serverStart()
	s.UseMiddleware(func(c *Context) {
		c.ResponseWriter.Write([]byte("1"))
		c.Next()
		c.ResponseWriter.Write([]byte("1"))
	})
	home := s.Group("/home")
	home.UseMiddleware(func(c *Context) {
		c.ResponseWriter.Write([]byte("2"))
		c.Next()
		c.ResponseWriter.Write([]byte("2"))
	})
	{
		home.RegisterRouter("/info", func(c *Context) { c.ResponseWriter.Write([]byte("/home/info")) })
		home.RegisterRouter("/*", func(c *Context) { c.ResponseWriter.Write([]byte("/home/*")) })
	}

	admin := s.Group("admin")
	admin.UseMiddleware(func(c *Context) {
		c.ResponseWriter.Write([]byte("3"))
		c.Next()
		c.ResponseWriter.Write([]byte("3"))
	})
	{
		admin.RegisterRouter("info", func(c *Context) { c.ResponseWriter.Write([]byte("/admin/info")) })
		admin.RegisterRouter("/:info", func(c *Context) { c.ResponseWriter.Write([]byte("/admin/" + c.param["info"])) })
	}
	go func() {
		s.Start()
	}()

	ClientGet(t, "/home/info", "12/home/info21")
	ClientGet(t, "/home/info1*", "12/home/*21")
	ClientGet(t, "/admin/info", "13/admin/info31")
	ClientGet(t, "/admin/infos", "13/admin/infos31")
	s.Stop()

}
