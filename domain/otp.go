package domain

import (
	"time"
)

type UserOtp struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
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
