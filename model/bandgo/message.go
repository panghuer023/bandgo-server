package bandgo

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	"gorm.io/datatypes"
)

// BandMessage 消息表
type BandMessage struct {
    global.GVA_MODEL
    Content        string         `json:"content" gorm:"type:text;comment:消息内容"`
    MediaURLs      datatypes.JSON `json:"media_urls" gorm:"type:json;comment:媒体URL数组"`
    SenderID       uint           `json:"sender_id" gorm:"index;comment:发送者ID"`
    ConversationID uint           `json:"conversation_id" gorm:"index;comment:对话ID"`
    ReadStatus     bool           `json:"read_status" gorm:"default:false;comment:阅读状态"`
}

// BandConversation 对话表
type BandConversation struct {
    global.GVA_MODEL
    Title     string `json:"title" gorm:"comment:对话标题"`
    IsGroup   bool   `json:"is_group" gorm:"default:false;comment:是否群聊"`
    CreatorID uint   `json:"creator_id" gorm:"comment:创建者ID"`
}

// BandConversationParticipant 对话参与者表
type BandConversationParticipant struct {
    ConversationID uint       `json:"conversation_id" gorm:"primaryKey;comment:对话ID"`
    UserID         uint       `json:"user_id" gorm:"primaryKey;comment:用户ID"`
    JoinedAt       time.Time  `json:"joined_at" gorm:"autoCreateTime;comment:加入时间"`
    LeftAt         *time.Time `json:"left_at" gorm:"comment:离开时间"`
    IsAdmin        bool       `json:"is_admin" gorm:"default:false;comment:是否管理员"`
    MutedUntil     *time.Time `json:"muted_until" gorm:"comment:禁言截止时间"`
} 

func (BandMessage) TableName() string {
	return "bandgo_messages"
}

func (BandConversation) TableName() string {
	return "bandgo_conversations"
}

func (BandConversationParticipant) TableName() string {
	return "bandgo_conversation_participants"
}