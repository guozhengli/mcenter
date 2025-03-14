# mcenter

基于Etcd的微服务注册中心

1. 配置MongoDB
```
use mcenter
db.createUser({user: "mcenter", pwd: "123456", roles: [{ role: "dbOwner", db: "mcenter" }]})
```

2. 初始化内置服务
```
make init

2022-06-11T17:36:23.304+0800    INFO    [INIT]  cmd/start.go:153        log level: debug
2022-06-11T17:36:23.356+0800    INFO    [init]  cmd/init.go:50  init app keyauth success, client_id: 4x6O2CgbiySkgswh7Ans2qCj, client_secret: idoeBMUAiSjcYJZfIIb7ptpKDiTtU5IQ
2022-06-11T17:36:23.359+0800    INFO    [init]  cmd/init.go:50  init app maudit success, client_id: 0y36drkRdoYKj7Jx3489AXB1, client_secret: t1rv1MJFyjO0wENEP1AodbgpuNiWUCEA
2022-06-11T17:36:23.361+0800    INFO    [init]  cmd/init.go:50  init app cmdb success, client_id: AB4S5Xl1zDgdYUh55q38XqB1, client_secret: niA8DYeRkbZFV2QIXutdAl4MZGdCIYEV
```

## 架构图



## 项目说明

```
├── protocol                       # 脚手架功能: rpc / http 功能加载
│   ├── grpc.go              
│   └── http.go    
├── client                         # 脚手架功能: grpc 客户端实现 
│   ├── client.go              
│   └── config.go    
├── cmd                            # 脚手架功能: 处理程序启停参数，加载系统配置文件
│   ├── root.go             
│   └── start.go                
├── conf                           # 脚手架功能: 配置文件加载
│   ├── config.go                  # 配置文件定义
│   ├── load.go                    # 不同的配置加载方式
│   └── log.go                     # 日志配置文件
├── dist                           # 脚手架功能: 构建产物
├── etc                            # 配置文件
│   ├── xxx.env
│   └── xxx.toml
├── apps                            # 具体业务场景的领域包
│   ├── all
│   │   |-- grpc.go                # 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序。  
│   │   |-- http.go                # 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载。                    
│   │   └── internal.go            #  注册所有内部服务模块, 无须对外暴露的服务, 用于内部依赖。 
│   ├── book                       # 具体业务场景领域服务 book
│   │   ├── http                   # http 
│   │   │    ├── book.go           # book 服务的http方法实现，请求参数处理、权限处理、数据响应等 
│   │   │    └── http.go           # 领域模块内的 http 路由处理，向系统层注册http服务
│   │   ├── impl                   # rpc
│   │   │    ├── book.go          # book 服务的rpc方法实现，请求参数处理、权限处理、数据响应等 
│   │   │    └── impl.go           # 领域模块内的 rpc 服务注册 ，向系统层注册rpc服务
│   │   ├──  pb                    # protobuf 定义
│   │   │     └── book.proto       # book proto 定义文件
│   │   ├── app.go                 # book app 只定义扩展
│   │   ├── book.pb.go             # protobuf 生成的文件
│   │   └── book_grpc.pb.go        # pb/book.proto 生成方法定义
├── version                        # 程序版本信息
│   └── version.go                    
├── README.md                    
├── main.go                        # Go程序唯一入口
├── Makefile                       # make 命令定义
└── go.mod                         # go mod 依赖定义
```


## 快速开发
make脚手架
```sh
➜  mcenter git:(master) ✗ make help
dep                            Get the dependencies
lint                           Lint Golang files
vet                            Run go vet
test                           Run unittests
test-coverage                  Run tests with coverage
build                          Local build
linux                          Linux build
run                            Run Server
clean                          Remove previous build
help                           Display this help screen
```

1. 使用安装依赖的Protobuf库(文件)
```sh
# 把依赖的probuf文件复制到/usr/local/include

# 创建protobuf文件目录
$ make -pv /usr/local/include/github.com/infraboard/mcube/pb

# 找到最新的mcube protobuf文件
$ ls `go env GOPATH`/pkg/mod/github.com/infraboard/

# 复制到/usr/local/include
$ cp -rf pb  /usr/local/include/github.com/infraboard/mcube/pb
```

2. 添加配置文件(默认读取位置: etc/mcenter.toml)
```sh
$ 编辑样例配置文件 etc/mcenter.toml.book
$ mv etc/mcenter.toml.book etc/mcenter.toml
```

3. 启动服务
```sh
# 编译protobuf文件, 生成代码
$ make gen
# 如果是MySQL, 执行SQL语句(docs/schema/tables.sql)
$ make init
# 下载项目的依赖
$ make dep
# 运行程序
$ make run
```

## 推送gitlab仓库

```
git remote add gitlab https://gitlab.com/infraboard/mcenter.git
git push -u gitlab --all 
```

## 相关文档

https://petstore.swagger.io/#/

+ [gitlab mcenter](https://gitlab.com/infraboard/mcenter)
