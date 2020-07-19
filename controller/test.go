package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	myHelper "github.com/livegoplayer/go_helper"
)

var redisClient *redis.Client
var prefix = ""

func MyTestHandler(c *gin.Context) {
	//myLogger.Info("test")
	myHelper.SuccessResp(c, "ok")
	myHelper.SuccessResp(c, "ok")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
