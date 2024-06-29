package domain

import (
	"peluang-server/dto"
)

type Packages struct {
	ID         string          `json:"id" gorm:"primaryKey;column:id"`
	MerchantID string          `json:"merchant_id" gorm:"column:merchant_id"`
	Name       string          `json:"name"`
	Price      int             `json:"price"`
	List       []PackageList   `json:"list" gorm:"foreignKey:PackageID;references:ID"`
	Aditional  []AditionalList `json:"aditional" gorm:"foreignKey:PackageID;references:ID"`
}

type PackageList struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	PackageID string `json:"package_id"`
	Name      string `json:"name"`
}

type AditionalList struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	PackageID string `json:"package_id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
}

type PackageRepository interface {
	FindAll() ([]Packages, error)
	FindByID(id string) (*Packages, error)
	Insert(data *Packages) error
	InsertLists(data []PackageList) error
	InsertAditionals(data []AditionalList) error
	Update(data *Packages) error
	UpdateList(data PackageList) error
	UpdateAditional(data AditionalList) error
	Delete(id string) error
	DeleteList(id []PackageList) error
	DeleteAditional(id []AditionalList) error
}

type PackageService interface {
	FindAll() ([]Packages, error)
	FindByID(id string) (*Packages, error)
	Insert(data *dto.PackageRequest, merchantID string) error
	InsertLists(data []dto.List, packageID string) error
	InsertAditionals(data []dto.Aditional, packageID string) error
	Update(id string, data *dto.PackageRequest) error
	UpdateList(data PackageList) error
	UpdateAditional(data AditionalList) error
	Delete(id string) error
	DeleteList(id uint) error
	DeleteAditional(id uint) error
}
