package redis

import (
	"util"
	"github.com/garyburd/redigo/redis"
	"fmt"
)
/* 将 Server url 存到 Redis */
func SetServerUrl2Redis() {

	ips := util.GetHostIP()

	/* 使用配置文件中的 Redis 服务器地址 */
	redisUrl := util.REDIS_HOST + ":" + util.REDIS_PORT

	c, err := redis.Dial("tcp", redisUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	for ip, _ := range ips {
		_, err := redis.String(c.Do("SET", util.SERVER_IP, ip))
		if err != nil {
			fmt.Println(err)
			return
		}

		/* TODO 检测server上空闲端口 */
		_, err = redis.String(c.Do("SET", util.SERVER_PORT, "9527"))
		if err != nil {
			fmt.Println(err)
			return
		}

	}
	fmt.Println("[info] write server ip to redis:", ips)
	defer c.Close()
}