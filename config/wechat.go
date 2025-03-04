package config
// WeChat 微信配置
type WeChat struct {
	AppID     string `mapstructure:"appid" json:"appid" yaml:"appid"`
	AppSecret string `mapstructure:"appsecret" json:"appsecret" yaml:"appsecret"`

	WxBaseTokenURL string  `mapstructure:"wxbasetokenurl" json:"wxbasetokenurl" yaml:"wxbasetokenurl"`
	WxJsapiTicketURL string  `mapstructure:"wxjsapiticketurl" json:"wxjsapiticketurl" yaml:"wxjsapiticketurl"`

	WxAccessTokenURL string  `mapstructure:"wxaccesstokenurl" json:"wxaccesstokenurl" yaml:"wxaccesstokenurl"`
	WxUserInfoURL string  `mapstructure:"wxuserinfourl" json:"wxuserinfourl" yaml:"wxuserinfourl"`
}