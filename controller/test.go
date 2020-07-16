package controller

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	helper "github.com/livegoplayer/go_helper"
)

var redisClient *redis.Client
var prefix = ""

func MyTestHandler(c *gin.Context) {
	session := &helper.UserSession{}
	key := "user_login_status_session_uid_" + strconv.Itoa(13)
	helper.SuccessResp(c, "ok", key)

	json := GetRedisKey(key).Val()
	res := helper.JsonDecode(helper.StringToBytes(json), session)

	helper.SuccessResp(c, "ok", res)
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func init() {
	// 根据redis配置初始化一个客户端
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "139.224.132.234:6379", // redis地址
		Password: "myredis",              // redis密码，没有则留空
		DB:       0,                      // 默认数据库，默认是0
	})

	prefix = "fs_redis_"
}

func GetRedisKey(key string) *redis.StringCmd {
	return redisClient.Get(prefix + key)
}

func SetRedisKey(key string, value []byte, expire time.Duration) *redis.StatusCmd {
	return redisClient.Set(prefix+key, value, expire)
}
