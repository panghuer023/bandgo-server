package bandgo

// const (
// 	// OAuth2 相关接口
// 	WxAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"
// 	WxUserInfoURL   = "https://api.weixin.qq.com/sns/userinfo"
// 	// 基础接口
// 	WxBaseTokenURL   = "https://api.weixin.qq.com/cgi-bin/token"
// 	WxJsapiTicketURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
// )

// WxCache 微信接口调用凭证缓存
type WxCache struct {
	AccessToken         string
	AccessTokenExpires int64
	JsapiTicket        string
	JsapiTicketExpires int64
} 