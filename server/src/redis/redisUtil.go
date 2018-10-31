package redis

import (
	"util"
	"github.com/garyburd/redigo/redis"
	"fmt"
)

/* 将 Server url 存到 Redis */
func SetServerUrl2Redis() {

	ips := util.GetHostIP()
	if len(ips) == 0 {
		fmt.Print("[error]: ips len is 0")
		return
	}

	/* 使用配置文件中的 Redis 服务器地址 */
	redisUrl := util.REDIS_HOST + ":" + util.REDIS_PORT

	fmt.Println("redisUrl: " + redisUrl + "\n")
	c, err := redis.Dial("tcp", redisUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(ips); i++ {
		_, err := redis.String(c.Do("SET", util.SERVER_IP, ips[i]))
		if err != nil {
			fmt.Println(err)
			return
		}

		/* TODO 检测server上空闲端口 */
		_, err = redis.String(c.Do("SET", util.SERVER_PORT, "9090"))
		if err != nil {
			fmt.Println(err)
			return
		}

	}
	fmt.Println("[info] write server ip to redis:", ips)
	defer c.Close()
}

/* 将 agent ip 存进 Redis */
func SetAgentIP2Redis(agentIP string, timestamp []byte) {

	/* 使用配置文件中的 Redis 服务器地址 */
	redisUrl := util.REDIS_HOST + ":" + util.REDIS_PORT

	fmt.Println("redisUrl: " + redisUrl + "\n")
	c, err := redis.Dial("tcp", redisUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("[info] write server ip to redis:", agentIP, len(timestamp))
	_, err = redis.String(c.Do("SET", agentIP, timestamp))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()
}