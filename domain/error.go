package domain

import "errors"

var (
	ErrEmailExist   = errors.New("email already exists")
	ErrInavlidToken = errors.New("invalid token")
	ErrInvalidOTP   = errors.New("invalid otp")
	ErrExpiredOTP   = errors.New("otp has expired")
)
