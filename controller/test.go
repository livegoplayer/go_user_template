package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	myHelper "github.com/livegoplayer/go_helper"
	myLogger "github.com/livegoplayer/go_logger"
)

var redisClient *redis.Client
var prefix = ""

func MyTestHandler(c *gin.Context) {
	myLogger.Info("test")
	myHelper.SuccessResp(c, "ok")
}

func GetRedisClient() *redis.Client {
	return redisClient
}
