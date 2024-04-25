package user

import (
	"peluang-server/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(con *gorm.DB) domain.UserRepository {
	return &repository{
		db: con,
	}
}

// FindByEmail implements domain.UserRepository.
func (r *repository) FindByEmail(email string) (user *domain.User, e error) {
	err := r.db.Where("email = ?", email).First(&user)
	if err != nil {
		return nil, err.Error
	}
	return user, nil
}

// FindAll implements domain.UserRepository.
func (r *repository) FindAll() ([]domain.User, error) {
	panic("unimplemented")
}

// FindByID implements domain.UserRepository.
func (r *repository) FindByID(id uuid.UUID) (*domain.User, error) {
	panic("unimplemented")
}

// FindByUsername implements domain.UserRepository.
func (r *repository) FindByUsername(username string) (*domain.User, error) {
	panic("unimplemented")
}

// Store implements domain.UserRepository.
func (r *repository) Store(user *domain.User) error {
	if err := r.db.Create(&user); err != nil {
		return err.Error
	}
	return nil
}

// Update implements domain.UserRepository.
func (r *repository) Update(user *domain.User) error {
	panic("unimplemented")
}

// Delete implements domain.UserRepository.
func (r *repository) Delete(id uuid.UUID) error {
	panic("unimplemented")
}
