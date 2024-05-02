package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email" gorm:"unique"`
	Password   string    `json:"password"`
	Telp       string    `json:"telp"`
	Role       string    `json:"role" gorm:"default:user"`
	ImgURL     string    `json:"img_url" gorm:"default:null"`
	VerifiedAt time.Time `json:"verified_at" gorm:"default:null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByToken(token string) (*User, error)
	FindByUsername(username string) (*User, error)
	Store(user *User) error
	Update(user *User) error
	Delete(id uuid.UUID) error
}

type UserService interface {
	Register(user *User, ctx context.Context) (*User, int, error)
	Login(user *User, ctx context.Context) (string, error)
	GetUser(token string) (*User, error)
	GetAllUser() ([]User, error)
	ValidateOTP(id string, otp int) error
	ResendOTP(id string) (int, error)
}
