package response

import "time"

// EventInfo 活动信息响应
type EventInfo struct {
    ID               uint      `json:"id"`
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
    Creator          UserInfo  `json:"creator"`
    ParticipantCount int       `json:"participantCount"`
}

// EventParticipantInfo 活动参与者信息响应
type EventParticipantInfo struct {
    Type            string    `json:"type"` // user 或 band
    ID              uint      `json:"id"`
    Name            string    `json:"name"`
    Role            string    `json:"role"`
    Status          string    `json:"status"`
    ProfileImageURL string    `json:"profileImageUrl"`
} 