package request

// Register 用户注册请求
type Configer struct {
    Url string `json:"url"`
}

// WeiXinLogin 登录请求
type WXLogin struct {
	Code  string `json:"code"`  // 授权码

}