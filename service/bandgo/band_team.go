package bandgo

import (
    "errors"
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
    "gorm.io/gorm"
)

type BandTeamService struct{}

// CreateBand 创建乐队
func (s *BandTeamService) CreateBand(band *bandgo.BandTeam) error {
    if !errors.Is(global.GVA_DB.Where("name = ?", band.Name).First(&bandgo.BandTeam{}).Error, gorm.ErrRecordNotFound) {
        return errors.New("乐队名称已存在")
    }
    return global.GVA_DB.Create(band).Error
}

// UpdateBand 更新乐队信息
func (s *BandTeamService) UpdateBand(bandID uint, band *bandgo.BandTeam) error {
    return global.GVA_DB.Model(&bandgo.BandTeam{}).Where("id = ?", bandID).
        Omit("id", "creator_id").Updates(band).Error
}

// GetBandInfo 获取乐队信息
func (s *BandTeamService) GetBandInfo(bandID uint) (*bandgo.BandTeam, error) {
    var band bandgo.BandTeam
    err := global.GVA_DB.First(&band, bandID).Error
    return &band, err
}

// AddBandMember 添加乐队成员
func (s *BandTeamService) AddBandMember(member *bandgo.BandMember) error {
    // 检查是否已经是成员
    var count int64
    global.GVA_DB.Model(&bandgo.BandMember{}).
        Where("band_id = ? AND user_id = ? AND status = 'active'", member.BandID, member.UserID).
        Count(&count)
    if count > 0 {
        return errors.New("该用户已经是乐队成员")
    }
    return global.GVA_DB.Create(member).Error
}

// GetBandMembers 获取乐队成员列表
func (s *BandTeamService) GetBandMembers(bandID uint) ([]bandgo.BandMember, error) {
    var members []bandgo.BandMember
    err := global.GVA_DB.Where("band_id = ? AND status = 'active'", bandID).Find(&members).Error
    return members, err
}

// ProcessJoinRequest 处理加入申请
func (s *BandTeamService) ProcessJoinRequest(requestID uint, status string, responderID uint, responseMsg string) error {
    return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
        var request bandgo.BandJoinRequest
        if err := tx.First(&request, requestID).Error; err != nil {
            return err
        }

        // 更新申请状态
        if err := tx.Model(&request).Updates(map[string]interface{}{
            "status": status,
            "responded_by": responderID,
            "response_message": responseMsg,
        }).Error; err != nil {
            return err
        }

        // 如果同意申请，添加为成员
        if status == "accepted" {
            member := &bandgo.BandMember{
                BandID: request.BandID,
                UserID: request.UserID,
                Role: request.ProposedRole,
                Status: "active",
            }
            if err := tx.Create(member).Error; err != nil {
                return err
            }
        }

        return nil
    })
} 