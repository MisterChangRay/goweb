# goweb

一个简单的go web脚手架

go版本: go version go1.22.5 windows/386


启动命令:
`go run .\main.go .\router.go`

编译命令:
`go build .\main.go .\router.go `


编译压缩命令
`go build -ldflags="-s -w" -o server main.go && upx -9 server`


使用方式
1. 导入conf 目录下的sql文件
2. 配置conf目录下的app.conf,修改数据库连接信息
3. 使用启动命令进入快速开发阶段