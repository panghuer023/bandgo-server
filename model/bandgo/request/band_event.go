package request

import "time"

// CreateEvent 创建活动请求
type CreateEvent struct {
    Title            string    `json:"title" binding:"required"`
    Description      string    `json:"description"`
    VenueName        string    `json:"venueName"`
    VenueAddress     string    `json:"venueAddress"`
    VenueCoordinates string    `json:"venueCoordinates"`
    StartDateTime    time.Time `json:"startDateTime" binding:"required"`
    EndDateTime      time.Time `json:"endDateTime"`
    EventType        string    `json:"eventType" binding:"required"`
    IsPublic         bool      `json:"isPublic"`
    CoverImageURL    string    `json:"coverImageUrl"`
}

// UpdateEvent 更新活动请求
type UpdateEvent struct {
    ID               uint      `json:"id" binding:"required"`
    Title            string    `json:"title"`
    Description      string    `json:"description"`
    VenueName        string    `json:"venueName"`
    VenueAddress     string    `json:"venueAddress"`
    VenueCoordinates string    `json:"venueCoordinates"`
    StartDateTime    time.Time `json:"startDateTime"`
    EndDateTime      time.Time `json:"endDateTime"`
    EventType        string    `json:"eventType"`
    IsPublic         bool      `json:"isPublic"`
    CoverImageURL    string    `json:"coverImageUrl"`
}

// EventSearch 活动搜索请求
type EventSearch struct {
    PageInfo
    EventType  string    `json:"eventType"`
    StartDate  time.Time `json:"startDate"`
    EndDate    time.Time `json:"endDate"`
    Location   string    `json:"location"`
    IsPublic   *bool     `json:"isPublic"`
    CreatorID  *uint     `json:"creatorId"`
} 