package user

import (
	"context"
	"peluang-server/domain"
	"peluang-server/internal/util"
	"time"

	"github.com/google/uuid"
)

type service struct {
	userRepo    domain.UserRepository
	userOtpRepo domain.UserOtpRepository
}

func NewService(userRepo domain.UserRepository, userOtpRepo domain.UserOtpRepository) domain.UserService {
	return &service{
		userRepo, userOtpRepo,
	}
}

// GetAllUser implements domain.UserService.
func (s *service) GetAllUser() ([]domain.User, error) {
	panic("unimplemented")
}

// GetUser implements domain.UserService.
func (s *service) GetUser(id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}

// Login implements domain.UserService.
func (s *service) Login(user *domain.User, ctx context.Context) error {
	panic("unimplemented")
}

// Register Test implements domain.UserService.
func (s *service) Register(user *domain.User, ctx context.Context) (otp int, err error) {
	_, err = s.userRepo.FindByEmail(user.Email)
	if err == nil {
		return 0, domain.ErrEmailExist
	}

	user.ID = uuid.New()
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword

	if err := s.userRepo.Store(user); err != nil {
		return 0, err
	}

	otp = util.GenerateOTP()
	err = s.userOtpRepo.Store(&domain.UserOtp{
		ID:        uuid.New(),
		UserID:    user.ID,
		OTP:       otp,
		ExpiredAt: time.Now().Add(time.Minute * 2).Unix(),
		IssuedAt:  time.Now(),
	})
	if err != nil {
		return 0, err
	}

	// AWS SES func
	// err = util.SendTemplatedEmailVerification(int64(otp), user.Email)
	// if err != nil {
	// 	return err
	// }

	return otp, nil
}

// ValidateOTP implements domain.UserService.
func (s *service) ValidateOTP(id string, otp int) error {
	panic("unimplemented")
}

// ResendOTP Test implements domain.UserService.
func (s *service) ResendOTP(id string) (int, error) {
	panic("unimplemented")
}
