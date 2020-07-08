package kerbalwzygo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client
var ctx = context.Background()

// 初始化RedisClient, 在程序启动阶段的代码中去执行
func InitRedisClient(addr, password string, db int) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // default no password set
		DB:       db,       // use default DB
	})

	_, err := redisClient.Ping(ctx).Result()
	if nil != err {
		log.Fatal(err)
	}
}

// 缓存读取中间件
func CacheInRedisMiddleware(cacheUrlTags []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		canLoad := false
		if c.Request.Method == "GET" {
			// check the url whether have contained the 'cacheUrlPartTag'
			for _, tag := range cacheUrlTags {
				if strings.Contains(c.Request.URL.String(), tag) {
					canLoad = true
					break
				}
			}
		}

		if canLoad {
			data, err := LoadCacheInRedis(c)
			if nil == err {
				// Load cache success, abort this request at now
				body := gin.H{}
				err = json.Unmarshal(data, body)
				if nil == err {
					c.JSON(200, body)
					c.Abort()
				}
			}
		}

		// Don't need load or load cache fail, continue the handler func
		c.Next()
	}
}

// 加载缓存
func LoadCacheInRedis(c *gin.Context) ([]byte, error) {
	data, err := redisClient.Get(ctx, c.Request.URL.String()).Bytes()
	return data, err
}

// 存储缓存数据
func StorageCacheInRedis(c *gin.Context, body gin.H, expire time.Duration) ([]byte, error) {
	data, err := json.Marshal(body)
	if nil != err {
		return nil, err
	}
	if _, err := redisClient.SetNX(ctx, c.Request.URL.String(), data, expire).Result(); nil != err {
		return nil, err
	}
	return data, nil
}

// 删除缓存
func CleanCacheInRedis(tag string) error {
	pattern := fmt.Sprintf("*%s*", tag)
	keys, err := redisClient.Keys(ctx, pattern).Result()
	if nil != err {
		return err
	}
	for _, delKey := range keys {
		redisClient.Del(ctx, delKey)
	}
	return nil
}
