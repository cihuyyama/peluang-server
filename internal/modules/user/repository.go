package user

import (
	"peluang-server/domain"
	"peluang-server/internal/util"

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
	tx := r.db.Where("email = ?", &email).First(&user)
	if tx.Error != nil {
		return &domain.User{}, tx.Error
	}
	return user, nil
}

// FindAll implements domain.UserRepository.
func (r *repository) FindAll() ([]domain.User, error) {
	panic("unimplemented")
}

// FindByID implements domain.UserRepository.
func (r *repository) FindByID(id string) (*domain.User, error) {
	var user *domain.User
	if tx := r.db.Where("id = ?", id).First(&user); tx.Error != nil {
		return &domain.User{}, tx.Error
	}
	return user, nil
}

// FindByToken implements domain.UserRepository.
func (r *repository) FindByToken(token string) (*domain.User, error) {
	claims, err := util.GetClaims(token)
	if err != nil {
		return &domain.User{}, err
	}
	user, err := r.FindByID(claims["id"].(string))
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

// FindByUsername implements domain.UserRepository.
func (r *repository) FindByUsername(username string) (*domain.User, error) {
	panic("unimplemented")
}

// Store implements domain.UserRepository.
func (r *repository) Store(user *domain.User) error {
	if tx := r.db.Create(&user); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements domain.UserRepository.
func (r *repository) Update(user *domain.User) error {
	if tx := r.db.Save(&user); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements domain.UserRepository.
func (r *repository) Delete(id uuid.UUID) error {
	panic("unimplemented")
}
