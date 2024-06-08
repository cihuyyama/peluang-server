package domain

import (
	"database/sql"
	"mime/multipart"
	"peluang-server/dto"
)

type Merchant struct {
	ID            string          `json:"id" gorm:"primaryKey;column:id"`
	Name          string          `json:"name" gorm:"unique"`
	Slug          string          `json:"slug" gorm:"unique"`
	Desc          string          `json:"desc"`
	Category      string          `json:"category"`
	BusinessModel string          `json:"business_model"`
	ImgUrl        string          `json:"img_url" gorm:"default:https://placehold.co/500x400.png"`
	Key           string          `json:"key" gorm:"default:null"`
	Images        []MerchantImage `json:"images" gorm:"foreignKey:MerchantID;references:ID"`
	VerifiedAt    sql.NullTime    `json:"verified_at"`
}

type MerchantImage struct {
	ID         string `json:"id" gorm:"primaryKey;column:id"`
	MerchantID string `json:"merchant_id" gorm:"column:merchant_id"`
	ImgURL     string `json:"img_url"`
	Key        string `json:"key"`
}

type MerchantRepository interface {
	FindAll() ([]Merchant, error)
	FindByID(id string) (*Merchant, error)
	FindImageByID(id string) (*MerchantImage, error)
	Insert(merchant *Merchant) error
	InsertImage(image *MerchantImage) error
	Update(merchant *Merchant) error
	UpdateAvatar(id, imgurl string, key string) error
	Delete(id string) error
	DeleteImage(merchantID string, imageID string) error
	DeleteAvatar(id string) error
}

type MerchantService interface {
	GetAllMerchants() ([]Merchant, error)
	GetMerchant(id string) (*Merchant, error)
	CreateMerchant(merchant *dto.MerchantRequest) error
	UpdateMerchant(id string, merchant *dto.MerchantRequest) error
	DeleteMerchant(id string) error

	UpdateAvatar(id string, file *multipart.FileHeader) error

	CreateImage(merchantID string, file *multipart.FileHeader) error
	DeleteImage(merchantID string, imageID string) error
}
