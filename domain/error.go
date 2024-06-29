package domain

import "errors"

//auth error
var (
	ErrEmailExist        = errors.New("email already exists")
	ErrInvalidOTP        = errors.New("invalid otp")
	ErrExpiredOTP        = errors.New("otp has expired")
	ErrInvalidCredential = errors.New("invalid credential")
	ErrAlreadyVerified   = errors.New("user already verified")
)

//token error
var (
	ErrInvalidToken   = errors.New("invalid token")
	ErrWrongTypeToken = errors.New("wrong type token")
	ErrExpiredToken   = errors.New("token has expired")
	ErrEmptyToken     = errors.New("empty token")
	ErrNoBerearToken  = errors.New("token no Bearer")
)

//merchant error
var (
	ErrMerchantAlreadyExist  = errors.New("merchant already exist")
	ErrMerchantNotFound      = errors.New("merchant not found")
	ErrMerchantImageNotFound = errors.New("merchant image not found")
	ErrBannerNotFound        = errors.New("banner not found")
	ErrPackageNotFound       = errors.New("package not found")
)
