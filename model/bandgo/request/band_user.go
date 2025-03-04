package request

// Register 用户注册请求
type Register struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Email    string `json:"email"`
    FullName string `json:"fullName"`
}

// Login 登录请求
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// UpdateProfile 更新用户资料请求
type UpdateProfile struct {
    Bio             string `json:"bio"`
    ProfileImageURL string `json:"profileImageURL"`
    CoverImageURL   string `json:"coverImageURL"`
    Location        string `json:"location"`
    PhoneNumber     string `json:"phoneNumber"`
} 