package controller

import (
	"github.com/gin-gonic/gin"
	. "github.com/livegoplayer/go_helper"
)

//获取二维码的方法
func CaptchaHandler(c *gin.Context) {
	//todo 请求次数限制，应该在服务器最外围做
	captchaId, captchaImg, err := MakeCaptcha(RandStringBytesMaskImprSrcUnsafe(6))
	if err != nil {
		ErrorResp(c, 1, err.Error())
	}

	m := make(map[string]string)
	m["captchaId"] = captchaId
	m["captchaImg"] = captchaImg
	SuccessResp(c, "ok", m)
}
