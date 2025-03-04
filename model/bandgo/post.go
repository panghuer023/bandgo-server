package bandgo

import (
    "gorm.io/datatypes"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// BandPost 帖子表
type BandPost struct {
    global.GVA_MODEL
    Content             string         `json:"content" gorm:"type:text;comment:内容"`
    MediaURLs          datatypes.JSON `json:"media_urls" gorm:"type:json;comment:媒体URL数组"`
    CreatorType        string         `json:"creator_type" gorm:"not null;comment:创建者类型(user/band)"`
    CreatorID          uint           `json:"creator_id" gorm:"not null;comment:创建者ID"`
    IsPublic           bool           `json:"is_public" gorm:"default:true;comment:是否公开"`
    IsPinned           bool           `json:"is_pinned" gorm:"default:false;comment:是否置顶"`
    Location           string         `json:"location" gorm:"comment:位置"`
    LocationCoordinates string         `json:"location_coordinates" gorm:"comment:位置坐标"`
    EventID            *uint          `json:"event_id" gorm:"comment:关联活动ID"`
    EngagementCount    datatypes.JSON `json:"engagement_count" gorm:"type:json;comment:互动数据"`
}

// BandComment 评论表
type BandComment struct {
    global.GVA_MODEL
    Content    string         `json:"content" gorm:"type:text;not null;comment:评论内容"`
    MediaURLs  datatypes.JSON `json:"media_urls" gorm:"type:json;comment:媒体URL数组"`
    UserID     uint           `json:"user_id" gorm:"index;comment:评论者ID"`
    PostID     uint           `json:"post_id" gorm:"index;comment:帖子ID"`
    ParentID   *uint          `json:"parent_id" gorm:"comment:父评论ID"`
}

// BandLike 点赞表
type BandLike struct {
    UserID      uint      `json:"user_id" gorm:"primaryKey;comment:用户ID"`
    ContentType string    `json:"content_type" gorm:"primaryKey;comment:内容类型(post/comment)"`
    ContentID   uint      `json:"content_id" gorm:"primaryKey;comment:内容ID"`
    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
} 


func (BandPost) TableName() string {
	return "bandgo_posts"
}

func (BandComment) TableName() string {
	return "bandgo_comments"
}

func (BandLike) TableName() string {
	return "bandgo_likes"
}




