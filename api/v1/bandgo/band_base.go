package bandgo

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	bandgoReq "github.com/flipped-aurora/gin-vue-admin/server/model/bandgo/request"
	bandgoRes "github.com/flipped-aurora/gin-vue-admin/server/model/bandgo/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// WeiXin
// @Tags      WeinXin
// @Summary   微信配置
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Configer true  "URL"
// @Success   200   {object}  response.Response{data=bandgoRes.WeiXinConfigResponse,msg=string}
// @Router    /base/wx/config [post]
func (b *BaseApi) WXConfig(c *gin.Context) {
	var req bandgoReq.Configer
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	url := req.Url
	if url == "" {
		response.FailWithMessage("请求参数缺失：url", c)
		return
	}

	// 生成签名所需参数
	noncestr := generateNonceStr()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// 获取ticket并生成签名
	ticket, err := WeiXinService.GetJsapiTicket()
	if err != nil {
		global.GVA_LOG.Error("获取api_ticket失败!", zap.Error(err))
		response.FailWithMessage("获取api_ticket失败： "+err.Error(), c)
		return
	}

	// 清理URL（移除hash部分）
	cleanURL := strings.Split(url, "#")[0]
	signature := generateSignature(ticket, noncestr, timestamp, cleanURL)

	configs := bandgoRes.WeiXinConfigResponse{
		AppID:     global.GVA_CONFIG.WeChat.AppID,
		Timestamp: timestamp,
		NonceStr:  noncestr,
		Signature: signature,
	}

	response.OkWithDetailed(configs, "获取微信配置成功", c)
}

// WeiXin
// @Tags      WeinXin
// @Summary   微信登录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      bandgoReq.WXLogin true  "code"
// @Success   200   {object}  response.Response{data=bandgoRes.WeiXinLoginResponse,msg=string}
// @Router    /base/wx/login [post]
func (b *BaseApi) WXLogin(c *gin.Context) {
	var req bandgoReq.WXLogin
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	code := req.Code
	if code == "" {
		response.FailWithMessage("请求参数缺失：code", c)
		return
	}

	accessToken, err := WeiXinService.GetAccessToken(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userInfo, err := WeiXinService.GetWeiXinUserInfo(accessToken)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(userInfo, "获取用户信息成功", c)

	// response.OkWithMessage("登录成功", c)
}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=bandgoRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /base/captcha [post]
func (b *BaseApi) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.GVA_CONFIG.Captcha.ImgHeight, global.GVA_CONFIG.Captcha.ImgWidth, global.GVA_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.GVA_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(bandgoRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.GVA_CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "验证码获取成功", c)
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}

// 生成随机字符串
func generateNonceStr() string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 15)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 生成JS-SDK签名
func generateSignature(ticket, noncestr, timestamp, url string) string {
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, noncestr, timestamp, url)
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
