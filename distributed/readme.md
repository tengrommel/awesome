# Tcp拥塞

- 接收方网络资源繁忙，因未及时相应ACK导致发送方重传大量数据，这样将会导致网络更加拥堵
- 拥塞控制是动态调整win大小，不只是依赖缓冲区大小确定窗口大小

# Tcp拥塞控制过程
- 慢开始和拥塞避免
- 快速重传和快速恢复

# ReverseProxy 功能点
- 更改内容支持
- 错误信息回调
- 支持自定义负载均衡
- 连接池功能
- url重写功能
- 支持websocket服务
- 支持https代理

#　请求头参数

- X-Forwarded-For
    - 记录最后直连实际服务器之前，整个代理过程
    - 可能会被伪造

- X-Real-IP
    - 请求实际服务器的IP
    - 每过一层代理都会被覆盖掉，只需第一代理设置转发
    - 不会被伪造（客户端）
    
#　一个加权负载均衡
- Weight
> 初始化时对节点约定的权重
- currentWeight
> 节点临时权重，每轮都会变化
- effectiveWeight
> 节点有效权重，默认与Weight相同
- totalWeight
> 所有节点有效权重之和：sum(effectiveWeight)