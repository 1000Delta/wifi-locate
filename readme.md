# wifi-locate 基于 WIFI 的室内定位

> 毕设项目，也作为 Go 后端开发练习

## 设计

采用微服务方式，前端请求发往 gateway, gateway 通过 RPC 向相关服务发起请求。

### 功能设计



### 服务架构

#### 请求网关

通过 Gin 构建网关服务，对向服务进行的请求进行包装，gin 的 handler 封装了对内网服务的 RPC 请求。

路径：`/gateway`

#### RPC 服务

// TODO
