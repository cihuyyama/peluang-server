package user

import (
	"context"
	"fmt"
	"peluang-server/domain"
	"peluang-server/dto"
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
func (s *service) GetUser(ctx context.Context) (*domain.User, error) {
	userID := ctx.Value("x-userid").(string)
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

// Login implements domain.UserService.
func (s *service) Login(userReq *dto.LoginRequest, ctx context.Context) (string, error) {
	user := new(domain.User)
	user.Email = userReq.Email
	user.Password = userReq.Password

	userRepo, err := s.userRepo.FindByEmail(user.Email)
	if err != nil {
		return "", domain.ErrInvalidCredential
	}

	if _, err := util.CheckPasswordHash(user.Password, userRepo.Password); err != nil {
		fmt.Println(user.Password, userRepo.Password)
		return "", err
	}

	token, err := util.GenerateToken(userRepo)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Register Test implements domain.UserService.
func (s *service) Register(user *domain.User, ctx context.Context) (*domain.User, int, error) {
	_, err := s.userRepo.FindByEmail(user.Email)
	if err == nil {
		return &domain.User{}, 0, domain.ErrEmailExist
	}

	user.ID = uuid.New().String()
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return &domain.User{}, 0, err
	}
	user.Password = hashedPassword

	if err := s.userRepo.Store(user); err != nil {
		return &domain.User{}, 0, err
	}

	otp := util.GenerateOTP()
	err = s.userOtpRepo.Store(&domain.UserOtp{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		OTP:       otp,
		ExpiredAt: time.Now().Add(time.Minute * 2).Unix(),
		IssuedAt:  time.Now(),
	})
	if err != nil {
		return &domain.User{}, 0, err
	}

	// AWS SES func
	err = util.SendTemplatedEmailVerification(int64(otp), user.Email)
	if err != nil {
		return user, otp, err
	}

	return user, otp, nil
}

// ValidateOTP implements domain.UserService.
func (s *service) ValidateOTP(id string, otp int) error {
	userOtp, err := s.userOtpRepo.FindByUserID(id)
	if err != nil {
		return err
	}

	if userOtp.OTP != otp {
		return domain.ErrInvalidOTP
	}

	if time.Now().Unix() > userOtp.ExpiredAt {
		return domain.ErrExpiredOTP
	}

	user, err := s.userRepo.FindByID(userOtp.UserID)
	if err != nil {
		return err
	} else if user.VerifiedAt.Valid {
		return domain.ErrAlreadyVerified
	}

	user.VerifiedAt.Time = time.Now()

	err = s.userRepo.Update(user)
	if err != nil {
		return err
	}
	return nil
}

// ResendOTP Test implements domain.UserService.
func (s *service) ResendOTP(id string) (int, error) {
	otp := util.GenerateOTP()

	userOtp, err := s.userOtpRepo.FindByUserID(id)
	if err != nil {
		return 0, err
	}
	userOtp.OTP = otp
	userOtp.ExpiredAt = time.Now().Add(time.Minute * 2).Unix()

	err = s.userOtpRepo.Update(userOtp)
	if err != nil {
		return 0, err
	}

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return 0, err
	} else if user.VerifiedAt.Valid {
		return 0, domain.ErrAlreadyVerified
	}

	// AWS SES func
	err = util.SendTemplatedEmailVerification(int64(otp), user.Email)
	if err != nil {
		return otp, err
	}

	return otp, nil
}
