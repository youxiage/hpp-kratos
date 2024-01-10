# hpp-kratos

kratos 框架微服务实践

## Layout 目录说明
目前针对一些改动的部分加了注释，其他的见官方文档：https://go-kratos.dev/docs/intro/layout/
```
.
├── api // rpc接口
│   └── helloworld
├── cmd // 项目入口
│   └── server // 重命名，服务端
├── configs // 配置文件
│   ├── dev // 本地环境
│   ├── prod // 线上
│   └── test // 测试
├── internal 
│   ├── biz
│   ├── conf
│   ├── data
│   ├── server
│   └── service
├── third_party
│   ├── errors
│   ├── google
│   ├── openapi
│   ├── validate
│   └── README.md
├── Dockerfile
├── LICENSE
├── Makefile // 自定义编译命令
├── OLD-README.md
├── README.md
├── go.mod
└── go.sum
```


## 生成目录结构

```
tree -L 2 --dirsfirst
```

## 容器运行

```
待补充
```

## 本地运行

```
// 配置文件改动执行
make config
// 依赖注入改动执行
make wire
// 编译后本地启动
make run
```