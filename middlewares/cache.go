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

// The tag string should be the part of full url, use to judge if this request need load cache
// and you want load cache. The follow values only en example, you need set by yourself.
var cacheUrlPartTags = []string{
    "/api/intro",
    "/api/qrcodes",
    "/api/news",
    "/api/new",
    "/api/company/info",
}

// Cache middleware, only the request, which method is 'GET' and url include the 'cacheUrlPartTag',
// would touch off.
func CacheMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        canLoad := false
        if c.Request.Method == "GET" {
            // check the url whether have contained the 'cacheUrlPartTag'
            for _, tag := range cacheUrlPartTags {
                if strings.Contains(c.Request.URL.String(), tag) {
                    canLoad = true
                    break
                }
            }
        }

        if canLoad {
            err := LoadCache(c)
            if nil == err {
                // Load cache success, abort this request at now
                c.Abort()

            }
        }

        // Don't need load or load cache fail, continue the handler func
        c.Next()
    }
}

// Try load the Cache for this request, the url as key
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

// Storage the data into cache DB, the key is request's url
// You can use this function in your controller method (http.HandlerFunc)
func StorageCache(c *gin.Context, body gin.H) {
    data, err := json.Marshal(body)
    if nil != err {
        utils.CustomLogger.Fprintln(c, 10000, fmt.Sprintf("Save Cache Error: %s", err))
    }
    if _, err := RedisClient.SetNX(c.Request.URL.String(), data, time.Hour*24).Result(); nil != err {
        utils.CustomLogger.Fprintln(c, 10000, fmt.Sprintf("Save Cache Error: %s", err))
    }
}

// WARNING !!!
// Clean the Cache data by the tag
// This is an simple demo, your may change a lot for your demand
func CleanCache(c *gin.Context, tag string) {
    pattern := fmt.Sprintf("*%s*", tag)
    keys, err := RedisClient.Keys(pattern).Result()
    if nil != err {
        utils.CustomLogger.Fprintln(c, 1000, fmt.Sprintf("Clean Cache Error: %s", err.Error()))
    }
    for _, delKey := range keys {
        RedisClient.Del(delKey)
    }
}
