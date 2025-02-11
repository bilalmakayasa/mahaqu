package helper

import (
	"encoding/json"
	"mahaqu/src/utility"
	"time"

	"github.com/go-redis/redis"
)

var (
	client = &redisClient{}
)

type redisClient struct {
	c *redis.Client
}

//GetClient get the redis client
func RedisInitialize() *redisClient {
	c := redis.NewClient(&redis.Options{
		Addr: utility.GoDotEnvVariable("REDIS_HOST"),
	})

	if err := c.Ping().Err(); err != nil {
		panic("Unable to connect to redis " + err.Error())
	}
	client.c = c
	return client
}

//GetKey get key
func (client *redisClient) GetRedis(key string, src interface{}) error {
	val, err := client.c.Get(key).Result()
	if err == redis.Nil || err != nil {
		return err
	}
	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}
	return nil
}

//SetKey set key
func (client *redisClient) SetRedis(key string, value interface{}, expiration time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = client.c.Set(key, cacheEntry, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
