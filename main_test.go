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
	"time"
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
	wlog.Info(str.String())
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
	r.RegisterRouter("GET", "/user", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user"))
	})
	r.RegisterRouter("GET", "/user/admin", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user/admin"))
	})

	r.Router("GET", "/user")
	r.Router("GET", "/user/admin")
}

func serverStart() *Service {

	s := New("127.0.0.1", 80)
	s.GET("/user", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user"))
	})
	s.GET("/user/admin", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user/admin"))
	})

	s.GET("/users/*", func(c *Context) {
		c.ResponseWriter.Write([]byte("/users"))
	})
	s.GET("/user/:user/info", func(c *Context) {
		c.ResponseWriter.Write([]byte("/user/" + c.param["user"] + "/info"))
	})

	//go func() {
	//	s.Start()
	//}()

	return s
}

func ClientGet(t *testing.T, url string, result string, method ...string) {
	c := &http.Client{}
	var res1 *http.Response
	if len(method) == 0 {
		method = []string{"GET"}
	}
	switch method[0] {
	case "POST":
		res1, _ = c.Post("http://127.0.0.1"+url, "application/json", nil)
	default:
		res1, _ = c.Get("http://127.0.0.1" + url)
	}
	b1, err := io.ReadAll(res1.Body)
	if err != nil {
		wlog.Error(err.Error())
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
	wlog.Info(string(b1))

	res2, _ := c.Get("http://127.0.0.1/user")
	b2 := make([]byte, 20)
	res2.Body.Read(b2)
	wlog.Info(string(b2))

	s.Stop()
}

func TestUrlParams(t *testing.T) {
	s := serverStart()
	go s.Start()
	time.Sleep(time.Millisecond * 10)
	t.Run("name", func(t *testing.T) {
		ClientGet(t, "/user/admins/info", "/user/admins/info")
		ClientGet(t, "/user/a/info", "/user/a/info")
		ClientGet(t, "/user/info/info", "/user/info/info")
		ClientGet(t, "/users/admina/nfo", "/users")
		ClientGet(t, "/users/adminaaa/info", "/users")

	})
	s.Stop()
}

func TestRouteGroup(t *testing.T) {
	s := serverStart()

	home := s.Group("/home")
	{
		home.GET("/info", func(c *Context) { c.ResponseWriter.Write([]byte("/home/info")) })
		home.GET("/*", func(c *Context) { c.ResponseWriter.Write([]byte("/home/*")) })
	}

	admin := s.Group("admin")
	{
		admin.GET("info", func(c *Context) { c.ResponseWriter.Write([]byte("/admin/info")) })
		admin.GET("/:info", func(c *Context) { c.ResponseWriter.Write([]byte("/admin/" + c.param["info"])) })
		admin.POST("/:info", func(c *Context) { c.ResponseWriter.Write([]byte("post/admin/" + c.param["info"])) })
	}
	go func() {
		s.Start()
	}()
	time.Sleep(10 * time.Millisecond)
	ClientGet(t, "/home/info", "/home/info")
	ClientGet(t, "/home/info*", "/home/*")
	ClientGet(t, "/admin/info", "/admin/info")
	ClientGet(t, "/admin/infos", "/admin/infos")
	ClientGet(t, "/admin/infos00", "/admin/infos", "POST")
	s.Stop()
}
func middle(count string) Handler {
	return func(c *Context) {
		c.ResponseWriter.Write([]byte(count))
		c.Next()
		c.ResponseWriter.Write([]byte(count))
	}
}
func TestMiddleware(t *testing.T) {
	s := serverStart()
	s.UseMiddleware(middle("1"))
	home := s.Group("/home")
	home.UseMiddleware(middle("2"))
	{
		home.GET("/info", func(c *Context) { c.ResponseWriter.Write([]byte("/home/info")) })
		home.GET("/*", func(c *Context) { c.ResponseWriter.Write([]byte("/home/*")) })
	}

	admin := s.Group("admin")
	admin.UseMiddleware(middle("3"))
	{
		admin.GET("info", func(c *Context) { c.ResponseWriter.Write([]byte("/admin/info")) })
		admin.GET("/:info", func(c *Context) { c.ResponseWriter.Write([]byte("/admin/" + c.param["info"])) })
	}
	go s.Start()
	time.Sleep(time.Millisecond * 10)
	ClientGet(t, "/home/info", "12/home/info21")
	ClientGet(t, "/home/info1*", "12/home/*21")
	ClientGet(t, "/admin/info", "13/admin/info31")
	ClientGet(t, "/admin/infos", "13/admin/infos31")
	s.Stop()
	//
}

func TestFileServer(t *testing.T) {
	s := serverStart()
	home := s.Group("/home")
	home.UseMiddleware(func(c *Context) {
		c.ResponseWriter.Write([]byte("2"))
		c.Next()
		c.ResponseWriter.Write([]byte("2"))
	})
	{
		home.GET("/info", func(c *Context) { c.ResponseWriter.Write([]byte("/home/info")) })
		home.GET("/*", func(c *Context) { c.ResponseWriter.Write([]byte("/home/*")) })
	}
	s.Static("static", "/.gitee")
	s.Start()

	//client := &http.Client{}
	//res, _ := client.Get("http://127.0.0.1/.gitee/PULL_REQUEST_TEMPLATE.zh-CN.md")
	//if res.StatusCode != 200 {
	//	t.Errorf("test static failed")
	//}
	//s.Stop()
}

func FormatTime(t time.Time) string {
	return t.AddDate(100, 100, 1).Format(time.DateTime)
}
func TestTemplate(t *testing.T) {
	s := serverStart()
	s.Static("/static/", "./.gitee/temp")
	//s.Func(template.FuncMap{"FormatTime": FormatTime})
	//s.Template("*gitee/*")

	home := s.Group("home/")
	home.UseMiddleware(func(c *Context) {
		//c.ResponseWriter.Write([]byte("2"))
		c.Next()
		//c.ResponseWriter.Write([]byte("2"))
	})
	//{
	//	home.RegisterRouter("/info", func(c *Context) {
	//		c.Html(200, "base.tmpl", map[string]any{"name": "小明", "gender": true, "age": 18, "time": time.Now()})
	//	})
	//	home.RegisterRouter("/*", func(c *Context) { c.ResponseWriter.Write([]byte("/home/*")) })
	//}

	s.Start()

	//client := &http.Client{}
	//res, _ := client.Get("http://127.0.0.1/.gitee/PULL_REQUEST_TEMPLATE.zh-CN.md")
	//if res.StatusCode != 200 {
	//	t.Errorf("test static failed")
	//}
	//s.Stop()
}
