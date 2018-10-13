# aolimonitor
分布式监控

1.自动发现
	节点已经配置好master地址，一旦上线就向master发送通知消息。Master更新节点状态，如果是新节点则创建一条新记录。
1)	从配置文件读出master地址和端口
2)	向master发送上线通告
3)	在未收到ack时重发上线通告

Master信息存在/aolimonitor/config路径下

2.网络拓扑图
3.程序更新
4.Log
5.推送数据到slave
6.数据上传
