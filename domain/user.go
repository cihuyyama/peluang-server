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
	VerifiedAt time.Time `json:"verified_at"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id uuid.UUID) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Store(user *User) error
	Update(user *User) error
	Delete(id uuid.UUID) error
}

type UserService interface {
	Register(user *User, ctx context.Context) (int, error)
	Login(user *User, ctx context.Context) error
	GetUser(id uuid.UUID) (*User, error)
	GetAllUser() ([]User, error)
	ValidateOTP(id string, otp int) error
	ResendOTP(id string) (int, error)
}
