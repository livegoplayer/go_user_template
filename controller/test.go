package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	ginHelper "github.com/livegoplayer/go_gin_helper"
)

var redisClient *redis.Client
var prefix = ""

func MyTestHandler(c *gin.Context) {
	//myLogger.Info("test")
	ginHelper.SuccessResp(c, "ok")
	ginHelper.SuccessResp(c, "ok")
	panic("sdfkj")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
