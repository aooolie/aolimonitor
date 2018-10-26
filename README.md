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
- 在未收到ack时重发上线通告
- server定期(5min)轮询所有节点
- agent离线时发送离线通告

```java
redis信息存在/opt/go/Applications/OnlineDiscover/conf路径下
```

### 网络拓扑图

### 程序更新

### Log

### 推送数据到slave

### 数据上传

