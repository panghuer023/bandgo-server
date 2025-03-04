package bandgo

import (
    "time"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BandProject 项目表
type BandProject struct {
    global.GVA_MODEL
    Title         string     `json:"title" gorm:"not null;comment:项目标题"`
    Description   string     `json:"description" gorm:"type:text;comment:项目描述"`
    ProjectType   string     `json:"project_type" gorm:"comment:项目类型"`
    StartDate     *time.Time `json:"start_date" gorm:"comment:开始日期"`
    EndDate       *time.Time `json:"end_date" gorm:"comment:结束日期"`
    Status        string     `json:"status" gorm:"default:active;comment:状态"`
    CoverImageURL string     `json:"cover_image_url" gorm:"comment:封面图URL"`
    CreatorID     uint       `json:"creator_id" gorm:"comment:创建者ID"`
}

// BandProjectParticipant 项目参与者表
type BandProjectParticipant struct {
    ProjectID      uint       `json:"project_id" gorm:"primaryKey;comment:项目ID"`
    ParticipantType string    `json:"participant_type" gorm:"primaryKey;comment:参与者类型(user/band)"`
    ParticipantID   uint      `json:"participant_id" gorm:"primaryKey;comment:参与者ID"`
    Role           string     `json:"role" gorm:"comment:角色"`
    JoinedAt       time.Time  `json:"joined_at" gorm:"autoCreateTime"`
    LeftAt         *time.Time `json:"left_at" gorm:"comment:离开时间"`
    Status         string     `json:"status" gorm:"default:active;comment:状态"`
}

// BandMusicWork 音乐作品表
type BandMusicWork struct {
    global.GVA_MODEL
    Title         string `json:"title" gorm:"not null;comment:作品标题"`
    Description   string `json:"description" gorm:"type:text;comment:作品描述"`
    ReleaseDate   *time.Time `json:"release_date" gorm:"comment:发布日期"`
    CoverImageURL string `json:"cover_image_url" gorm:"comment:封面图URL"`
    AudioURL      string `json:"audio_url" gorm:"comment:音频URL"`
    Duration      uint   `json:"duration" gorm:"comment:时长(秒)"`
    Lyrics        string `json:"lyrics" gorm:"type:text;comment:歌词"`
    CopyrightInfo string `json:"copyright_info" gorm:"type:text;comment:版权信息"`
    ProjectID     *uint  `json:"project_id" gorm:"comment:关联项目ID"`
    WorkType      string `json:"work_type" gorm:"comment:作品类型"`
} 

func (BandProject) TableName() string {
	return "bandgo_projects"
}

func (BandProjectParticipant) TableName() string {
	return "bandgo_project_participants"
}

func (BandMusicWork) TableName() string {
	return "bandgo_music_works"
}
