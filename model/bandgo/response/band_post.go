package response

// PostInfo 帖子信息响应
type PostInfo struct {
    ID                uint     `json:"id"`
    Content           string   `json:"content"`
    MediaURLs         []string `json:"mediaUrls"`
    CreatorType       string   `json:"creatorType"`
    Creator           UserInfo `json:"creator"`
    IsPublic          bool     `json:"isPublic"`
    IsPinned          bool     `json:"isPinned"`
    Location          string   `json:"location"`
    LikeCount         int      `json:"likeCount"`
    CommentCount      int      `json:"commentCount"`
    IsLiked          bool     `json:"isLiked"`
    CreatedAt         string   `json:"createdAt"`
}

// CommentInfo 评论信息响应
type CommentInfo struct {
    ID          uint     `json:"id"`
    Content     string   `json:"content"`
    MediaURLs   []string `json:"mediaUrls"`
    User        UserInfo `json:"user"`
    LikeCount   int      `json:"likeCount"`
    IsLiked     bool     `json:"isLiked"`
    CreatedAt   string   `json:"createdAt"`
    ParentID    *uint    `json:"parentId"`
} 