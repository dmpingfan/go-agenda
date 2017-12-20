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
+ [f58fd87] 完成 Dockerfile
+ [d9984c9] 完成 Dockerfile 和 travis
+ [8df3cd5] merge cli
+ [c6be716] merge cli
+ [f1a3eea] cli bug
+ [76c58b1] 完成
+ [7c17f61] 删除本地 db 文件
+ [8885cf7] 添加函数的注释
+ [c8583be] README

## 项目小结

通过本次 go-agenda 项目，我的确接触和学习到了许多关于后端开发的知识。

+ Docker 部分。本次项目是将 cli 和 service 两个程序部署在一个 Docker 镜像中。
  + 对于 Go + Docker 项目的开发，应该将项目放在镜像的 $GOPATH/src/ 目录下，不然编译时会出错。
  + Docker 中 VOLUME 是一个很重要的概念，用于同步本地路径与 Docker 容器路径。需要在 Dockerfile 文件中声明 VOLUME 路径，然后在 run 的时候再指定具体的路径。这样一来，容器中凡是存储在 VOLUME 目录下的文件都会相应地存在与本地的指定路径，有点类型与 Linux 系统中的 mount。
  + 