package request

// CreateBand 创建乐队请求
type CreateBand struct {
    Name          string `json:"name" binding:"required"`
    Description   string `json:"description"`
    Location      string `json:"location"`
    LogoURL       string `json:"logoUrl"`
    CoverImageURL string `json:"coverImageUrl"`
}

// UpdateBand 更新乐队信息请求
type UpdateBand struct {
    ID            uint   `json:"id" binding:"required"`
    Description   string `json:"description"`
    Location      string `json:"location"`
    LogoURL       string `json:"logoUrl"`
    CoverImageURL string `json:"coverImageUrl"`
    Status        string `json:"status"`
}

// BandJoinRequest 申请加入乐队请求
type BandJoinRequest struct {
    BandID       uint   `json:"bandId" binding:"required"`
    Message      string `json:"message"`
    ProposedRole string `json:"proposedRole" binding:"required"`
}

// ProcessJoinRequest 处理加入申请请求
type ProcessJoinRequest struct {
    RequestID       uint   `json:"requestId" binding:"required"`
    Status         string `json:"status" binding:"required,oneof=accepted rejected"`
    ResponseMessage string `json:"responseMessage"`
} 