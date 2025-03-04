package bandgo

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bandgo"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bandgo/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type BandUserService struct{}

// Register 用户注册
func (s *BandUserService) Register(req request.Register) (userInter *bandgo.BandUser, err error) {
	if !errors.Is(global.GVA_DB.Where("username = ?", req.Username).First(&bandgo.BandUser{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名已存在")
	}
	if !errors.Is(global.GVA_DB.Where("email = ?", req.Email).First(&bandgo.BandUser{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("邮箱已被使用")
	}

	user := &bandgo.BandUser{
		Username:     req.Username,
		Email:        req.Email,
		FullName:     req.FullName,
		Password: utils.BcryptHash(req.Password),
	}
	err = global.GVA_DB.Create(user).Error
	return user, err
}

// Login 用户登录
func (s *BandUserService) Login(req request.Login) (*bandgo.BandUser, error) {
	var user bandgo.BandUser
	err := global.GVA_DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

// UpdateProfile 更新用户资料
func (s *BandUserService) UpdateProfile(userID uint, req request.UpdateProfile) error {
	return global.GVA_DB.Model(&bandgo.BandUser{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"bio":               req.Bio,
			"profile_image_url": req.ProfileImageURL,
			"cover_image_url":   req.CoverImageURL,
			"location":          req.Location,
			"phone_number":      req.PhoneNumber,
		}).Error
}

// GetUserProfile 获取用户资料
func (s *BandUserService) GetUserProfile(userID uint) (*bandgo.BandUser, error) {
	var user bandgo.BandUser
	err := global.GVA_DB.First(&user, userID).Error
	return &user, err
}

// AddUserSkill 添加用户技能
func (s *BandUserService) AddUserSkill(userID uint, skill *bandgo.BandUserSkill) error {
	skill.UserID = userID
	return global.GVA_DB.Create(skill).Error
}

// UpdateUserSkill 更新用户技能
func (s *BandUserService) UpdateUserSkill(userID uint, skillID uint, skill *bandgo.BandUserSkill) error {
	return global.GVA_DB.Model(&bandgo.BandUserSkill{}).
		Where("id = ? AND user_id = ?", skillID, userID).
		Updates(skill).Error
}

// GetUserSkills 获取用户所有技能
func (s *BandUserService) GetUserSkills(userID uint) ([]bandgo.BandUserSkill, error) {
	var skills []bandgo.BandUserSkill
	err := global.GVA_DB.Where("user_id = ?", userID).Find(&skills).Error
	return skills, err
}

// UpdatePassword 更新密码
func (s *BandUserService) UpdatePassword(userID uint, oldPassword, newPassword string) error {
	var user bandgo.BandUser
	err := global.GVA_DB.First(&user, userID).Error
	if err != nil {
		return err
	}

	if ok := utils.BcryptCheck(oldPassword, user.Password); !ok {
		return errors.New("原密码错误")
	}

	return global.GVA_DB.Model(&user).Update("password_hash", utils.BcryptHash(newPassword)).Error
}

// GetUserGenres 获取用户音乐类型偏好
func (s *BandUserService) GetUserGenres(userID uint) ([]bandgo.BandUserGenre, error) {
	var genres []bandgo.BandUserGenre
	err := global.GVA_DB.Where("user_id = ?", userID).Find(&genres).Error
	return genres, err
}

// UpdateUserGenres 更新用户音乐类型偏好
func (s *BandUserService) UpdateUserGenres(userID uint, genres []bandgo.BandUserGenre) error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 删除旧的偏好
		if err := tx.Where("user_id = ?", userID).Delete(&bandgo.BandUserGenre{}).Error; err != nil {
			return err
		}

		// 添加新的偏好
		for i := range genres {
			genres[i].UserID = userID
		}
		return tx.Create(&genres).Error
	})
}
