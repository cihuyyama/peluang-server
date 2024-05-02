package dto

type OTPRequest struct {
	UserID string `json:"user_id" validate:"required"`
	OTP    int    `json:"otp" validate:"required"`
}

type ReOTPRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
