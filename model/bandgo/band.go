package bandgo

import (
    "time"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BandTeam 乐队表
type BandTeam struct {
    global.GVA_MODEL
    Name          string     `json:"name" gorm:"not null;comment:乐队名称"`
    Description   string     `json:"description" gorm:"type:text;comment:乐队描述"`
    FoundedDate   *time.Time `json:"founded_date" gorm:"comment:成立日期"`
    LogoURL       string     `json:"logo_url" gorm:"comment:logo URL"`
    CoverImageURL string     `json:"cover_image_url" gorm:"comment:封面图URL"`
    Location      string     `json:"location" gorm:"comment:地区"`
    Status        string     `json:"status" gorm:"default:active;comment:状态"`
    Verified      bool       `json:"verified" gorm:"default:false;comment:是否认证"`
    CreatorID     uint       `json:"creator_id" gorm:"comment:创建者ID"`
}

// BandMember 乐队成员表
type BandMember struct {
    BandID          uint       `json:"band_id" gorm:"primaryKey;comment:乐队ID"`
    UserID          uint       `json:"user_id" gorm:"primaryKey;comment:用户ID"`
    Role            string     `json:"role" gorm:"not null;comment:角色"`
    RoleDescription string     `json:"role_description" gorm:"type:text;comment:角色描述"`
    JoinDate        time.Time  `json:"join_date" gorm:"not null;comment:加入日期"`
    LeaveDate       *time.Time `json:"leave_date" gorm:"comment:离开日期"`
    IsAdmin         bool       `json:"is_admin" gorm:"default:false;comment:是否管理员"`
    Status          string     `json:"status" gorm:"default:active;comment:状态"`
}

// BandTeamGenre 乐队音乐类型关联表
type BandTeamGenre struct {
    BandID  uint `json:"band_id" gorm:"primaryKey;comment:乐队ID"`
    GenreID uint `json:"genre_id" gorm:"primaryKey;comment:类型ID"`
}

// BandJoinRequest 乐队加入申请表
type BandJoinRequest struct {
    global.GVA_MODEL
    BandID          uint   `json:"band_id" gorm:"index;comment:乐队ID"`
    UserID          uint   `json:"user_id" gorm:"index;comment:用户ID"`
    Message         string `json:"message" gorm:"type:text;comment:申请消息"`
    ProposedRole    string `json:"proposed_role" gorm:"comment:期望角色"`
    Status          string `json:"status" gorm:"default:pending;comment:状态"`
    RespondedBy     uint   `json:"responded_by" gorm:"comment:处理人ID"`
    ResponseMessage string `json:"response_message" gorm:"type:text;comment:回复消息"`
}

// BandInvitation 乐队邀请表
type BandInvitation struct {
    global.GVA_MODEL
    BandID          uint   `json:"band_id" gorm:"index;comment:乐队ID"`
    InvitedUserID   uint   `json:"invited_user_id" gorm:"index;comment:被邀请用户ID"`
    InviterID       uint   `json:"inviter_id" gorm:"comment:邀请人ID"`
    Message         string `json:"message" gorm:"type:text;comment:邀请消息"`
    ProposedRole    string `json:"proposed_role" gorm:"comment:提议角色"`
    Status          string `json:"status" gorm:"default:pending;comment:状态"`
    ResponseMessage string `json:"response_message" gorm:"type:text;comment:回复消息"`
} 

func (BandTeam) TableName() string {
	return "bandgo_teams"
}

func (BandMember) TableName() string {
	return "bandgo_members"
}

func (BandTeamGenre) TableName() string {
	return "bandgo_team_genres"
}

func (BandJoinRequest) TableName() string {
	return "bandgo_join_requests"
}   

func (BandInvitation) TableName() string {
	return "bandgo_invitations"
}


