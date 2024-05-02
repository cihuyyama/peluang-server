package domain

import (
	"time"

	"github.com/google/uuid"
)

type UserOtp struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	OTP       int       `json:"otp"`
	ExpiredAt int64     `json:"expired_at"`
	IssuedAt  time.Time `json:"issued_at"`
}

type UserOtpRepository interface {
	FindByUserID(userID string) (*UserOtp, error)
	Store(userOtp *UserOtp) error
	Update(userOtp *UserOtp) error
}

type UserOtpService interface {
}
