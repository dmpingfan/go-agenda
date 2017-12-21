# painterdrown in go-agenda

## fork 项目的位置

https://github.com/painterdrown/go-agenda

## 个人工作摘要

+ [x] 在 Github 上创建 go-agenda 仓库，并搭好项目框架
+ [x] 使用 API Blueprint 写好 API 文档
+ [x] 负责数据库部分，即 entities 目录
+ [x] 负责 Docker 镜像的制作和生成，成功后 push 到 Docker Hub
+ [x] 负责 travis CI 连续集成
+ [x] 编写 _test.go 文件并进行测试
+ [x] 负责 README 文档，相关文档的说明

## 个人 commit 记录

+ [1f37191](https://github.com/painterdrown/go-agenda/commit/1f37191) init
+ [f58fd87](https://github.com/painterdrown/go-agenda/commit/f58fd87) 完成 Dockerfile
+ [d9984c9](https://github.com/painterdrown/go-agenda/commit/d9984c9) 完成 Dockerfile 和 travis
+ [8df3cd5](https://github.com/painterdrown/go-agenda/commit/8df3cd5) merge cli
+ [c6be716](https://github.com/painterdrown/go-agenda/commit/c6be716) merge cli
+ [f1a3eea](https://github.com/painterdrown/go-agenda/commit/f1a3eea) cli bug
+ [76c58b1](https://github.com/painterdrown/go-agenda/commit/76c58b1) 完成
+ [7c17f61](https://github.com/painterdrown/go-agenda/commit/7c17f61) 删除本地 db 文件
+ [8885cf7](https://github.com/painterdrown/go-agenda/commit/8885cf7) 添加函数的注释
+ [c8583be](https://github.com/painterdrown/go-agenda/commit/c8583be) README

## 项目小结

通过本次 go-agenda 项目，我的确接触和学习到了许多关于后端开发的知识

+ API 设计部分。本次项目要求使用 RESTful 规范进行 API 的设计。除了登陆这个 API 之外，其他都能很好的根据服务资源的 CRUD 操作进行设计。但是如果将“登录”的 API 设计为：`POST users/login`，则不太符合 RESTful 规范，因为不该在 URL 中使用动词，动作应该通过 HTTP 方法来体现。Google 一番之后，觉得目前最合理的“登录” API 设计是：`POST sessions`。理解：登录操作会在服务器创建一个 session 的资源。

+ Docker 部分。本次项目是将 cli 和 service 两个程序部署在一个 Docker 镜像中。
  + 对于 Go + Docker 项目的开发，应该将项目放在镜像的 $GOPATH/src/ 目录下，不然编译时会出错。
  + Docker 中 VOLUME 是一个很重要的概念，用于同步本地路径与 Docker 容器路径。需要在 Dockerfile 文件中声明 VOLUME 路径，然后在 run 的时候再指定具体的路径。这样一来，容器中凡是存储在 VOLUME 目录下的文件都会相应地存在与本地的指定路径，有点类型与 Linux 系统中的 mount。

+ Travis CI 部分。我是第一次接触到“持续集成”这个概念。阅读官方文档后了解到持续集成是用于小规模代码修改的项目的集成测试，注意下这里指的是“小规模”，也就是说，Travis CI 最大的使用价值在于保证我们每次 Github 上的新的 commit（非大规模代码修改）都能成功部署。