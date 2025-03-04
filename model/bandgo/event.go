package bandgo

import (
	"time"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BandEvent 活动表
type BandEvent struct {
	global.GVA_MODEL
	Title            string     `json:"title" gorm:"not null;comment:活动标题"`
	Description      string     `json:"description" gorm:"type:text;comment:活动描述"`
	VenueName        string     `json:"venue_name" gorm:"comment:场地名称"`
	VenueAddress     string     `json:"venue_address" gorm:"type:text;comment:场地地址"`
	VenueCoordinates string     `json:"venue_coordinates" gorm:"comment:场地坐标"`
	StartDateTime    time.Time  `json:"start_datetime" gorm:"not null;comment:开始时间"`
	EndDateTime      *time.Time `json:"end_datetime" gorm:"comment:结束时间"`
	EventType        string     `json:"event_type" gorm:"comment:活动类型"`
	IsPublic         bool       `json:"is_public" gorm:"default:true;comment:是否公开"`
	CoverImageURL    string     `json:"cover_image_url" gorm:"comment:封面图URL"`
	CreatorID        uint       `json:"creator_id" gorm:"comment:创建者ID"`
}

// BandRehearsal 排练表
type BandRehearsal struct {
	global.GVA_MODEL
	BandID              uint      `json:"band_id" gorm:"index;comment:乐队ID"`
	Title               string    `json:"title" gorm:"not null;comment:排练标题"`
	Description         string    `json:"description" gorm:"type:text;comment:排练描述"`
	StartTime           time.Time `json:"start_time" gorm:"not null;comment:开始时间"`
	EndTime            time.Time `json:"end_time" gorm:"not null;comment:结束时间"`
	Location           string    `json:"location" gorm:"comment:地点"`
	LocationCoordinates string    `json:"location_coordinates" gorm:"comment:地点坐标"`
	Notes              string    `json:"notes" gorm:"type:text;comment:备注"`
	CreatedBy          uint      `json:"created_by" gorm:"comment:创建者ID"`
} 
func (BandEvent) TableName() string {
	return "bandgo_events"
}

func (BandRehearsal) TableName() string {
	return "bandgo_rehearsals"
}
