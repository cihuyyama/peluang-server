package merchant

import (
	"fmt"
	"mime/multipart"
	"peluang-server/domain"
	"peluang-server/dto"
	"peluang-server/internal/util"
	"strings"

	"github.com/google/uuid"
)

type merchantService struct {
	merchantRepo domain.MerchantRepository
}

func NewService(merchantRepo domain.MerchantRepository) domain.MerchantService {
	return &merchantService{
		merchantRepo,
	}
}

// CreateImage implements domain.MerchantService.
func (m *merchantService) CreateImage(merchantID string, file *multipart.FileHeader) error {
	merchantDomain, err := m.merchantRepo.FindByID(merchantID)
	if err != nil {
		return domain.ErrMerchantNotFound
	}

	filename := fmt.Sprintf("%s.%s", util.ToSlug(strings.Split(file.Filename, ".")[0]), strings.Split(file.Filename, ".")[1])

	url, key, err := util.UploadFileToS3(file, filename, merchantDomain.Slug)
	if err != nil {
		return err
	}

	image := &domain.MerchantImage{
		ID:         uuid.New().String(),
		MerchantID: merchantID,
		ImgURL:     url,
		Key:        key,
	}

	if err := m.merchantRepo.InsertImage(image); err != nil {
		return err
	}
	return nil
}

// DeleteImage implements domain.MerchantService.
func (m *merchantService) DeleteImage(merchantID string, imageID string) error {
	merchantImage, err := m.merchantRepo.FindImageByID(imageID)
	if err != nil {
		return domain.ErrMerchantImageNotFound
	}

	if err := util.DeleteFileFromS3(merchantImage.Key); err != nil {
		return err
	}

	if err := m.merchantRepo.DeleteImage(merchantImage.MerchantID, imageID); err != nil {
		return err
	}
	return nil
}

// DeleteMerchant implements domain.MerchantService.
func (m *merchantService) DeleteMerchant(id string) error {
	if _, err := m.merchantRepo.FindByID(id); err != nil {
		return domain.ErrMerchantNotFound
	}

	if err := m.merchantRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

// UpdateAvatar implements domain.MerchantService.
func (m *merchantService) UpdateAvatar(id string, file *multipart.FileHeader) error {
	merchantDomain, err := m.merchantRepo.FindByID(id)
	if err != nil {
		return domain.ErrMerchantNotFound
	}

	if merchantDomain.ImgUrl != "https://placehold.co/500x400.png" {
		if err := util.DeleteFileFromS3(merchantDomain.Key); err != nil {
			return err
		}
	}

	filename := fmt.Sprintf("avatar.%s", strings.Split(file.Filename, ".")[1])

	url, key, err := util.UploadFileToS3(file, filename, merchantDomain.Slug)
	if err != nil {
		return err
	}

	if err := m.merchantRepo.UpdateAvatar(id, url, key); err != nil {
		return err
	}

	return nil
}

// UpdateMerchant implements domain.MerchantService.
func (m *merchantService) UpdateMerchant(id string, merchant *dto.MerchantRequest) error {
	merchantDomain, err := m.merchantRepo.FindByID(id)
	if err != nil {
		return domain.ErrMerchantNotFound
	}

	merchantDomain.Name = merchant.Name
	merchantDomain.Desc = merchant.Desc
	merchantDomain.Category = merchant.Category
	merchantDomain.BusinessModel = merchant.BusinessModel

	if err := m.merchantRepo.Update(merchantDomain); err != nil {
		return err
	}

	return nil
}

// CreateMerchant implements domain.MerchantService.
func (m *merchantService) CreateMerchant(merchant *dto.MerchantRequest) error {
	if _, err := m.merchantRepo.FindByID(merchant.Name); err == nil {
		return domain.ErrMerchantAlreadyExist
	}

	merchantDomain := &domain.Merchant{
		ID:            uuid.New().String(),
		Name:          merchant.Name,
		Slug:          util.ToSlug(merchant.Name),
		Desc:          merchant.Desc,
		Category:      merchant.Category,
		BusinessModel: merchant.BusinessModel,
	}

	if err := m.merchantRepo.Insert(merchantDomain); err != nil {
		return err
	}

	return nil
}

// GetAllMerchants implements domain.MerchantService.
func (m *merchantService) GetAllMerchants() ([]domain.Merchant, error) {
	merchants, err := m.merchantRepo.FindAll()
	if err != nil {
		return []domain.Merchant{}, err
	}
	return merchants, nil
}

// GetMerchant implements domain.MerchantService.
func (m *merchantService) GetMerchant(id string) (*domain.Merchant, error) {
	merchant, err := m.merchantRepo.FindByID(id)
	if err != nil {
		return &domain.Merchant{}, err
	}
	return merchant, nil
}
