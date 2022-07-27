## lebron
Highly concurrent mall system built on go-zero

## Architecture 
A busy sum, when there is time to make up for it

商品服务(product) - 商品的添加、信息查询、库存管理等功能

购物车服务(cart) - 购物车的增删改查

订单服务(order) - 生成订单，订单管理

支付服务(pay) - 通过调用第三方支付实现支付功能

账号服务(user) - 用户信息、等级、封禁、地址管理

推荐服务(recommend) - 首页商品推荐

评论服务(reply) - 商品的评论功能、评论的回复功能

api - 对外的BFF服务，接受来自客户端的请求，暴露HTTP接口

rpc - 对内的微服务，仅接受来自内部其他微服务或者BFF的请求，暴露gRPC接口

rmq - 负责进行流式任务处理，上游一般依赖消息队列，比如kafka等

admin - 也是对内的服务，区别于rpc，更多的是面向运营侧的且数据权限较高，通过隔离可带来更好的代码级别的安全，直接提供HTTP接口
## Series of Courses
[go-zero 微服务实战系列（一、开篇）](https://mp.weixin.qq.com/s?__biz=Mzg2ODU1MTI0OA==&mid=2247485597&idx=1&sn=7e85894b7847cc50df51d66092792453&scene=21#wechat_redirect)

[go-zero 微服务实战系列（二、服务拆分）](https://mp.weixin.qq.com/s?__biz=Mzg2ODU1MTI0OA==&mid=2247485645&idx=1&sn=d329f56741dbe1f3e09713a6e4d1f7f0&scene=21#wechat_redirect)

[go-zero微服务实战系列（三、API定义和表结构设计）](https://mp.weixin.qq.com/s/ZWfzuJuJKeyJM3PMJ-SysQ)

[go-zero微服务实战系列（四、CRUD热身）](https://mp.weixin.qq.com/s/AIcJkMKTL0odqy1NzeJkxg)

[go-zero微服务实战系列（五、缓存代码怎么写）](https://mp.weixin.qq.com/s/QqrLOq7DcDVuIM_1YAaVTw)

[go-zero微服务实战系列（六、缓存一致性保证）](https://mp.weixin.qq.com/s/422ZHs81y7nN9Sgb_ESsgg)

[go-zero微服务实战系列（七、请求量这么高该如何优化）](https://mp.weixin.qq.com/s/pPPSPZJispmITY9Wsi7hUg)

[go-zero微服务实战系列（八、如何处理每秒上万次的下单请求）](https://mp.weixin.qq.com/s/OAbuzj876SrrcB5WO_2FuA)

[go-zero微服务实战系列（九、极致优化秒杀性能）](https://mp.weixin.qq.com/s/8VSS9WNSy4jkOSSIA4BmLw)


## Basic Environment
| Name    | Description              | Link                    |
| ------- | ------------------------ | ----------------------- |
| Go-Zero | Web & Rpc Go Frame       | https://go-zero.dev/cn/ |
| Mysql   | DB                       | https://www.mysql.com/  |
| Redis   | Cache                    | https://redis.io/       |
| Docker  | Code Runtime Environment | https://www.docker.com/ |
| MQ      |                          |                         |

## Code Components

| Name   | Description                               | Link                     |
| ------ | ----------------------------------------- | ------------------------ |
| sqlx   | db table crud                             |                          |
| copier | copy value from struct to struct and more | github.com/jinzhu/copier |




## Business Function
![img.png](doc/imgs/img.png)

## Server Port 

### API

| server name    | port  |
| -------------- | ------|
| api bff        | 8001  |


### RPC
| server name    | port  |
| -------------- | ------|
| user rpc       | 9001  |
| product rpc    | 9002  |
| order rpc      | 9003  |
