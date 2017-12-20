# go-agenda

[![Build Status](https://travis-ci.org/painterdrown/go-agenda.svg?branch=master)](https://travis-ci.org/painterdrown/go-agenda.svg?branch=master)

Agenda based on Go, with CLI and service.

## 生成 Docker 镜像

```
docker build -t go-agenda .
```

## Docker 使用方法

  + 下载镜像
```
docker pull painterdrown/go-agenda
```

  + 启动 service 服务器
```
docker run --name go-agenda-sevice -d -v $DATA_PATH:/data -p 3000:3000 go-agenda service
```
> 比如：docker run --name go-agenda-sevice -d -v ~:/data -p 3000:3000 go-agenda service
> 成功后访问 localhost:3000 来使用 go-agenda-service。

  + 运行 cli 客户端
```
docker run --name go-agenda-cli -v $DATA_PATH:/data go-agenda cli [COMMAND] [ARG...]
```
> 比如：docker run --name go-agenda-cli -v ~:/data go-agenda cli --help

其中，$DATA_PATH 是你本地对某个路径，用来存储 go-agenda 的数据库文件。
