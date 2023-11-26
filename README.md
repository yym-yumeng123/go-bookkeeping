# go-bookkeeping

go 记账 api

```
api/ 后端接口描述模块
cmd 命令行
go.mod
go.sum
internal/ 内部模块
main.go 入口
test 测试
web 静态资源
```


```
go test
go mod
```

### 流程

1. 需求分析
2. 概要设计: 经典网站风格
   - 两层架构: B/S或C/S
   - 三层架构: 表现层+业务层+数据层
   - 单体架构: 组件互相依赖, 整体编译, 部署
   - 微服务架构: 松耦合 独立开发 独立部署
   - 无服务架构(Serveless): 不关心硬件 直接上云
   - MVC架构: 和三层相似
3. 概要设计: 接口风格
   - RESTful API 风格 (资源 + 动词)
   - RPC 风格 `POST /sayHello HTTP/1.1` 要求: 后端有 sayHello 函数
4. 工作排期
   - (标准工时 / 个人效率) + 测试 + 联调 + 缓冲 buffer
5. 与前端约定接口
   - 草拟接口名, 字段, 流程
6. 分配任务, 安排例会
7. 开发 前后端联调