## go module使用
- 环境变量设置
> $ go env -w GO111MODULE=on   //开启module
> $ go env -w GOPROXY=https://goproxy.cn,direct  //更换包下载代理
> go env -w GOSUMDB=off //包校验关闭，因为默认的校验地址被墙了

- 新建工程
> 保证新建的工程不在$GOPATH下

- module初始化
> 1.带git的项目： ~~直接从github上面clone一个项目下来。直接执行 go mod init 会自动生成带git地址的packagename~~  和2一样
> 2.不带git的项目：$go mod init PACKAGENAME