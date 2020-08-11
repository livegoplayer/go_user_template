package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ginHelper "github.com/livegoplayer/go_gin_helper"
	. "github.com/livegoplayer/go_helper"
	myLogger "github.com/livegoplayer/go_logger"
	userRpc "github.com/livegoplayer/go_user_rpc/user"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	. "github.com/livegoplayer/go_user/controller"
	. "github.com/livegoplayer/go_user/routers"
)

func main() {
	// 初始化一个http服务对象
	//默认有写入控制台的日志
	// 把这两个处理器替换
	r := gin.New()
	r.NoMethod(ginHelper.HandleNotFound)
	r.NoRoute(ginHelper.HandleNotFound)

	//加载.env文件
	LoadEnv()

	//设置gin的运行模式
	switch viper.GetString("ENV") {
	case PRODUCTION_ENV:
		gin.SetMode(gin.ReleaseMode)
	case DEVELOPMENT_ENV:
		gin.SetMode(gin.DebugMode)
		//额外放置一个可以在控制台打印access_log的中间件
		r.Use(gin.Logger())
	default:
		gin.SetMode(gin.DebugMode)
		r.Use(gin.Logger())
	}

	//gin的格式化参数
	//改造access log, 输出到文件
	r.Use(myLogger.GetGinAccessFileLogger(viper.GetString("log.access_log_file_path"), viper.GetString("app_name")+"_"+viper.GetString("log.access_log_file_name")))

	//如果是debug模式的话，使用logger另外打印一份输出到控制台的logger
	if gin.IsDebugging() {
		r.Use(gin.Logger())
		//额外输出错误异常栈
	}
	r.Use(ginHelper.ErrHandler())

	//app_log
	//如果是debug模式的话，直接打印到控制台
	var appLogger *logrus.Logger
	if gin.IsDebugging() {
		appLogger = myLogger.GetConsoleLogger()
	} else {
		appLogger = myLogger.GetMysqlLogger(viper.GetString("log.app_log_mysql_host"), viper.GetString("log.app_log_mysql_port"), viper.GetString("log.app_log_mysql_db_name"), viper.GetString("log.app_log_mysql_table_name"), viper.GetString("log.app_log_mysql_user"), viper.GetString("log.app_log_mysql_pass"))
	}
	myLogger.SetLogger(appLogger)

	//解决跨域问题的中间件
	r.Use(ginHelper.Cors(viper.GetStringSlice("client_list")))

	//更换校验器
	binding.Validator = ValidatorV10

	InitAppRouter(r)

	if !gin.IsDebugging() {
		ginHelper.AuthenticationMiddleware(CommonCheckTokenHandler)
	}

	userRpc.InitUserClient(viper.GetString("sso_host"))

	fmt.Printf("userClient connect success")
	err := r.Run(":" + viper.GetString("app_port"))
	if err != nil {
		fmt.Printf("server start error : " + err.Error())
		return
	}
}
