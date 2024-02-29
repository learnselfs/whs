
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
  <a href="https://github.com/learnselfs/wlog/">
    <img src="logo.png" alt="Logo"  height="80">
  </a>

<h3 align="center"></h3>
  <p align="center">
wlog is a structured logger for Go (golang), completely API compatible with the standard library logger.
    <br />
    <a href="https://github.com//learnselfs/wlog"><strong> markdown »</strong></a>
    <br />
    <br />
    <a href="https://github.com//learnselfs/wlog">Demo</a>
    ·
    <a href="https://github.com//learnselfs/wlog/issues">Bug</a>
    ·
    <a href="https://github.com//learnselfs/wlog/issues">Issues</a>
  </p>

</p>

English | [中文](./README_ch.md)
### Guide
1. create server 
```go
s := New("127.0.0.1", 80)
```
2. routes
````go
s.RegisterRouter("/user", func(c *Context) {
c.ResponseWriter.Write([]byte("/user"))
}
````
3. route group
```go
home := s.Group("/home")
{
    home.RegisterRouter("/index", func(c *Context) { c.ResponseWriter.Write([]byte("/home/index")) })
}}
)
```
4. middleware
```go
home.UseMiddleware(func(c *Context) {
c.ResponseWriter.Write([]byte("2"))
c.Next()
c.ResponseWriter.Write([]byte("2"))
})
```
5. template
```go
s.Static("static/", "/static")
s.Func(template.FuncMap{"FormatTime": FormatTime})
s.Template("template/*")```
6. start and stop 
```text
s.Start()
s.stop()
```
###### Pre development Configuration Requirements

1. go version 1.21.1

###### **Installation**

1. `go get github.com/learnselfs/wlog`
   [github.com/learnselfs/wlog](https://pkg.go.dev/github.com/learnselfs/wlog)

### Contributor
Please read **CONTABUTING.md** to find out the developers who have contributed to this project.


#### Open Source Projects

Contributing makes the open source community an excellent place to learn, motivate, and create.
Any contribution you make is greatly appreciated.


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### License

This project has signed an Apache license, please refer to for details.
[LICENSE](https://github.com/learnselfs/wlog/blob/master/LICENSE)

### Thanks


- [gin](https://github.com/gin-gonic/gin)
- [Best_README_template](https://github.com/shaojintian/Best_README_template)
- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)

<!-- links -->
[your-project-path]:/learnselfs/wlog
[contributors-shield]: https://img.shields.io/github/contributors/learnselfs/wlog.svg?style=flat-square
[contributors-url]: https://github.com//learnselfs/wlog/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks//learnselfs/wlog.svg?style=flat-square
[forks-url]: https://github.com/learnselfs/wlog/network/members
[stars-shield]: https://img.shields.io/github/stars//learnselfs/wlog.svg?style=flat-square
[stars-url]: https://github.com//learnselfs/wlog/stargazers
[issues-shield]: https://img.shields.io/github/issues/learnselfs/wlog.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues//learnselfs/wlog.svg
[license-shield]: https://img.shields.io/github/license//learnselfs/wlog.svg?style=flat-square
[license-url]: https://github.com/learnselfs/wlog/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian