package statics

import (
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func init()  {
	config := GetConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Ip, config.Redis.Port),
		Password: config.Redis.Password, // no password set
		DB: config.Redis.DB,  // use default DB
		PoolSize:     15, // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
	})
}

func GetRedis() *redis.Client {
	return client
}