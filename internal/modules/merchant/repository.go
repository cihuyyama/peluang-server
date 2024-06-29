package merchant

import (
	"peluang-server/domain"

	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

func NewRepository(con *gorm.DB) domain.MerchantRepository {
	return &merchantRepository{
		db: con,
	}
}

// FindBySlug implements domain.MerchantRepository.
func (m *merchantRepository) FindBySlug(slug string) (*domain.Merchant, error) {
	var merchant *domain.Merchant
	if tx := m.db.
		Preload("Images").
		Preload("Packages").
		Preload("Packages.List").
		Preload("Packages.Aditional").
		Where("slug = ?", slug).First(&merchant); tx.Error != nil {
		return &domain.Merchant{}, tx.Error
	}
	return merchant, nil
}

// Delete implements domain.MerchantRepository.
func (m *merchantRepository) Delete(id string) error {
	query := `DELETE FROM merchants WHERE id = ?`
	tx := m.db.Exec(query, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DeleteAvatar implements domain.MerchantRepository.
func (m *merchantRepository) DeleteAvatar(id string) error {
	panic("unimplemented")
}

// DeleteImage implements domain.MerchantRepository.
func (m *merchantRepository) DeleteImage(merchantID string, imageID string) error {
	merchantImage := domain.MerchantImage{ID: imageID, MerchantID: merchantID}
	if tx := m.db.Delete(&merchantImage); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// InsertImage implements domain.MerchantRepository.
func (m *merchantRepository) InsertImage(image *domain.MerchantImage) error {
	merchant := domain.Merchant{ID: image.MerchantID}

	if err := m.db.Model(&merchant).Association("Images").Append(image); err != nil {
		return err
	}
	return nil

}

// Update implements domain.MerchantRepository.
func (m *merchantRepository) Update(merchant *domain.Merchant) error {
	// Directly update the merchant using its ID
	if tx := m.db.Model(&domain.Merchant{}).Where("id = ?", merchant.ID).Updates(merchant); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// UpdateAvatar implements domain.MerchantRepository.
func (m *merchantRepository) UpdateAvatar(id string, imgurl string, key string) error {
	// Directly update the merchant's img_url and key using the ID
	if tx := m.db.Model(&domain.Merchant{}).Where("id = ?", id).Updates(map[string]interface{}{"img_url": imgurl, "key": key}); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FindAll implements domain.MerchantRepository.
func (m *merchantRepository) FindAll() ([]domain.Merchant, error) {
	var merchants []domain.Merchant
	if tx := m.db.Find(&merchants); tx.Error != nil {
		return []domain.Merchant{}, tx.Error
	}
	return merchants, nil
}

// FindByID implements domain.MerchantRepository.
func (m *merchantRepository) FindByID(id string) (*domain.Merchant, error) {
	var merchant *domain.Merchant
	if tx := m.db.Preload("Images").Preload("Packages").Where("id = ?", id).First(&merchant); tx.Error != nil {
		return &domain.Merchant{}, tx.Error
	}
	return merchant, nil
}

// FindImageByID implements domain.MerchantRepository.
func (m *merchantRepository) FindImageByID(id string) (*domain.MerchantImage, error) {
	var merchantImage *domain.MerchantImage
	if tx := m.db.Where("id = ?", id).First(&merchantImage); tx.Error != nil {
		return &domain.MerchantImage{}, tx.Error
	}
	return merchantImage, nil
}

// Insert implements domain.MerchantRepository.
func (m *merchantRepository) Insert(merchant *domain.Merchant) error {
	if tx := m.db.Create(merchant); tx.Error != nil {
		return tx.Error
	}
	return nil
}
