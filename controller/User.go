package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	ginHelper "github.com/livegoplayer/go_gin_helper"
	myHelper "github.com/livegoplayer/go_helper"
	myLogger "github.com/livegoplayer/go_logger"
	"github.com/livegoplayer/go_user_rpc/user"
	userpb "github.com/livegoplayer/go_user_rpc/user/grpc"
	"github.com/spf13/viper"
)

func RegisterHandler(c *gin.Context) {
	captchaId := c.Request.FormValue("captchaId")
	answer := c.Request.FormValue("answer")

	//验证一下二维码是否正确
	CaptchaRes := myHelper.VerifyCaptchaWithId(captchaId, answer)
	if !CaptchaRes {
		ginHelper.ErrorResp(1, "验证码验证失败")
	}

	registerRequest := &userpb.RegisterRequest{}

	//todo
	err := c.Bind(registerRequest)
	ginHelper.CheckError(err, "数据验证失败")

	userClient := user.GetUserClient()
	res, err := userClient.Register(c, registerRequest)

	ginHelper.CheckError(err, "新建用户失败")
	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func LoginHandler(c *gin.Context) {
	loginRequest := &userpb.LoginRequest{}
	err := c.Bind(loginRequest)

	userClient := user.GetUserClient()
	res, err := userClient.Login(c, loginRequest)

	ginHelper.CheckError(err, "登录失败")

	data := res.GetData()

	//设置本域名下的cookie
	if data.Uid > 0 {
		c.SetCookie("us_user_cookie", data.Token, int(time.Hour.Seconds()*6), "/", "", false, false)
	}

	ginHelper.SuccessResp("ok", data)
}

func LogoutHandler(c *gin.Context) {
	//设置本域名下的cookie
	c.SetCookie("us_user_cookie", "", -1, "/", "", false, false)
	logoutRequest := &userpb.LogoutRequest{}
	err := c.Bind(logoutRequest)
	ginHelper.CheckError(err)

	userClient := user.GetUserClient()
	response, err := userClient.Logout(c, logoutRequest)
	ginHelper.CheckError(err)

	data := response.GetData()

	ginHelper.SuccessResp("ok", data)
}

func AddUserHandler(c *gin.Context) {
	addUserRequest := &userpb.AddUserRequest{}
	err := c.Bind(addUserRequest)

	userClient := user.GetUserClient()
	res, err := userClient.AddUser(c, addUserRequest)

	ginHelper.CheckError(err, "添加用户失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func GetUserListHandler(c *gin.Context) {
	//获取用户列表，需要权限验证
	userClient := user.GetUserClient()

	getUserList := &userpb.GetUserListRequest{}
	err := c.Bind(getUserList)
	ginHelper.CheckError(err, "获取用户列表")

	res, err := userClient.GetUserList(c, getUserList)

	ginHelper.CheckError(err, "获取用户列表")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func DelUserHandler(c *gin.Context) {
	delUserRequest := &userpb.DelUserRequest{}
	err := c.Bind(delUserRequest)

	userClient := user.GetUserClient()
	res, err := userClient.DelUser(c, delUserRequest)

	ginHelper.CheckError(err, "删除用户失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

type checkUserStatusRes struct {
	IsLogin     bool                 `json:"isLogin"`
	UserSession *userpb.UserSessions `json:"userSession"`
	Token       string               `json:"token"`
}

func CheckUserStatusHandler(c *gin.Context) {
	token, err := c.Cookie("us_user_cookie")
	ginHelper.CheckError(err, "获取cookie失败")

	//如果没有token，证明没有登录
	data := &checkUserStatusRes{}
	if token == "" {
		ginHelper.SuccessResp("ok", data)
	}

	checkUserStatusRequest := &userpb.CheckUserStatusRequest{}
	checkUserStatusRequest.Token = token

	userClient := user.GetUserClient()
	res, err := userClient.CheckUserStatus(c, checkUserStatusRequest)
	ginHelper.CheckError(err, "检查用户登录状态失败")

	data.UserSession = res.GetData().UserSession
	data.IsLogin = res.GetData().IsLogin
	data.Token = res.GetData().Token

	ginHelper.SuccessResp("ok", data)
}

func CheckUserAuthorityHandler(c *gin.Context) {
	checkUserAuthorityRequest := &userpb.CheckUserAuthorityRequest{}
	err := c.Bind(checkUserAuthorityRequest)

	userClient := user.GetUserClient()
	res, err := userClient.CheckUserAuthority(c, checkUserAuthorityRequest)

	ginHelper.CheckError(err, "检查用户登录状态失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func GetUserAuthorityListHandler(c *gin.Context) {
	getUserAuthorityListRequest := &userpb.GetUserAuthorityListRequest{}
	err := c.Bind(getUserAuthorityListRequest)

	userClient := user.GetUserClient()
	res, err := userClient.GetUserAuthorityList(c, getUserAuthorityListRequest)

	ginHelper.CheckError(err, "获取用户权限列表失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func GetAuthorityListHandler(c *gin.Context) {
	getAuthorityListRequest := &userpb.GetAuthorityListRequest{}
	err := c.Bind(getAuthorityListRequest)

	userClient := user.GetUserClient()
	res, err := userClient.GetAuthorityList(c, getAuthorityListRequest)

	ginHelper.CheckError(err, "获取权限列表失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func AddUserRoleHandler(c *gin.Context) {
	addUserRoleRequest := &userpb.AddUserRoleRequest{}
	err := c.Bind(addUserRoleRequest)

	userClient := user.GetUserClient()
	res, err := userClient.AddUserRole(c, addUserRoleRequest)

	ginHelper.CheckError(err, "获取用户角色列表失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func DelUserRoleHandler(c *gin.Context) {
	delUserRoleRequest := &userpb.DelUserRoleRequest{}
	err := c.Bind(delUserRoleRequest)

	userClient := user.GetUserClient()
	res, err := userClient.DelUserRole(c, delUserRoleRequest)

	ginHelper.CheckError(err, "删除用户角色失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func GetRoleListHandler(c *gin.Context) {
	getRoleListRequest := &userpb.GetRoleListRequest{}
	err := c.Bind(getRoleListRequest)

	userClient := user.GetUserClient()
	res, err := userClient.GetRoleList(c, getRoleListRequest)

	ginHelper.CheckError(err, "获取角色列表失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

func GetUserRoleListHandler(c *gin.Context) {
	getUserRoleListRequest := &userpb.GetUserRoleListRequest{}
	err := c.Bind(getUserRoleListRequest)

	userClient := user.GetUserClient()
	res, err := userClient.GetUserRoleList(c, getUserRoleListRequest)

	ginHelper.CheckError(err, "获取角色列表失败")

	data := res.GetData()

	ginHelper.SuccessResp("ok", data)
}

//主服务器请求检查是否登录
func CommonCheckTokenHandler(c *gin.Context) {
	//获取token，如果没有就设置
	token, err := c.Cookie("us_user_cookie")
	if token == "" {
		token = c.Request.FormValue("token")
		if token == "" {
			ginHelper.AuthResp("没有权限，请先登录", viper.GetString("user_app_host"))
		}
		//设置一下cookie
		c.SetCookie("us_user_cookie", token, int(time.Hour.Seconds()*6), "/", "", false, false)
	}

	//如果没有token，证明没有登录
	data := &checkUserStatusRes{}
	checkUserStatusRequest := &userpb.CheckUserStatusRequest{}
	checkUserStatusRequest.Token = token

	userClient := user.GetUserClient()
	res, err := userClient.CheckUserStatus(c, checkUserStatusRequest)
	if err != nil {
		if gin.IsDebugging() {
			ginHelper.CheckError(err)
		} else {
			myLogger.Error("获取用户鉴权失败" + err.Error())
			ginHelper.AuthResp("没有权限，请先登录", viper.GetString("user_app_host"))
		}
	}

	//如果没有登录的话
	if res.GetData().IsLogin == false {
		ginHelper.AuthResp("没有权限，请先登录", viper.GetString("user_app_host"))
	}
	data.UserSession = res.GetData().UserSession
	data.IsLogin = res.GetData().IsLogin
	data.Token = res.GetData().Token
}
