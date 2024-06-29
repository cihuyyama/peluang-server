package banner

import (
	"peluang-server/domain"

	"gorm.io/gorm"
)

type bannerRepository struct {
	db *gorm.DB
}

func NewRepository(con *gorm.DB) domain.BannerRepository {
	return &bannerRepository{
		db: con,
	}
}

// Delete implements domain.BannerRepository.
func (b *bannerRepository) Delete(id string) error {
	if tx := b.db.Delete(&domain.Banner{ID: id}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindAll implements domain.BannerRepository.
func (b *bannerRepository) FindAll() ([]domain.Banner, error) {
	var banners []domain.Banner
	if tx := b.db.Find(&banners); tx.Error != nil {
		return nil, tx.Error
	}
	return banners, nil
}

// FindByID implements domain.BannerRepository.
func (b *bannerRepository) FindByID(id string) (*domain.Banner, error) {
	var banner domain.Banner
	if tx := b.db.Where("id = ?", id).First(&banner); tx.Error != nil {
		return nil, tx.Error
	}
	return &banner, nil
}

// Insert implements domain.BannerRepository.
func (b *bannerRepository) Insert(banner *domain.Banner) error {
	if tx := b.db.Create(banner); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements domain.BannerRepository.
func (b *bannerRepository) Update(banner *domain.Banner) error {
	if tx := b.db.Model(&domain.Banner{}).Where("id = ?", banner.ID).Updates(banner); tx.Error != nil {
		return tx.Error
	}
	return nil
}
