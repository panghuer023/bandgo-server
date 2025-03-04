package bandgo

import (
	// "time"

	// "github.com/flipped-aurora/gin-vue-admin/server/api/v1/bandgo"
	// sys "github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	// "github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bandgo/request"
	bandgoRes "github.com/flipped-aurora/gin-vue-admin/server/model/bandgo/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	// "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	
)

type BandUserApi struct{}

// Register
// @Tags      User
// @Summary   用户注册
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Register  true  "用户名, 密码, 邮箱"
// @Success   200   {object}  response.Response{data=bandgo.BandUser,msg=string}
// @Router    /user/register [post]
func (b *BandUserApi) Register(c *gin.Context) {
	var req request.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user, err := bandUserService.Register(req)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage("注册失败: "+err.Error(), c)
		return
	}
	response.OkWithData(user, c)
}

// Login
// @Tags      User
// @Summary   用户登录
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Login  true  "用户名, 密码"
// @Success   200   {object}  response.Response{data=bandgoRes.LoginResponse,msg=string}
// @Router    /user/login [post]
func (b *BandUserApi) Login(c *gin.Context) {
	var req request.Login
	err := c.ShouldBindJSON(&req)
	// key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// if err := c.ShouldBindJSON(&req); err != nil {
	//     response.FailWithMessage(err.Error(), c)
	//     return
	// }

	// 判断验证码是否开启
	// openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	// openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	// v, ok := global.BlackCache.Get(key)
	// if !ok {
	// 	global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	// }

	// var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	// if !oc || (req.CaptchaId != "" && req.Captcha != "" && store.Verify(req.CaptchaId, req.Captcha, true)) {
	// 	u := &bandgo.BandUser{Username: req.Username, Password: req.Password}
	// 	user, err := bandUserService.Login(u)
	// 	if err != nil {
	// 		global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
	// 		// 验证码次数+1
	// 		global.BlackCache.Increment(key, 1)
	// 		response.FailWithMessage("用户名不存在或者密码错误", c)
	// 		return
	// 	}
	// 	// if user.Enable != 1 {
	// 	// 	global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
	// 	// 	// 验证码次数+1
	// 	// 	global.BlackCache.Increment(key, 1)
	// 	// 	response.FailWithMessage("用户被禁止登录", c)
	// 	// 	return
	// 	// }
	// 	b.TokenNext(c, *user)
	// 	return
	// }

	// 验证码次数+1
	// global.BlackCache.Increment(key, 1)
	// response.FailWithMessage("验证码错误", c)

	user, err := bandUserService.Login(req)
	if err != nil {
		global.GVA_LOG.Error("登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败: "+err.Error(), c)
		return
	}

	// 生成token等登录逻辑
	// token, err := utils.GenerateToken(user.ID)
	// if err != nil {
	//     global.GVA_LOG.Error("生成token失败!", zap.Error(err))
	//     response.FailWithMessage("登录失败: "+err.Error(), c)
	//     return
	// }

	response.OkWithData(bandgoRes.LoginResponse{
		User:  *user,
		Token: "sdfsdfsdf",
	}, c)
}

// // TokenNext 登录以后签发jwt
// func (b *BandUserApi) TokenNext(c *gin.Context, user bandgo.BandUser) {
// 	token, claims, err := utils.LoginToken(&user)
// 	if err != nil {
// 		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
// 		response.FailWithMessage("获取token失败", c)
// 		return
// 	}
// 	if !global.GVA_CONFIG.System.UseMultipoint {
// 		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
// 		response.OkWithDetailed(systemRes.LoginResponse{
// 			User:      user,
// 			Token:     token,
// 			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
// 		}, "登录成功", c)
// 		return
// 	}

