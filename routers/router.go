package routers

import (
	"github.com/gin-gonic/gin"

	. "github.com/livegoplayer/go_user/controller"
)

func InitAppRouter(r gin.IRoutes) {
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
}
