# go-agenda

Agenda based on Go, with CLI and service.

## Docker 使用方法

  + 下载镜像
```
docker pull painterdrown/go-agenda
```

  + 启动 service 服务器
```
docker run --name go-agenda-sevice -d -v $DATA_PATH:/data -p 8080:8080 painterdrown/go-agenda service
```

  + 运行命令行客户端
```
docker run --rm --network host -v $DATA_PATH:/data painterdrown/go-agenda cli help
```