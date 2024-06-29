package domain

import "mime/multipart"

type Banner struct {
	ID     string `json:"id" gorm:"primaryKey;column:id"`
	ImgURL string `json:"img_url"`
	Key    string `json:"key"`
}

type BannerRepository interface {
	FindAll() ([]Banner, error)
	FindByID(id string) (*Banner, error)
	Insert(banner *Banner) error
	Update(banner *Banner) error
	Delete(id string) error
}

type BannerService interface {
	GetAllBanners() ([]Banner, error)
	GetBanner(id string) (*Banner, error)
	CreateBanner(file *multipart.FileHeader) error
	UpdateBanner(id string, file *multipart.FileHeader) error
	DeleteBanner(id string) error
}
