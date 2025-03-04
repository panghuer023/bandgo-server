package response

// BandInfo 乐队信息响应
type BandInfo struct {
    ID            uint   `json:"id"`
    Name          string `json:"name"`
    Description   string `json:"description"`
    Location      string `json:"location"`
    LogoURL       string `json:"logoUrl"`
    CoverImageURL string `json:"coverImageUrl"`
    Status        string `json:"status"`
    CreatorID     uint   `json:"creatorId"`
    MemberCount   int    `json:"memberCount"`
    IsVerified    bool   `json:"isVerified"`
}

// BandMemberInfo 乐队成员信息响应
type BandMemberInfo struct {
    UserInfo
    Role        string `json:"role"`
    JoinDate    string `json:"joinDate"`
    IsAdmin     bool   `json:"isAdmin"`
}

// BandJoinRequestInfo 加入申请信息响应
type BandJoinRequestInfo struct {
    ID              uint     `json:"id"`
    User            UserInfo `json:"user"`
    Message         string   `json:"message"`
    ProposedRole    string   `json:"proposedRole"`
    Status          string   `json:"status"`
    CreatedAt       string   `json:"createdAt"`
} 