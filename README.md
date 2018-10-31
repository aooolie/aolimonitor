## 奥利分布式监控

### 自动发现

​	agent: 监控器 server: 服务器，agent将数据上传到server，由server展示数据。

- server启动后将地址及端口写入redis。
- agent从配置文件读出redis地址和端口
- 向server发送上线通告
- server将节点信息存入redis并发送ack
```
server开一个进程监听9527端口，有请求过来就调用goroutine处理
```
- server检查agent 5min内有无上报数据，判断是否存活(Redis)

```java
redis信息存在/opt/go/Applications/OnlineDiscover/conf路径下
```

### 网络拓扑图

### 程序更新

### Log

### 推送数据到slave

### 数据上传

