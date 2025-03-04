package bandgo

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// BandUserVerification 用户认证表
type BandUserVerification struct {
    global.GVA_MODEL
    UserID            uint           `json:"user_id" gorm:"index;comment:用户ID"`
    DocumentType      string         `json:"document_type" gorm:"comment:证件类型"`
    DocumentURLs      datatypes.JSON `json:"document_urls" gorm:"type:json;comment:证件URL数组"`
    VerificationStatus string         `json:"verification_status" gorm:"default:pending;comment:认证状态"`
    VerifiedBy        *uint          `json:"verified_by" gorm:"comment:审核人ID"`
    Notes             string         `json:"notes" gorm:"type:text;comment:备注"`
}

// BandUserReview 用户评价表
type BandUserReview struct {
    global.GVA_MODEL
    ReviewerID  uint   `json:"reviewer_id" gorm:"comment:评价人ID"`
    ReviewedID  uint   `json:"reviewed_id" gorm:"comment:被评价人ID"`
    ProjectID   *uint  `json:"project_id" gorm:"comment:关联项目ID"`
    Rating      uint   `json:"rating" gorm:"comment:评分(1-5)"`
    Content     string `json:"content" gorm:"type:text;comment:评价内容"`
}

// BandTeamReview 乐队评价表
type BandTeamReview struct {
    global.GVA_MODEL
    ReviewerID uint   `json:"reviewer_id" gorm:"comment:评价人ID"`
    BandID     uint   `json:"band_id" gorm:"comment:乐队ID"`
    EventID    *uint  `json:"event_id" gorm:"comment:关联活动ID"`
    Rating     uint   `json:"rating" gorm:"comment:评分(1-5)"`
    Content    string `json:"content" gorm:"type:text;comment:评价内容"`
} 

func (BandUserVerification) TableName() string {
	return "bandgo_user_verifications"
}

func (BandUserReview) TableName() string {
	return "bandgo_user_reviews"
}

func (BandTeamReview) TableName() string {
	return "bandgo_team_reviews"
}
