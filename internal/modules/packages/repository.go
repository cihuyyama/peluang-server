package packages

import (
	"peluang-server/domain"

	"gorm.io/gorm"
)

type packagesRepository struct {
	db *gorm.DB
}

func NewRepository(con *gorm.DB) domain.PackageRepository {
	return &packagesRepository{
		db: con,
	}
}

// DeleteAditional implements domain.PackageRepository.
func (p *packagesRepository) DeleteAditional(list []domain.AditionalList) error {
	if tx := p.db.Delete(&list); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// InsertAditionals implements domain.PackageRepository.
func (p *packagesRepository) InsertAditionals(data []domain.AditionalList) error {
	if tx := p.db.Create(&data); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateAditional implements domain.PackageRepository.
func (p *packagesRepository) UpdateAditional(data domain.AditionalList) error {
	if tx := p.db.Model(&domain.AditionalList{}).Where("id = ?", data.ID).Updates(data); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateList implements domain.PackageRepository.
func (p *packagesRepository) UpdateList(data domain.PackageList) error {
	if tx := p.db.Model(&domain.PackageList{}).Where("id = ?", data.ID).Updates(data); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteLists implements domain.PackageRepository.
func (p *packagesRepository) DeleteList(list []domain.PackageList) error {
	if tx := p.db.Delete(&list); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// InsertLists implements domain.PackageRepository.
func (p *packagesRepository) InsertLists(data []domain.PackageList) error {
	if tx := p.db.Create(&data); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements domain.PackageRepository.
func (p *packagesRepository) Delete(id string) error {
	if tx := p.db.Delete(&domain.Packages{ID: id}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindAll implements domain.PackageRepository.
func (p *packagesRepository) FindAll() ([]domain.Packages, error) {
	var packages []domain.Packages
	if tx := p.db.
		Preload("List").
		Preload("Aditional").
		Find(&packages); tx.Error != nil {
		return nil, tx.Error
	}
	return packages, nil
}

// FindByID implements domain.PackageRepository.
func (p *packagesRepository) FindByID(id string) (*domain.Packages, error) {
	var packages domain.Packages
	if tx := p.db.
		Preload("List").
		Preload("Aditional").
		Where("id = ?", id).
		First(&packages); tx.Error != nil {
		return nil, tx.Error
	}
	return &packages, nil
}

// Insert implements domain.PackageRepository.
func (p *packagesRepository) Insert(data *domain.Packages) error {
	if tx := p.db.Create(data); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements domain.PackageRepository.
func (p *packagesRepository) Update(data *domain.Packages) error {
	if tx := p.db.Model(&domain.Packages{}).Where("id = ?", data.ID).Updates(data); tx.Error != nil {
		return tx.Error
	}
	return nil
}
