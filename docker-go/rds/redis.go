package rds

import "github.com/go-redis/redis"

var (
	RedisCli *redis.Client
)

func InitRedis(host string, password string, db int) error {

	RedisCli = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	return RedisCli.Ping().Err()

}

func GetRedisClient() *redis.Client {
	return RedisCli
}
