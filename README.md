# goweb

一个简单的go web脚手架,用于快速开发一些简单的web接口项目。由于可以快速打包到各个平台。比较适合一些开发不交付源码的。

go版本: go version go1.22.5 windows/386

### 1.使用方式
1. 导入conf 目录下的sql文件
2. 配置conf目录下的app.conf,修改数据库连接信息
3. 使用启动命令进入快速开发阶段
4. 打包异地运行时，需要拷贝配置文件目录

### 2. 常用命令
如果需要压缩，需要提前到github下载upx并添加到path。

启动命令:
`go run .\main.go .\router.go`

编译命令:
`go build .\main.go .\router.go `


编译压缩命令
`go build -ldflags="-s -w" -o server.exe main.go && upx -9 server.exe`


交叉编译
```
# 编译linux执行程序
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o server_linux.exe main.go && upx -9 server_linux.exe


# 编译windows执行程序
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o server_win.exe main.go && upx -9 server_win.exe
```

