# web-servic demo---teste
A simple web application with go

### web-service版本1.0.1版本相关特性
* 支持toml配置文件配置相关type
* 支持go-vendor管理相关依赖包
* 基础web组件服务聚合
* dao层service层基础搭建

### User Guide
* 更改web-service.toml文件中的数据库配置（假定go相关环境已创建）
* 执行go build命令:go build
* 执行运行命令:go run main.go -conf /xxx/web-service/web-service.toml 配置文件的绝对路径
  -conf 后跟的文件路径为当前
* 打开浏览器，访问localhost:9090 (或者执行命令：curl -i localhost:9090)

### And More
* 如何通过govendor更新或下载相关的依赖？
   1. cd vendor
   2. 使用govendor list命令查看当前已下载的依赖包
   3. 选择其中的某个依赖进行更新，譬如：govendor update "github.com/BurntSushi/toml"（下载时改为fetch）
   4. 通过govendor --help可查看相关命令的使用

<hr>

##### Version 1.0.0

> 1.基础初始化版本

#### Version 1.0.1

>1.提供web接口，提供数据库接入，提供基础页面支持

>2.pull request 测试
