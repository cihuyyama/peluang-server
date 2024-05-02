package otp

import (
	"peluang-server/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(con *gorm.DB) domain.UserOtpRepository {
	return &repository{
		db: con,
	}
}

// FindByUserID implements domain.UserOtpRepository.
func (r *repository) FindByUserID(userID string) (*domain.UserOtp, error) {
	var userOtp *domain.UserOtp
	if tx := r.db.Where("user_id = ?", userID).First(&userOtp); tx.Error != nil {
		return &domain.UserOtp{}, tx.Error
	}
	return userOtp, nil
}

// Store implements domain.UserOtpRepository.
func (r *repository) Store(userOtp *domain.UserOtp) error {
	if tx := r.db.Create(userOtp); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements domain.UserOtpRepository.
func (r *repository) Update(userOtp *domain.UserOtp) error {
	if tx := r.db.Save(userOtp); tx.Error != nil {
		return tx.Error
	}
	return nil
}
