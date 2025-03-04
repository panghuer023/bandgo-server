package request

// CreatePost 创建帖子请求
type CreatePost struct {
    Content     string   `json:"content"`
    MediaURLs   []string `json:"mediaUrls"`
    IsPublic    bool     `json:"isPublic"`
    Location    string   `json:"location"`
    EventID     *uint    `json:"eventId"`
}

// UpdatePost 更新帖子请求
type UpdatePost struct {
    ID          uint     `json:"id" binding:"required"`
    Content     string   `json:"content"`
    MediaURLs   []string `json:"mediaUrls"`
    IsPublic    bool     `json:"isPublic"`
    IsPinned    bool     `json:"isPinned"`
}

// CreateComment 创建评论请求
type CreateComment struct {
    PostID      uint     `json:"postId" binding:"required"`
    Content     string   `json:"content" binding:"required"`
    MediaURLs   []string `json:"mediaUrls"`
    ParentID    *uint    `json:"parentId"`
} 