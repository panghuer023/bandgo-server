package bandgo

import (
    "time"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
	"github.com/google/uuid"
)

// BandUser 用户基础信息表
type BandUser struct {
    global.GVA_MODEL
    UUID      uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
    Username              string         `json:"username" gorm:"uniqueIndex;not null;comment:用户名"`
    Password             string         `json:"-" gorm:"not null;comment:用户登录密码"`
    FullName             string         `json:"full_name" gorm:"not null;comment:全名"`
    NickName             string         `json:"nickname" gorm:"comment:昵称"`
    Bio                  string         `json:"bio" gorm:"type:text;comment:个人简介"`
    ProfileImageURL      string         `json:"profile_image_url" gorm:"comment:头像URL"`
    CoverImageURL        string         `json:"cover_image_url" gorm:"comment:封面图URL"`
    Location             string         `json:"location" gorm:"comment:地区"`
    DateOfBirth         *time.Time     `json:"date_of_birth" gorm:"comment:出生日期"`
    AccountStatus       string         `json:"account_status" gorm:"default:active;comment:账号状态"`
    VerificationStatus  bool           `json:"verification_status" gorm:"default:false;comment:验证状态"`
    LastLoginAt         *time.Time     `json:"last_login_at" gorm:"comment:最后登录时间"`
    PreferredLanguage   string         `json:"preferred_language" gorm:"default:zh;comment:首选语言"`
    // swagger:strfmt json
    NotificationPrefs   datatypes.JSON `json:"notification_preferences" gorm:"type:json;comment:通知偏好"`
    IsProUser           bool           `json:"is_pro_user" gorm:"default:false;comment:是否专业用户"`
    // swagger:strfmt json
    SocialLinks         datatypes.JSON `json:"social_links" gorm:"type:json;comment:社交链接"`
    // 添加微信相关字段
	OpenID    string    `json:"openid" gorm:"unique;comment:微信用户唯一标识"`
	UnionID   string    `json:"unionid" gorm:"comment:微信用户统一标识"`
	Sex       int       `json:"sex" gorm:"default:0;comment:性别 0-未知 1-男 2-女"`
	Province  string    `json:"province" gorm:"comment:省份"`
	City      string    `json:"city" gorm:"comment:城市"`
	Country   string    `json:"country" gorm:"comment:国家"`
	Avatar    string    `json:"avatar" gorm:"comment:头像URL"`
	Phone     string    `json:"phone"  gorm:"comment:用户手机号"`
	Email     string    `json:"email"  gorm:"comment:用户邮箱"`
	Enable    int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
}


// BandUserSkill 用户音乐技能表
type BandUserSkill struct {
    global.GVA_MODEL
    UserID           uint   `json:"user_id" gorm:"index;comment:用户ID"`
    InstrumentID     uint   `json:"instrument_id" gorm:"index;comment:乐器ID"`
    SkillLevel       string `json:"skill_level" gorm:"comment:技能等级"`
    YearsExperience  uint   `json:"years_experience" gorm:"comment:演奏年限"`
    IsPrimary        bool   `json:"is_primary" gorm:"default:false;comment:是否主要乐器"`
    SkillDescription string `json:"skill_description" gorm:"type:text;comment:技能描述"`
}

// BandInstrument 乐器表
type BandInstrument struct {
    global.GVA_MODEL
    Name     string `json:"name" gorm:"not null;comment:乐器名称"`
    Category string `json:"category" gorm:"not null;comment:乐器类别"`
    IconURL  string `json:"icon_url" gorm:"comment:图标URL"`
}

// BandGenre 音乐类型表
type BandGenre struct {
    global.GVA_MODEL
    Name          string `json:"name" gorm:"not null;comment:类型名称"`
    ParentGenreID *uint  `json:"parent_genre_id" gorm:"comment:父类型ID"`
    IconURL       string `json:"icon_url" gorm:"comment:图标URL"`
}

// BandUserGenre 用户音乐类型偏好表
type BandUserGenre struct {
    UserID          uint `json:"user_id" gorm:"primaryKey;comment:用户ID"`
    GenreID         uint `json:"genre_id" gorm:"primaryKey;comment:类型ID"`
    PreferenceLevel uint `json:"preference_level" gorm:"comment:偏好程度(1-10)"`
} 

func (BandUser) TableName() string {
	return "bandgo_users"
}

func (BandUserSkill) TableName() string {
	return "bandgo_user_skills"
}

func (BandInstrument) TableName() string {
	return "bandgo_instruments"
}

func (BandGenre) TableName() string {
	return "bandgo_genres"
}

func (BandUserGenre) TableName() string {
	return "bandgo_user_genres"
}
