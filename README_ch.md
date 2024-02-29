<!-- PROJECT SHIELDS -->

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />

<p align="center">
  <a href="https://github.com/learnselfs/whs/">
    <img src="logo.png" alt="Logo"  height="80">
  </a>

<h3 align="center"></h3>
  <p align="center">
web http(s) service
    <br />
    <a href="https://github.com//learnselfs/whs"><strong>探索本项目的文档 »</strong></a>
    <br />
    <br />
    <a href="https://github.com//learnselfs/whs">查看Demo</a>
    ·
    <a href="https://github.com//learnselfs/whs/issues">报告Bug</a>
    ·
    <a href="https://github.com//learnselfs/whs/issues">提出新特性</a>
  </p>

</p>

[English](./README.md) | 中文
## 目录

- [上手指南](#上手指南)
    - [开发前的配置要求](#开发前的配置要求)
    - [安装步骤](#安装步骤)
- [部署](#部署)
- [使用到的框架](#使用到的框架)
- [贡献者](#贡献者)
    - [如何参与开源项目](#如何参与开源项目)
- [版本控制](#版本控制)
- [鸣谢](#鸣谢)

### 上手指南
1. 创建服务 
```go
s := New("127.0.0.1", 80)
```
2. 路由 
````go
s.RegisterRouter("/user", func(c *Context) {
c.ResponseWriter.Write([]byte("/user"))
}
````
3. 路由组
```go
home := s.Group("/home")
{
    home.RegisterRouter("/index", func(c *Context) { c.ResponseWriter.Write([]byte("/home/index")) })
}}
)
```
4. 中间件 
```go
home.UseMiddleware(func(c *Context) {
c.ResponseWriter.Write([]byte("2"))
c.Next()
c.ResponseWriter.Write([]byte("2"))
})
```
5. 模板
```go
s.Static("static/", "/static")
s.Func(template.FuncMap{"FormatTime": FormatTime})
s.Template("template/*")```
6. start and stop 
```text
s.Start()
s.stop()
```

###### 开发前的配置要求

1. go version 1.21.1

###### **安装步骤**

1. `go get github.com/learnselfs/whs`
   [github.com/learnselfs/whs](https://pkg.go.dev/github.com/learnselfs/wlog)

### 贡献者

请阅读**CONTRIBUTING.md** 查阅为该项目做出贡献的开发者。

#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

*您也可以在贡献者名单中参看所有参与该项目的开发者。*

### 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE](https://github.com//learnselfs/whs/blob/master/LICENSE)

### 鸣谢


- [gin](https://github.com/gin-gonic/gin)
- [Best_README_template](https://github.com/shaojintian/Best_README_template)
- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)

<!-- links -->
[your-project-path]:/learnselfs/whs
[contributors-shield]: https://img.shields.io/github/contributors/learnselfs/whs.svg?style=flat-square
[contributors-url]: https://github.com//learnselfs/whs/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks//learnselfs/whs.svg?style=flat-square
[forks-url]: https://github.com/learnselfs/whs/network/members
[stars-shield]: https://img.shields.io/github/stars//learnselfs/whs.svg?style=flat-square
[stars-url]: https://github.com//learnselfs/whs/stargazers
[issues-shield]: https://img.shields.io/github/issues/learnselfs/whs.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues//learnselfs/whs.svg
[license-shield]: https://img.shields.io/github/license//learnselfs/whs.svg?style=flat-square
[license-url]: https://github.com/learnselfs/whs/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian
