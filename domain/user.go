package domain

import (
	"context"
	"database/sql"
	"peluang-server/dto"
	"time"
)

type User struct {
	ID         string         `json:"id"`
	Username   string         `json:"username"`
	Email      string         `json:"email" gorm:"unique"`
	Password   string         `json:"password"`
	Telp       string         `json:"telp"`
	Role       string         `json:"role" gorm:"default:user"`
	ImgURL     sql.NullString `json:"img_url" gorm:"default:null"`
	VerifiedAt sql.NullTime   `json:"verified_at" gorm:"default:null"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByToken(token string) (*User, error)
	FindByUsername(username string) (*User, error)
	Store(user *User) error
	Update(user *User) error
	Delete(id string) error
}

type UserService interface {
	Register(user *User, ctx context.Context) (*User, int, error)
	Login(user *dto.LoginRequest, ctx context.Context) (string, error)
	GetUser(ctx context.Context) (*User, error)
	GetAllUser() ([]User, error)
	ValidateOTP(id string, otp int) error
	ResendOTP(id string) (int, error)
}
