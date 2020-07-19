package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	. "github.com/livegoplayer/go_helper"
	myLogger "github.com/livegoplayer/go_logger"
	"github.com/spf13/viper"

	. "github.com/livegoplayer/go_user/controller"
)

func main() {
	// 初始化一个http服务对象
	//默认有写入控制台的日志
	// 把这两个处理器替换
	r := gin.New()
	r.Use(gin.Recovery())
	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)

	//加载.env文件
	LoadEnv()
	//gin的格式化参数
	//改造access log, 插入到数据库
	r.Use(myLogger.GetGinAccessFileLogger(viper.GetString("log.access_log_file_path"), viper.GetString("log.access_log_file_name")))

	//app_log
	myLogger.SetLogger(myLogger.GetMysqlLogger(viper.GetString("app_log_mysql_host"), viper.GetString("app_log_mysql_port"), viper.GetString("app_log_mysql_db_name")))

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

	//解决跨域问题的中间件
	r.Use(Cors())

	//更换binding
	binding.Validator = ValidatorV10

	r.POST("/api/user/register", RegisterHandler)
	r.POST("/api/user/login", LoginHandler)
	r.POST("/api/user/logout", LogoutHandler)
	r.POST("/api/user/addUser", AddUserHandler)
	r.POST("/api/user/delUser", DelUserHandler)
	r.GET("/api/user/checkUserStatus", CheckUserStatusHandler)
	r.GET("/api/user/checkUserAuthority", CheckUserAuthorityHandler)
	r.GET("/api/user/getUserAuthorityList", GetUserAuthorityListHandler)
	r.GET("/api/user/getAuthorityList", GetAuthorityListHandler)
	r.POST("/api/user/addUserRole", AddUserRoleHandler)
	r.POST("/api/user/delUserRole", DelUserRoleHandler)
	r.GET("/api/user/getRoleList", GetRoleListHandler)
	r.GET("/api/user/getUserRoleList", GetUserRoleListHandler)
	r.GET("/api/user/getUserList", GetUserListHandler)
	r.GET("/api/captcha/getCaptcha", CaptchaHandler)
	r.GET("/test", MyTestHandler)

	err := r.Run(":9091") // 监听并在 9091 上启动服务
	if err != nil {
		fmt.Printf("user server start error : " + err.Error())
		return
	}
}
