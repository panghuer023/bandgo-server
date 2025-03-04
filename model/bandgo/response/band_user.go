package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
	"github.com/google/uuid"
)

// UserInfo 用户信息响应
type UserInfo struct {
    UUID              uuid.UUID   `json:"uuid"`
    Username        string `json:"username"`
    Email           string `json:"email"`
    FullName        string `json:"fullName"`
    Bio             string `json:"bio"`
    ProfileImageURL string `json:"profileImageURL"`
    CoverImageURL   string `json:"coverImageURL"`
    Location        string `json:"location"`
    AccountStatus   string `json:"accountStatus"`
}

// LoginResponse 登录响应
type LoginResponse struct {
    User  bandgo.BandUser `json:"user"`
    Token string   `json:"token"`
} 
// RegisterResponse 注册响应
type RegisterResponse struct {
    User  bandgo.BandUser `json:"user"`
    Token string   `json:"token"`
} 

