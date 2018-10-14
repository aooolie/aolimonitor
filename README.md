
## 奥利分布式监控
### 自动发现

节点已经配置好master地址，一旦上线就向master发送通知消息。Master更新节点状态，如果是新节点则创建一条新记录。


* 从配置文件读出master地址和端口
* 向master发送上线通告
* 在未收到ack时重发上线通告
```java
Master信息存在/aolimonitor/config路径下
```
### 网络拓扑图
### 程序更新
### Log
### 推送数据到slave
### 数据上传


