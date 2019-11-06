package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-dev-tools/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

// NOTICE!!
// If you don't need output log with "gin log style", you can remove the codes.
// This function use to initial the "RedisClient", you need use your own config information.
func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1",
		Password: "", // default no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if nil != err {
		log.Fatal(err)
	}

}

// This slice use to save feature prefix of path of the request which include query params
// and you want load cache. The follow values only en example, you need set by yourself.
var CacheUrlPrefixSlice = []string{"/api/intro", "/api/qrcodes", "/api/news", "/api/new"}

// Cache middleware, only use for GET methods to load cache from redis database.
// If you want the request try load cache, you need add the request path prefix into the
// slice "CacheUrlPrefixSlice", and the prefix have be defined by your demand.
// When loading cache, the full request path and the query params will as a key.
func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		useCache := false
		for _, urlPrefix := range CacheUrlPrefixSlice {
			if strings.HasPrefix(c.Request.URL.String(), urlPrefix) && c.Request.Method == "GET" {
				useCache = true
				break
			}
		}
		if useCache {
			err := LoadCache(c)
			if nil == err {
				// Abort this request at now
				c.Abort()
			}
		}

		c.Next()

	}
}

// Try loading the Cache for this request, the full request path and the query params will as a key.
func LoadCache(c *gin.Context) error {
	data, err := RedisClient.Get(c.Request.URL.String()).Bytes()
	if nil != err {
		return err
	}
	body := new(gin.H)
	err = json.Unmarshal(data, body)
	if nil != err {
		return err
	}
	c.JSON(200, body)
	return nil
}

// Storage the data as cache, the full request path and the query params will as a key.
// You can use this function in your controller method
func StorageCache(c *gin.Context, body gin.H) {
	data, err := json.Marshal(body)
	if nil != err {
		utils.CustomLogger.Fprintln(c, 10000, fmt.Sprintf("Save Cache Error%s", err))
	}
	if ok, err := RedisClient.SetNX(c.Request.URL.String(), data, time.Hour*24).Result(); !ok {
		utils.CustomLogger.Fprintln(c, 10000, fmt.Sprintf("Save Cache Error%s", err))
	}
}

// Clean the cache data of one feature prefix.
// This is an simple demo, your may change a lot for your demand
func CleanCache(c *gin.Context, key string) {
	switch key {
	case "/api/intro", "/api/qrcodes":
		RedisClient.Del(key)
	case "/api/news", "/api/new":
		keys, err := RedisClient.Keys(key + "*").Result()
		if nil != err {
			utils.CustomLogger.Fprintln(c, 1000, fmt.Sprintf("Clean Cache Error: %s", err.Error()))
		}
		for _, delKey := range keys {
			RedisClient.Del(delKey)
		}
	}
}
