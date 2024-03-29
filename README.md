# gin-code-generator gin项目代码生成器

简称gcg

gcg 是一个基于 Go 语言开发的命令工具，可以快捷生成 model、router、service、api 文件甚至是整个项目。

我们的目标是1分钟新建一个项目并跑起来,1分钟生成一个接口并可以访问！

本生成器使用的[项目规范](#项目规范)

## 功能特性

- 生成Model
- 生成Service
- 生成Router
- 生成Api

## 数据库支持

- Oracle
- MySQL
- PostgreSQL

## 软件架构


## 项目结构说明

## 项目规范
- api api接口目录
  - v1 版本目录
  - 服务名Api.go
- config 配置
- core http服务器入口
- docs swagger接口文档地址
- global 全局设置
- initialize 初始化orm、redis等
- log 日志文件
- middleware 中间件
  - 中间件名Middleware.go
- model 数据库模型
  - 服务名Model.go
  - request 请求数据模型
    - 服务名ReqModel.go
  - response 返回数据模型
    - 服务名RespModel.go
- router 路由
  - 服务名Router.go
- service 服务层
  - 服务名Service.go
- utils 工具包
- config.yaml 配置
- main.go 程序入口

## 快速开始

### 依赖检查

**Minimum Requirements**

go 1.17


### 构建

1. 代码包下载

```
$ git clone https://github.com/jiangwu10057/gin-code-generator
```

2. 模板文件引入
```bash
$ go get -u github.com/jteeuwen/go-bindata/...
$ go-bindata -pkg=assets -o=assets/bindata.go assets/...
```
3. 编译

```bash
$ go build -o bin/gcg cmd/cli.go
```

## 项目配置文件
- [模板](./assets/config.yaml)

### 运行

- 新项目
```bash
gcg -module project -name project_name -path fullPath
```

- 生成接口

```bash
gcg -author jiangwu10057 -module quick -name auth -withtest -withcurd -tags swagger接口分组tag -apiv v1
```

1. 参数说明

| 名称| 说明|可选值|
|----------------|------------------|----------------|
|module|要生成的模块|model、router、service、api、quick（同时生成model、router、service、api）|
|name|项目名或表名||
|author|代码生成者|默认为计算机名|
|tags|api接口swagger分组tag||
|apiv|api版本号||
|withtest|是否同时生成单测文件|默认否|
|withcurd|是否同时生成curd接口|默认否|
|force|是否强制覆盖旧文件以及创建目录|默认否|


2. window运行注意事项

- 下载[instance client](https://oracle.github.io/odpi/doc/installation.html#id1)，并解压
- 设置环境变量PATH
- 控制台重启

## 使用指南

[Documentation](docs/guide/zh-CN)

## 如何贡献

欢迎贡献代码，贡献流程可以参考 [developer's documentation](docs/devel/zh-CN/development.md)。

## 社区

You are encouraged to communicate most things via [GitHub issues](https://github.com/jiangwu10057/gin-code-generator/issues/new/choose) or pull requests.

## 关于作者

- jiangwu10057 <jiangwu10057@qq.com>

为了方便交流，我建了微信群，可以加我 **微信：**，拉你入群，方便交流。

## 谁在用

<!-- 如果你有项目在使用，也欢迎联系作者，加入使用案例。 -->

## 许可证

gcg is licensed under the Apache License. See [LICENSE](LICENSE) for the full license text.