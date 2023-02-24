package cache

import (
	"gin-api/global"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// RedisSet 重新redis 方法
func RedisSet(ctx *gin.Context, key string, value string, expiration time.Duration) bool {
	redisClient := global.App.RDB
	_, err := redisClient.Set(ctx, key, value, expiration).Result()
	if err != nil {
		log.Printf("redis set error key = %v value = %v err = %v", key, value, err)
		return false
	}
	return true

}

// RedisGet 重新redis 方法 当key不存在 | 报错时 返回字符串
func RedisGet(ctx *gin.Context, key string) string {
	redisClient := global.App.RDB
	val, err := redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Println("key 不存在 = ", key)
		val = ""
	} else if err != nil {
		log.Printf("redis get error key = %v err = %v", key, err)
		val = ""
	}
	return val
}
