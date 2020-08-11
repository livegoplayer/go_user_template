package controller

import (
	"github.com/gin-gonic/gin"
	ginHelper "github.com/livegoplayer/go_gin_helper"
)

func MyTestHandler(c *gin.Context) {
	//myLogger.Info("test")
	ginHelper.SuccessResp("ok")
	ginHelper.SuccessResp("ok")
	panic("sdfkj")
}