// 	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
// 		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
// 			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
// 			response.FailWithMessage("设置登录状态失败", c)
// 			return
// 		}
// 		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
// 		response.OkWithDetailed(systemRes.LoginResponse{
// 			User:      user,
// 			Token:     token,
// 			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
// 		}, "登录成功", c)
// 	} else if err != nil {
// 		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
// 		response.FailWithMessage("设置登录状态失败", c)
// 	} else {
// 		var blackJWT system.JwtBlacklist
// 		blackJWT.Jwt = jwtStr
// 		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
// 			response.FailWithMessage("jwt作废失败", c)
// 			return
// 		}
// 		if err := jwtService.SetRedisJWT(token, user.GetUsername()); err != nil {
// 			response.FailWithMessage("设置登录状态失败", c)
// 			return
// 		}
// 		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
// 		response.OkWithDetailed(systemRes.LoginResponse{
// 			User:      user,
// 			Token:     token,
// 			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
// 		}, "登录成功", c)
// 	}
// }

// GetProfile
// @Tags      BandUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=bandgo.BandUser,msg=string}
// @Router    /user/profile [get]
func (b *BandUserApi) GetProfile(c *gin.Context) {
	userID := utils.GetUserID(c)
	user, err := bandUserService.GetUserProfile(userID)
	if err != nil {
		global.GVA_LOG.Error("获取资料失败!", zap.Error(err))
		response.FailWithMessage("获取资料失败", c)
		return
	}
	response.OkWithData(user, c)
}

// UpdateProfile
// @Tags      User
// @Summary   更新用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.UpdateProfile  true  "用户信息"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/profile [put]
func (b *BandUserApi) UpdateProfile(c *gin.Context) {
	var req request.UpdateProfile
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	userID := utils.GetUserID(c)
	if err := bandUserService.UpdateProfile(userID, req); err != nil {
		global.GVA_LOG.Error("更新资料失败!", zap.Error(err))
		response.FailWithMessage("更新资料失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetUserSkills
// @Tags      User
// @Summary   获取用户技能列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]bandgo.BandUserSkill,msg=string}
// @Router    /user/skills [get]
func (b *BandUserApi) GetUserSkills(c *gin.Context) {
	userID := utils.GetUserID(c)
	skills, err := bandUserService.GetUserSkills(userID)
	if err != nil {
		global.GVA_LOG.Error("获取技能失败!", zap.Error(err))
		response.FailWithMessage("获取技能失败", c)
		return
	}
	response.OkWithData(skills, c)
}

// AddUserSkill
// @Tags      User
// @Summary   添加用户技能
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      bandgo.BandUserSkill  true  "技能信息"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /bandUser/skills [post]
// func (b *BandUserApi) AddUserSkill(c *gin.Context) {
//     var skill bandgo.BandUserSkill
//     if err := c.ShouldBindJSON(&skill); err != nil {
//         response.FailWithMessage(err.Error(), c)
//         return
//     }

//     userID := utils.GetUserID(c)
//     skill.UserID = userID
//     if err := bandUserService.AddUserSkill(userID, &skill); err != nil {
//         global.GVA_LOG.Error("添加技能失败!", zap.Error(err))
//         response.FailWithMessage("添加技能失败", c)
//         return
//     }
//     response.OkWithMessage("添加成功", c)
// }

// // UpdateUserSkill
// // @Tags      User
// // @Summary   更新用户技能
// // @Security  ApiKeyAuth
// // @accept    application/json
// // @Produce   application/json
// // @Param     id    path      uint                  true  "技能ID"
// // @Param     data  body      bandgo.BandUserSkill  true  "技能信息"
// // @Success   200   {object}  response.Response{msg=string}
// // @Router    /user/skills/{id} [put]
// func (b *BandUserApi) UpdateUserSkill(c *gin.Context) {
//     var skill bandgo.BandUserSkill
//     if err := c.ShouldBindJSON(&skill); err != nil {
//         response.FailWithMessage(err.Error(), c)
//         return
//     }

//     userID := utils.GetUserID(c)
//     skillID := c.Param("id")
//     if err := bandUserService.UpdateUserSkill(userID, skillID, &skill); err != nil {
//         global.GVA_LOG.Error("更新技能失败!", zap.Error(err))
//         response.FailWithMessage("更新技能失败", c)
//         return
//     }
//     response.OkWithMessage("更新成功", c)
// }


