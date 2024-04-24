package otp

import (
	"peluang-server/domain"

	"github.com/google/uuid"
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
func (r *repository) FindByUserID(userID uuid.UUID) (*domain.UserOtp, error) {
	panic("unimplemented")
}

// Store implements domain.UserOtpRepository.
func (r *repository) Store(userOtp *domain.UserOtp) error {
	if err := r.db.Create(userOtp); err != nil {
		return err.Error
	}
	return nil
}

// Update implements domain.UserOtpRepository.
func (r *repository) Update(userOtp *domain.UserOtp) error {
	panic("unimplemented")
}
