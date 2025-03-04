package bandgo

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
    "time"
)

type BandEventService struct{}

// CreateEvent 创建活动
func (s *BandEventService) CreateEvent(event *bandgo.BandEvent) error {
    return global.GVA_DB.Create(event).Error
}

// UpdateEvent 更新活动信息
func (s *BandEventService) UpdateEvent(eventID uint, event *bandgo.BandEvent) error {
    return global.GVA_DB.Model(&bandgo.BandEvent{}).Where("id = ?", eventID).
        Omit("id", "creator_id").Updates(event).Error
}

// GetEventInfo 获取活动详情
func (s *BandEventService) GetEventInfo(eventID uint) (*bandgo.BandEvent, error) {
    var event bandgo.BandEvent
    err := global.GVA_DB.First(&event, eventID).Error
    return &event, err
}

// ListEvents 获取活动列表
func (s *BandEventService) ListEvents(query struct {
    CreatorID *uint
    EventType *string
    StartTime *time.Time
    EndTime   *time.Time
    IsPublic  *bool
}) ([]bandgo.BandEvent, error) {
    db := global.GVA_DB.Model(&bandgo.BandEvent{})

    if query.CreatorID != nil {
        db = db.Where("creator_id = ?", *query.CreatorID)
    }
    if query.EventType != nil {
        db = db.Where("event_type = ?", *query.EventType)
    }
    if query.StartTime != nil {
        db = db.Where("start_datetime >= ?", *query.StartTime)
    }
    if query.EndTime != nil {
        db = db.Where("start_datetime <= ?", *query.EndTime)
    }
    if query.IsPublic != nil {
        db = db.Where("is_public = ?", *query.IsPublic)
    }

    var events []bandgo.BandEvent
    err := db.Find(&events).Error
    return events, err
} 