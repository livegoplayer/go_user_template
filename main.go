package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	. "github.com/livegoplayer/go_helper"

	. "github.com/livegoplayer/go_user/controller"
)

func main() {
	// 初始化一个http服务对象
	//默认有写入控制台的日志
	r := gin.Default()
	// 把这两个处理器替换
	r.NoMethod(HandleNotFound)
	r.NoRoute(HandleNotFound)
	//增加一个recover在 中间件的执行链的最内层，不破坏原来Recover handler的结构，在最内层渲染并且返回api请求结果
	r.Use(ErrHandler())

	//解决跨域问题的中间件
	r.Use(Cors())

	//更换binding
	binding.Validator = ValidatorV10

	r.POST("/api/user/register", RegisterHandler)
	r.POST("/api/user/login", LoginHandler)
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
	r.GET("/api/captcha/getCaptcha", CaptchaHandler)

	err := r.Run(":9091") // 监听并在 9091 上启动服务
	if err != nil {
		fmt.Printf("user server start error : " + err.Error())
		return
	}
}
