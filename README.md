# whs

## 介绍
web http service

## 安装教程
`go get github.com/learnselfs/whs`

## 使用说明
- 初始化

```go
s := New("127.0.0.1", 80)
```

- 路由
```go
s.RegisterRouter("/user", func(c *Context) {
    c.ResponseWriter.Write([]byte("/user"))
}
```
- 路由组

```go
home := s.Group("/home")
{
    home.RegisterRouter("/index", func(c *Context) { c.ResponseWriter.Write([]byte("/home/index")) })
}}
)


```
- 参数获取

- 中间件

```go
home.UseMiddleware(func(c *Context) {
    c.ResponseWriter.Write([]byte("2"))
    c.Next()
    c.ResponseWriter.Write([]byte("2"))
})
```
- 启动
```go
s.Start()
```
- 停止
```go
s.Stop()
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
