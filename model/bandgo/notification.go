package bandgo

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// BandNotification 通知表
type BandNotification struct {
    global.GVA_MODEL
    UserID            uint   `json:"user_id" gorm:"index;comment:用户ID"`
    Content           string `json:"content" gorm:"type:text;not null;comment:通知内容"`
    NotificationType  string `json:"notification_type" gorm:"not null;comment:通知类型"`
    RelatedEntityType string `json:"related_entity_type" gorm:"comment:关联实体类型"`
    RelatedEntityID   uint   `json:"related_entity_id" gorm:"comment:关联实体ID"`
    IsRead            bool   `json:"is_read" gorm:"default:false;comment:是否已读"`
}

// BandUserDevice 用户设备表
type BandUserDevice struct {
    global.GVA_MODEL
    UserID      uint      `json:"user_id" gorm:"index;comment:用户ID"`
    DeviceToken string    `json:"device_token" gorm:"not null;comment:设备令牌"`
    DeviceType  string    `json:"device_type" gorm:"not null;comment:设备类型"`
    IsActive    bool      `json:"is_active" gorm:"default:true;comment:是否活跃"`
    LastUsedAt  time.Time `json:"last_used_at" gorm:"autoUpdateTime;comment:最后使用时间"`
} 

func (BandNotification) TableName() string {
	return "bandgo_notifications"
}

func (BandUserDevice) TableName() string {
	return "bandgo_user_devices"
}
