package banner

import (
	"fmt"
	"mime/multipart"
	"peluang-server/domain"
	"peluang-server/internal/util"
	"strings"

	"github.com/google/uuid"
)

type bannerService struct {
	repo domain.BannerRepository
}

func NewService(repo domain.BannerRepository) domain.BannerService {
	return &bannerService{
		repo: repo,
	}
}

// CreateBanner implements domain.BannerService.
func (b *bannerService) CreateBanner(file *multipart.FileHeader) error {
	filename := fmt.Sprintf("%s.%s", util.ToSlug(strings.Split(file.Filename, ".")[0]), strings.Split(file.Filename, ".")[1])

	url, key, err := util.UploadFileToS3(file, filename, "banner")
	if err != nil {
		return err
	}

	banner := &domain.Banner{
		ID:     uuid.New().String(),
		ImgURL: url,
		Key:    key,
	}

	if err := b.repo.Insert(banner); err != nil {
		return err
	}
	return nil
}

// DeleteBanner implements domain.BannerService.
func (b *bannerService) DeleteBanner(id string) error {
	banner, err := b.repo.FindByID(id)
	if err != nil {
		return domain.ErrBannerNotFound
	}

	if err := util.DeleteFileFromS3(banner.Key); err != nil {
		return err
	}

	if err := b.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

// GetAllBanners implements domain.BannerService.
func (b *bannerService) GetAllBanners() ([]domain.Banner, error) {
	banners, err := b.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return banners, nil
}

// GetBanner implements domain.BannerService.
func (b *bannerService) GetBanner(id string) (*domain.Banner, error) {
	banner, err := b.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return banner, nil
}

// UpdateBanner implements domain.BannerService.
func (b *bannerService) UpdateBanner(id string, file *multipart.FileHeader) error {
	panic("unimplemented")
}
